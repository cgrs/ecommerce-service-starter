package main

import (
	"github.com/cgrs/ecommerce-service-starter/middlewares"
	"github.com/cgrs/ecommerce-service-starter/server"
)

func main() {
	server.Start(
		server.CreateServer(
			"",
			middlewares.WithLogger(
				server.MainMux,
				nil,
			),
		),
	)
}
