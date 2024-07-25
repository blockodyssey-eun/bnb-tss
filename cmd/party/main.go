package main

import (
	"log"
	"net"

	"tss_project/internal/tss/keygen"
	"tss_project/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTSSServiceServer(grpcServer, &keygen.Server{})
	log.Printf("gRPC server listening on :8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
