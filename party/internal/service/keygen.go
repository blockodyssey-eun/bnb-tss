package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"party/internal/config"
	"party/internal/proto"

	"google.golang.org/grpc"
)

type KeygenService struct {
	proto.UnimplementedKeygenServiceServer
}

func NewKeygenService() *KeygenService {
	return &KeygenService{}
}

func (s *KeygenService) GenerateKey(ctx context.Context, req *proto.KeygenRequest) (*proto.KeygenResponse, error) {
	// 키 생성 작업 시뮬레이션
	time.Sleep(2 * time.Second)

	publicKey := fmt.Sprintf("generated_key_n%d_m%d", req.N, req.M)

	// KeygenFinished 메시지를 Gateway로 보냅니다.
	go s.sendKeygenFinished(publicKey)

	return &proto.KeygenResponse{Publickey: publicKey}, nil
}

func (s *KeygenService) sendKeygenFinished(publicKey string) {
	cfg := config.Get()
	gatewayAddress := fmt.Sprintf("%s:%d", cfg.Gateway.Host, cfg.Gateway.Port)

	// Gateway와의 gRPC 연결 설정
	conn, err := grpc.Dial(gatewayAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to Gateway: %v", err)
		return
	}
	defer conn.Close()

	client := proto.NewKeygenServiceClient(conn)

	// KeygenFinished 스트림 시작
	stream, err := client.KeygenFinished(context.Background())
	if err != nil {
		log.Printf("Error creating KeygenFinished stream: %v", err)
		return
	}

	// KeygenFinished 메시지 전송
	err = stream.Send(&proto.KeygenFinishedRequest{
		Publickey: publicKey,
	})
	if err != nil {
		log.Printf("Error sending KeygenFinished message: %v", err)
		return
	}

	// 스트림 종료 및 응답 수신
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error receiving KeygenFinished response: %v", err)
		return
	}

	log.Printf("KeygenFinished message sent successfully. Gateway response: %s", reply.Message)
}
