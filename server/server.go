package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	g *gin.Engine
	addr string
}

func CreateServer(address string) *Server {
	if address == "" {
		address = "localhost:3000"
	}

	gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	AddItemRoutes(g.Group("/api"))

	return &Server{g,address}
}

func (s *Server) Start() error {
	return s.g.Run(s.addr)
}