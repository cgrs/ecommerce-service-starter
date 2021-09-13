package v1

import (
	"context"

	"github.com/asim/go-micro/v3/server"
	pb "github.com/cgrs/ecommerce-service-starter/gen/proto/go/items/v1"
	"github.com/google/uuid"
)

type itemService struct {
	repository Repository
}

func New(r Repository) pb.ItemsHandler {
	return &itemService{repository: r}
}

func (s *itemService) Create(ctx context.Context, r *pb.CreateItemRequest, cr *pb.CreateItemResponse) (err error) {
	i := r.GetItem()
	i.Id = uuid.NewString()
	it, err := s.repository.Create(ctx, i)
	cr.Item = it
	return
}

func (s *itemService) List(ctx context.Context, r *pb.ListItemRequest, lr *pb.ListItemResponse) (err error) {
	lr.Items, err = s.repository.List(ctx)
	return
}

func (s *itemService) Filter(ctx context.Context, r *pb.FilterItemRequest, lr *pb.FilterItemResponse) (err error) {
	lr.Items = s.repository.Filter(ctx, r.GetIds())
	return
}

func (s *itemService) Find(ctx context.Context, r *pb.FindItemRequest, fr *pb.FindItemResponse) (err error) {
	it := s.repository.Find(ctx, r.Id)
	if it == nil {
		err = &ErrNotFound{r.Id}
		return
	}
	fr.Item = it
	return
}

func (s *itemService) Update(ctx context.Context, r *pb.UpdateItemRequest, ur *pb.UpdateItemResponse) (err error) {
	it, err := s.repository.Update(ctx, r.Item)
	ur.Item = it
	return err
}

func (s *itemService) Delete(ctx context.Context, r *pb.DeleteItemRequest, dr *pb.DeleteItemResponse) error {
	return s.repository.Delete(ctx, r.Id)
}

func Register(s server.Server, impl pb.ItemsHandler) error {
	return pb.RegisterItemsHandler(s, impl)
}
