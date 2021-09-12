package v1

import (
	"context"

	pb "github.com/cgrs/ecommerce-service-starter/gen/proto/go/orders/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type orderService struct {
	repository Repository
	pb.UnimplementedOrdersServiceServer
}

func New(r Repository) pb.OrdersServiceServer {
	return &orderService{repository: r}
}

func (s *orderService) Create(ctx context.Context, r *pb.CreateOrderRequest) (*pb.Order, error) {
	i := r.GetOrder()
	i.Id = uuid.NewString()
	return s.repository.Create(ctx, i)
}

func (s *orderService) List(ctx context.Context, r *pb.ListOrderRequest) (*pb.ListResponse, error) {
	list := s.repository.Filter(ctx, r.Ids)
	return &pb.ListResponse{Orders: list}, nil
}

func (s *orderService) Find(ctx context.Context, r *pb.FindOrderRequest) (*pb.Order, error) {
	item := s.repository.Find(ctx, r.Id)
	if item == nil {
		return nil, &ErrNotFound{r.Id}
	}
	return item, nil
}

func (s *orderService) Update(ctx context.Context, r *pb.UpdateOrderRequest) (*pb.Order, error) {
	return s.repository.Update(ctx, r.Order)
}

func (s *orderService) Delete(ctx context.Context, r *pb.DeleteOrderRequest) (*emptypb.Empty, error) {
	return nil, s.repository.Delete(ctx, r.Id)
}

func Register(server *grpc.Server, impl pb.OrdersServiceServer) {
	server.RegisterService(&pb.OrdersService_ServiceDesc, impl)
}
