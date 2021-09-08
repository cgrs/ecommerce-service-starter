package main

import (
	"log"

	"github.com/cgrs/ecommerce-service-starter/server"
)

func main() {
	s := server.New()
	if err := s.Start(":50051"); err != nil {
		log.Fatal(err)
	}
}
