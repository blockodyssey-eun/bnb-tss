package party

import (
	"log"
	"net"
	"tss_project/handler/party"
	"tss_project/proto"

	"google.golang.org/grpc"
)

type TSSService struct {
	proto.UnimplementedTSSServiceServer
	keygenHandler *party.KeygenHandler
}

func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	keygenHandler := &party.KeygenHandler{}
	tssService := &TSSService{
		keygenHandler: keygenHandler,
	}
	proto.RegisterTSSServiceServer(grpcServer, tssService)

	return grpcServer
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := NewServer()
	log.Println("Starting Party server on :9090")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
