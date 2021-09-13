package server

import (
	"github.com/asim/go-micro/v3"
	itemsv1 "github.com/cgrs/ecommerce-service-starter/items/v1"
	ordersv1 "github.com/cgrs/ecommerce-service-starter/orders/v1"
	"github.com/cgrs/ecommerce-service-starter/storage/memory"
)

type server struct {
	ms micro.Service
}

type Server interface {
	Start() error
}

func New(opts ...micro.Option) Server {
	ms := micro.NewService(opts...)
	ms.Init()
	s := &server{ms}
	s.setup()
	return s
}

func (s *server) Start() error {
	return s.ms.Run()
}

func (s *server) setup() {
	itemsv1.Register(s.ms.Server(), itemsv1.New(itemsv1.NewRepository(memory.New())))
	ordersv1.Register(s.ms.Server(), ordersv1.New(ordersv1.NewRepository(memory.New())))
}
