package main

import (
	"log"

	"party/internal/config"
	"party/internal/grpc"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	cfg := config.Get()

	server := grpc.NewServer()
	if err := server.Start(cfg.GRPC.Port); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
