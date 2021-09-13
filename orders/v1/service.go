package v1

import (
	"context"

	"github.com/asim/go-micro/v3/server"
	pb "github.com/cgrs/ecommerce-service-starter/gen/proto/go/orders/v1"
	"github.com/google/uuid"
)

type orderService struct {
	repository Repository
}

func New(r Repository) pb.OrdersHandler {
	return &orderService{repository: r}
}

func (s *orderService) Create(ctx context.Context, r *pb.CreateOrderRequest, or *pb.CreateOrderResponse) (err error) {
	o := r.GetOrder()
	o.Id = uuid.NewString()
	order, err := s.repository.Create(ctx, o)
	or.Order = order
	return
}

func (s *orderService) List(ctx context.Context, r *pb.ListOrderRequest, lr *pb.ListOrderResponse) (err error) {
	list, err := s.repository.List(ctx)
	lr.Orders = list
	return
}

func (s *orderService) Filter(ctx context.Context, r *pb.FilterOrderRequest, fr *pb.FilterOrderResponse) (err error) {
	fr.Orders = s.repository.Filter(ctx, r.GetIds())
	return
}

func (s *orderService) Find(ctx context.Context, r *pb.FindOrderRequest, fr *pb.FindOrderResponse) (err error) {
	o := s.repository.Find(ctx, r.Id)
	if o == nil {
		err = &ErrNotFound{r.Id}
		return
	}
	fr.Order = o
	return
}

func (s *orderService) Update(ctx context.Context, r *pb.UpdateOrderRequest, ur *pb.UpdateOrderResponse) (err error) {
	or, err := s.repository.Update(ctx, r.Order)
	ur.Order = or
	return
}

func (s *orderService) Delete(ctx context.Context, r *pb.DeleteOrderRequest, dr *pb.DeleteOrderResponse) error {
	return s.repository.Delete(ctx, r.Id)
}

func Register(s server.Server, impl pb.OrdersHandler) error {
	return pb.RegisterOrdersHandler(s, impl)
}
