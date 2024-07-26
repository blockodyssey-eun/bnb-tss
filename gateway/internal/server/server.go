package server

import (
	"fmt"

	"gateway/internal/config"
	"gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	server := &Server{router: router}

	server.routes()

	return server
}

func (s *Server) routes() {
	s.router.POST("/keygen", handler.Keygen)
}

func (s *Server) Run(addr string) {
	cfg := config.Get()
	s.router.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
