package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer(address string, handler http.Handler) *http.Server {
	if address == "" {
		address = "localhost:3000"
	}

	return &http.Server{
		Addr:    address,
		Handler: handler,
	}
}

// Starts the server
func Start(s *http.Server) error {
	log.Printf("Server is listening on http://%s\n", s.Addr)
	return s.ListenAndServe()
}

type Server interface {
	Start() error
}

type ServerOptions func(e *gin.Engine) error
type server struct {
	address string
	engine  *gin.Engine
}

func New(address string, options ...ServerOptions) (Server, error) {
	e := gin.New()
	e.HandleMethodNotAllowed = true
	e.Use(gin.Logger(), gin.Recovery())
	for _, opt := range options {
		if err := opt(e); err != nil {
			return nil, err
		}
	}
	return &server{address, e}, nil
}

func (s *server) Start() error {
	return s.engine.Run(s.address)
}
