package grpc

import (
	"context"
	"fmt"
	"time"

	"gateway/internal/proto"

	"google.golang.org/grpc"
	corev1 "k8s.io/api/core/v1"
)

func CallKeygenService(podIP string, n, m int32, allPods []*corev1.Pod) (*proto.KeygenResponse, error) {
	fmt.Println("podIP:", podIP)
	conn, err := grpc.Dial(fmt.Sprintf("%s:50051", podIP), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to pod %s: %v", podIP, err)
	}
	defer conn.Close()

	client := proto.NewKeygenServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	podInfos := make([]*proto.PodInfo, len(allPods))
	for i, pod := range allPods {
		podInfos[i] = &proto.PodInfo{
			Ip:   pod.Status.PodIP,
			Port: 50051, // 모든 Pod이 같은 포트를 사용한다고 가정
		}
	}

	req := &proto.KeygenRequest{
		N:    n,
		M:    m,
		Pods: podInfos,
	}

	return client.GenerateKey(ctx, req)
}
