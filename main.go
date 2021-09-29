package main

import (
	"github.com/cgrs/ecommerce-service-starter/server"
	"log"
)

func main() {
	log.Fatal(server.CreateServer("").Start())
}
