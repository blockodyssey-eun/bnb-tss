package k8s

import (
	"fmt"
	"sync"

	v1 "k8s.io/api/core/v1"
)

type PodPool struct {
	mu   sync.Mutex
	pods []*v1.Pod
}

func NewPodPool() *PodPool {
	return &PodPool{
		pods: make([]*v1.Pod, 0),
	}
}

func (p *PodPool) AddPod(pod *v1.Pod) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pods = append(p.pods, pod)
}

func (p *PodPool) GetPod() *v1.Pod {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.pods) == 0 {
		return nil
	}
	pod := p.pods[0]
	p.pods = p.pods[1:]
	return pod
}

func (p *PodPool) GetAvailablePods(m int) ([]*v1.Pod, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	var availablePods []*v1.Pod
	for _, pod := range p.pods {
		if CheckPodResourceAvailability(pod) {
			availablePods = append(availablePods, pod)
			if len(availablePods) == m {
				break
			}
		}
	}
	if len(availablePods) < m {
		return nil, fmt.Errorf("not enough available pods in the pool")
	}
	return availablePods, nil
}
