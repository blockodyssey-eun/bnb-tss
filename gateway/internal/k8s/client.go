package k8s

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gateway/internal/config"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var podPool *PodPool

func init() {
	podPool = NewPodPool()
}

func getClientset() (*kubernetes.Clientset, error) {
	if _, exists := os.LookupEnv("KUBERNETES_SERVICE_HOST"); exists {
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to get in-cluster config: %v", err)
		}
		return kubernetes.NewForConfig(config)
	} else {
		kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, fmt.Errorf("failed to get kubeconfig: %v", err)
		}
		return kubernetes.NewForConfig(config)
	}
}

func ListExistingPods() ([]*corev1.Pod, error) {
	clientset, err := getClientset()
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %v", err)
	}

	cfg := config.Get()

	// 디버깅: 네임스페이스와 라벨 출력
	log.Printf("Searching for pods in namespace: %s with label: app=tss-party", cfg.Kubernetes.Namespace)

	pods, err := clientset.CoreV1().Pods(cfg.Kubernetes.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=keygen",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list existing pods: %v", err)
	}

	// 디버깅: 찾은 Pod 수 출력
	log.Printf("Found %d pods", len(pods.Items))

	var existingPods []*corev1.Pod
	for _, pod := range pods.Items {
		// 디버깅: 각 Pod의 이름과 상태 출력
		log.Printf("Pod: %s, Status: %s", pod.Name, pod.Status.Phase)
		if pod.Status.Phase == corev1.PodRunning {
			existingPods = append(existingPods, &pod)
		}
	}

	return existingPods, nil
}

func CreatePods(m int) error {
	clientset, err := getClientset()
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %v", err)
	}

	cfg := config.Get()

	for i := 0; i < m; i++ {
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("%s-%d", cfg.Kubernetes.PodPrefix, i),
				Labels: map[string]string{
					"app": "tss-party",
				},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:            "tss-party-container",
						Image:           cfg.Kubernetes.PodPrefix,
						ImagePullPolicy: corev1.PullNever,

						Ports: []corev1.ContainerPort{
							{
								ContainerPort: 50051,
							},
						},
					},
				},
				RestartPolicy: corev1.RestartPolicyOnFailure,
			},
		}

		createdPod, err := clientset.CoreV1().Pods(cfg.Kubernetes.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
		if err != nil {
			return fmt.Errorf("failed to create pod %d: %v", i, err)
		}

		// Pod이 Running 상태가 될 때까지 대기
		err = waitForPodRunning(clientset, createdPod.Name, cfg.Kubernetes.Namespace)
		if err != nil {
			return fmt.Errorf("error waiting for pod to be running: %v", err)
		}

		log.Printf("Created pod: %s in namespace %s with IP: %s", createdPod.Name, createdPod.Namespace, createdPod.Status.PodIP)
		podPool.AddPod(createdPod)
	}

	return nil
}

func GetPodFromPool() *corev1.Pod {
	return podPool.GetPod()
}

func GetPodsFromPool(m int) ([]*corev1.Pod, error) {
	var selectedPods []*corev1.Pod
	for i := 0; i < m; i++ {
		pod := podPool.GetPod()
		if pod == nil {
			return nil, fmt.Errorf("not enough pods in the pool")
		}
		selectedPods = append(selectedPods, pod)
	}
	return selectedPods, nil
}

func CheckPodResourceAvailability(pod *corev1.Pod) bool {
	// Pod의 리소스 상태를 확인하는 로직을 추가합니다.
	// 예를 들어, Pod의 상태가 Running인지 확인합니다.
	return pod.Status.Phase == corev1.PodRunning
}

func GetPodPool() *PodPool {
	return podPool
}

func waitForPodRunning(clientset *kubernetes.Clientset, podName, namespace string) error {
	return wait.PollImmediate(time.Second, time.Minute*5, func() (bool, error) {
		pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		return pod.Status.Phase == corev1.PodRunning, nil
	})
}
