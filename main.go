package main

import (
	"log"

	"github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/plugins/wrapper/validator/v3"
	"github.com/asim/go-micro/v3"
	ms "github.com/asim/go-micro/v3/server"
	"github.com/cgrs/ecommerce-service-starter/server"
)

func main() {
	s := server.New(
		micro.Server(grpc.NewServer(ms.Address(":50051"))),
		micro.Name("ecommerce"),
		micro.WrapHandler(validator.NewHandlerWrapper()),
	)
	log.Fatal(s.Start())
}
