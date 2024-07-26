package grpc

import (
	"log"
	"net"

	"gateway/internal/proto"

	"google.golang.org/grpc"
)

type KeygenServiceServer struct {
	proto.UnimplementedKeygenServiceServer
	FinishedKeys chan string
}

func NewKeygenServiceServer() *KeygenServiceServer {
	return &KeygenServiceServer{
		FinishedKeys: make(chan string, 100),
	}
}

// KeygenFinished는 Pod로부터 키 생성 완료 메시지를 스트리밍으로 받는 gRPC 메서드입니다.
func (s *KeygenServiceServer) KeygenFinished(stream proto.KeygenService_KeygenFinishedServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == grpc.ErrServerStopped {
				return nil
			}
			return err
		}
		s.FinishedKeys <- req.Publickey
	}
}

// StartGRPCServer는 gRPC 서버를 시작합니다.
func StartGRPCServer(server *KeygenServiceServer) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterKeygenServiceServer(grpcServer, server)

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
