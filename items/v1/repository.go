package v1

import (
	"context"
	"fmt"

	pb "github.com/cgrs/ecommerce-service-starter/gen/proto/go/items/v1"
	"github.com/cgrs/ecommerce-service-starter/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository interface {
	Create(context.Context, *pb.Item) (*pb.Item, error)
	List(context.Context) ([]*pb.Item, error)
	Filter(context.Context, []string) []*pb.Item
	Find(context.Context, string) *pb.Item
	Update(context.Context, *pb.Item) (*pb.Item, error)
	Delete(context.Context, string) error
}

type repository struct {
	storage storage.Storage
}

type ErrNotFound struct {
	id string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf(`item with id "%s" not found`, e.id)
}

func (e *ErrNotFound) GRPCStatus() *status.Status {
	return status.New(codes.NotFound, e.Error())
}

func NewRepository(s storage.Storage) Repository {
	return &repository{s}
}

func (r *repository) Create(ctx context.Context, item *pb.Item) (*pb.Item, error) {
	return item, r.storage.Insert(ctx, item.Id, item)
}

func (r *repository) List(ctx context.Context) ([]*pb.Item, error) {
	res := []*pb.Item{}
	r.storage.Range(ctx, func(key, value interface{}) bool {
		i, _ := value.(*pb.Item)
		res = append(res, i)
		return true
	})
	return res, nil
}

func (r *repository) Find(ctx context.Context, id string) *pb.Item {
	i, _ := r.storage.Find(ctx, id).(*pb.Item)
	return i
}

func (r *repository) Update(ctx context.Context, i *pb.Item) (*pb.Item, error) {
	item := r.Find(ctx, i.Id)
	if item == nil {
		return nil, &ErrNotFound{i.Id}
	}
	item.Name = i.Name
	item.Description = i.Description
	item.UnitPrice = i.UnitPrice
	return item, r.storage.Update(ctx, item.Id, item)
}

func (r *repository) Delete(ctx context.Context, id string) error {
	item := r.Find(ctx, id)
	if item == nil {
		return &ErrNotFound{id}
	}
	return r.storage.Delete(ctx, id)
}

func (r *repository) Filter(ctx context.Context, ids []string) []*pb.Item {
	result := []*pb.Item{}
	for _, id := range ids {
		i := r.Find(ctx, id)
		if i != nil {
			result = append(result, i)
		}
	}
	return result
}
