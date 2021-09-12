package server

import (
	"net"

	itemsv1 "github.com/cgrs/ecommerce-service-starter/items/v1"
	ordersv1 "github.com/cgrs/ecommerce-service-starter/orders/v1"
	"github.com/cgrs/ecommerce-service-starter/storage/memory"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

type server struct {
	gs *grpc.Server
}

type Server interface {
	Start(addr string) error
}

func New() Server {
	gs := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_recovery.StreamServerInterceptor(),
				grpc_validator.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
				grpc_validator.UnaryServerInterceptor(),
			),
		),
	)
	s := &server{gs}
	s.setup()
	return s
}

func (s *server) Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.gs.Serve(lis)
}

func (s *server) setup() {
	itemsv1.Register(s.gs, itemsv1.New(itemsv1.NewRepository(memory.New())))
	ordersv1.Register(s.gs, ordersv1.New(ordersv1.NewRepository(memory.New())))
}
