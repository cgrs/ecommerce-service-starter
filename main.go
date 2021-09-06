package main

import (
	"github.com/cgrs/ecommerce-service-starter/server"
)

func main() {
	s, _ := server.New(":3000", server.Router)
	s.Start()
}
