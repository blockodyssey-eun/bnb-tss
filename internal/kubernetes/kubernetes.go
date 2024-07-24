package kubernetes

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"path/filepath"

	"github.com/bnb-chain/tss-lib/tss"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const (
	inUseLabel = "tss.inuse"
	poolSize   = 5 // 미리 생성할 파드 수
)

type PartyPod struct {
	ID    string
	Name  string
	IP    string
	Party *party
}

type party struct {
	id        *tss.PartyID
	params    *tss.Parameters
	out       chan tss.Message
	in        chan tss.Message
	shareData []byte
	sendMsg   Sender
	logger    *zap.SugaredLogger
	closeChan chan struct{}
}

type Sender func(msg []byte, isBroadcast bool, to uint16)

func NewParty(id uint16, logger *zap.SugaredLogger) *party {
	return &party{
		id:        tss.NewPartyID(fmt.Sprintf("%d", id), "", big.NewInt(int64(id))),
		out:       make(chan tss.Message, 1000),
		in:        make(chan tss.Message, 1000),
		logger:    logger,
		closeChan: make(chan struct{}),
	}
}

var (
	podPool       = make(chan *corev1.Pod, poolSize)
	poolInitOnce  sync.Once
	clientset     *kubernetes.Clientset
	clientsetInit sync.Once
)

// getClientset returns a Kubernetes clientset using the in-cluster config if available, otherwise falls back to local kubeconfig
func getClientset() (*kubernetes.Clientset, error) {
	var err error
	clientsetInit.Do(func() {
		var config *rest.Config
		if _, err = rest.InClusterConfig(); err == nil {
			config, err = rest.InClusterConfig()
			if err != nil {
				err = fmt.Errorf("failed to create in-cluster config: %v", err)
				return
			}
		} else {
			kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
			config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
			if err != nil {
				err = fmt.Errorf("failed to build config from flags: %v", err)
				return
			}
		}

		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			err = fmt.Errorf("failed to create clientset: %v", err)
		}
	})
	return clientset, err
}

// initPodPool initializes the pod pool
func initPodPool() {
	clientset, err := getClientset()
	if err != nil {
		fmt.Printf("Failed to get clientset: %v\n", err)
		return
	}

	for i := 0; i < poolSize; i++ {
		podName := fmt.Sprintf("tss-party-pool-%d", i)
		deployment, err := createOrUpdatePartyDeployment(clientset, podName)
		if err != nil {
			fmt.Printf("Failed to create or update deployment: %v\n", err)
			continue
		}

		pod, err := waitForPod(clientset, deployment.Name)
		if err != nil {
			fmt.Printf("Failed to wait for pod: %v\n", err)
			continue
		}

		podPool <- pod
	}
}

// createOrUpdatePartyDeployment creates or updates a deployment for a party pod
func createOrUpdatePartyDeployment(clientset *kubernetes.Clientset, podName string) (*appsv1.Deployment, error) {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": podName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": podName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            podName,
							Image:           "tss-party:latest",
							ImagePullPolicy: corev1.PullNever,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 9090,
								},
							},
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("100m"),
									corev1.ResourceMemory: resource.MustParse("128Mi"),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("500m"),
									corev1.ResourceMemory: resource.MustParse("512Mi"),
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		if k8serrors.IsAlreadyExists(err) {
			existingDeployment, getErr := clientset.AppsV1().Deployments("default").Get(context.TODO(), podName, metav1.GetOptions{})
			if getErr != nil {
				return nil, fmt.Errorf("failed to get existing deployment: %v", getErr)
			}
			existingDeployment.Spec = deployment.Spec
			result, err = clientset.AppsV1().Deployments("default").Update(context.TODO(), existingDeployment, metav1.UpdateOptions{})
			if err != nil {
				return nil, fmt.Errorf("failed to update existing deployment: %v", err)
			}
			return result, nil
		}
		return nil, fmt.Errorf("failed to create deployment: %v", err)
	}
	return result, nil
}

// findAvailablePod finds an available pod from the pool
func findAvailablePod() (*corev1.Pod, error) {
	select {
	case pod := <-podPool:
		return pod, nil
	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("no available pod in the pool")
	}
}

// markPodAsInUse marks a pod as in use
func markPodAsInUse(clientset *kubernetes.Clientset, pod *corev1.Pod) error {
	if pod.Labels == nil {
		pod.Labels = make(map[string]string)
	}
	pod.Labels[inUseLabel] = "true"

	_, err := clientset.CoreV1().Pods(pod.Namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	return err
}

// markPodAsNotInUse marks a pod as not in use and returns it to the pool
func markPodAsNotInUse(clientset *kubernetes.Clientset, pod *corev1.Pod) error {
	delete(pod.Labels, inUseLabel)

	_, err := clientset.CoreV1().Pods(pod.Namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err == nil {
		podPool <- pod
	}
	return err
}

// CreatePartyPods creates the specified number of party pods
func CreatePartyPods(totalParties int) ([]PartyPod, error) {
	poolInitOnce.Do(initPodPool)

	clientset, err := getClientset()
	if err != nil {
		return nil, err
	}

	var partyPods []PartyPod

	for i := 0; i < totalParties; i++ {
		availablePod, err := findAvailablePod()
		if err != nil {
			return nil, fmt.Errorf("failed to find available pod: %v", err)
		}

		err = markPodAsInUse(clientset, availablePod)
		if err != nil {
			return nil, fmt.Errorf("failed to mark pod as in use: %v", err)
		}

		partyPod := PartyPod{
			ID:   strconv.Itoa(i),
			Name: availablePod.Name,
			IP:   availablePod.Status.PodIP,
		}
		partyPods = append(partyPods, partyPod)
	}

	return partyPods, nil
}

// DeletePartyPods marks the party pods as not in use
func DeletePartyPods(partyPods []PartyPod) error {
	clientset, err := getClientset()
	if err != nil {
		return fmt.Errorf("failed to create clientset: %v", err)
	}

	for _, partyPod := range partyPods {
		pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), partyPod.Name, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("failed to get pod %s: %v", partyPod.Name, err)
		}

		err = markPodAsNotInUse(clientset, pod)
		if err != nil {
			return fmt.Errorf("failed to mark pod %s as not in use: %v", partyPod.Name, err)
		}
	}

	return nil
}

func int32Ptr(i int32) *int32 { return &i }

func waitForPod(clientset *kubernetes.Clientset, deploymentName string) (*corev1.Pod, error) {
	for {
		pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{
			LabelSelector: fmt.Sprintf("app=%s", deploymentName),
		})
		if err != nil {
			return nil, err
		}
		if len(pods.Items) > 0 && pods.Items[0].Status.Phase == corev1.PodRunning {
			return &pods.Items[0], nil
		}
		time.Sleep(2 * time.Second)
	}
}
