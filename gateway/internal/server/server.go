package server

import (
	grpcClient "gateway/internal/grpc"
	"gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router       *gin.Engine
	keygenServer *grpcClient.KeygenServiceServer
}

func NewServer(keygenServer *grpcClient.KeygenServiceServer) *Server {
	router := gin.Default()
	server := &Server{
		router:       router,
		keygenServer: keygenServer,
	}

	server.routes()

	return server
}

func (s *Server) routes() {
	s.router.POST("/keygen", handler.Keygen(s.keygenServer))
}

func (s *Server) Run(addr string) {
	s.router.Run(addr)
}
