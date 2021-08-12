package server

import (
	"log"
	"net/http"
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
