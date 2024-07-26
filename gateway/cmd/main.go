package main

import (
	"fmt"
	"log"

	"gateway/internal/config"

	grpcServer "gateway/internal/grpc"
	"gateway/internal/k8s"
	"gateway/internal/server"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	cfg := config.Get()
	log.Printf("Starting server on port: %d", cfg.Server.Port)

	// gRPC 서버 생성
	keygenServer := grpcServer.NewKeygenServiceServer()

	// gRPC 서버 실행 (별도의 고루틴에서)
	go grpcServer.StartGRPCServer(keygenServer)

	// 서버 시작 시 기존 Pod 확인 및 대기 풀에 추가
	existingPods, err := k8s.ListExistingPods()
	if err != nil {
		log.Fatalf("Failed to list existing pods: %v", err)
	}
	fmt.Println("existingPods: ", existingPods)
	podPool := k8s.GetPodPool()
	for _, pod := range existingPods {
		podPool.AddPod(pod)
	}
	log.Printf("Added %d existing pods to the pool", len(existingPods))

	// 필요한 경우 새로운 Pod 생성
	initialPodCount := cfg.Kubernetes.InitialPodCount
	fmt.Println("initialPodCount", initialPodCount)
	if len(existingPods) < initialPodCount {
		if err := k8s.CreatePods(initialPodCount - len(existingPods)); err != nil {
			log.Fatalf("Failed to create initial pods: %v", err)
		}
	}

	// HTTP 서버에 keygenServer 전달
	srv := server.NewServer(keygenServer)
	srv.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
