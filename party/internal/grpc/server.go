package grpc

import (
	"fmt"
	"net"

	"party/internal/proto"
	"party/internal/service"

	"google.golang.org/grpc"
)

type Server struct {
	keygenService *service.KeygenService
}

func NewServer() *Server {
	return &Server{
		keygenService: service.NewKeygenService(),
	}
}

func (s *Server) Start(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterKeygenServiceServer(grpcServer, s.keygenService)

	fmt.Printf("gRPC server listening on :%d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
