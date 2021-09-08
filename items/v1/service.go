package v1

import (
	"context"

	pb "github.com/cgrs/ecommerce-service-starter/gen/proto/go/items/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type itemService struct {
	repository Repository
	pb.UnimplementedItemsServiceServer
}

func New(r Repository) pb.ItemsServiceServer {
	return &itemService{repository: r}
}

func (s *itemService) Create(ctx context.Context, r *pb.CreateItemRequest) (*pb.Item, error) {
	i := r.GetItem()
	i.Id = uuid.NewString()
	return s.repository.Create(ctx, i)
}

func (s *itemService) List(ctx context.Context, r *emptypb.Empty) (*pb.ListResponse, error) {
	list, err := s.repository.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error: %s", err.Error())
	}
	return &pb.ListResponse{Items: list}, nil
}

func (s *itemService) Find(ctx context.Context, r *pb.FindItemRequest) (*pb.Item, error) {
	item := s.repository.Find(ctx, r.Id)
	if item == nil {
		return nil, &ErrNotFound{r.Id}
	}
	return item, nil
}

func (s *itemService) Update(ctx context.Context, r *pb.UpdateItemRequest) (*pb.Item, error) {
	return s.repository.Update(ctx, r.Item)
}

func (s *itemService) Delete(ctx context.Context, r *pb.DeleteItemRequest) (*emptypb.Empty, error) {
	return nil, s.repository.Delete(ctx, r.Id)
}

func Register(server *grpc.Server, impl pb.ItemsServiceServer) {
	server.RegisterService(&pb.ItemsService_ServiceDesc, impl)
}
