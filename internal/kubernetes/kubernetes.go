package kubernetes

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type PartyPod struct {
	ID   string
	Name string
	IP   string
}

// getClientset returns a Kubernetes clientset using the in-cluster config if available, otherwise falls back to local kubeconfig
func getClientset() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error

	if _, err = rest.InClusterConfig(); err == nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to create in-cluster config: %v", err)
		}
	} else {
		// Path to the local kubeconfig file
		kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, fmt.Errorf("failed to build config from flags: %v", err)
		}
	}

	// Create the Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %v", err)
	}

	return clientset, nil
}

// CreatePartyPods creates the specified number of party pods
func CreatePartyPods(totalParties int) ([]PartyPod, error) {
	clientset, err := getClientset()
	if err != nil {
		return nil, err
	}

	var partyPods []PartyPod

	for i := 0; i < totalParties; i++ {
		podName := fmt.Sprintf("tss-party-%d", i)

		// Create deployment
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
								Name:  podName,
								Image: "tss-party:latest",
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

		// Create Deployment
		result, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to create deployment: %v", err)
		}

		// Wait for the pod to be ready
		pod, err := waitForPod(clientset, result.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to wait for pod: %v", err)
		}

		partyPod := PartyPod{
			ID:   strconv.Itoa(i),
			Name: pod.Name,
			IP:   pod.Status.PodIP,
		}
		partyPods = append(partyPods, partyPod)
	}

	return partyPods, nil
}

// DeletePartyPods deletes the party pods
func DeletePartyPods(partyPods []PartyPod) error {
	clientset, err := getClientset()
	if err != nil {
		return fmt.Errorf("failed to create clientset: %v", err)
	}

	for _, pod := range partyPods {
		err := clientset.AppsV1().Deployments("default").Delete(context.TODO(), pod.Name, metav1.DeleteOptions{})
		if err != nil {
			return fmt.Errorf("failed to delete deployment %s: %v", pod.Name, err)
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
