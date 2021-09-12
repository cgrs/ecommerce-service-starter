package v1

import (
	"context"
	"fmt"

	pb "github.com/cgrs/ecommerce-service-starter/gen/proto/go/orders/v1"
	"github.com/cgrs/ecommerce-service-starter/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository interface {
	Create(context.Context, *pb.Order) (*pb.Order, error)
	List(context.Context) ([]*pb.Order, error)
	Find(context.Context, string) *pb.Order
	Filter(context.Context, []string) []*pb.Order
	Update(context.Context, *pb.Order) (*pb.Order, error)
	Delete(context.Context, string) error
}

type repository struct {
	storage storage.Storage
}

type ErrNotFound struct {
	id string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf(`order with id "%s" not found`, e.id)
}

func (e *ErrNotFound) GRPCStatus() *status.Status {
	return status.New(codes.NotFound, e.Error())
}

func NewRepository(s storage.Storage) Repository {
	return &repository{s}
}

func (r *repository) Create(ctx context.Context, order *pb.Order) (*pb.Order, error) {
	return order, r.storage.Insert(ctx, order.Id, order)
}

func (r *repository) List(ctx context.Context) ([]*pb.Order, error) {
	res := []*pb.Order{}
	r.storage.Range(ctx, func(key, value interface{}) bool {
		o, _ := value.(*pb.Order)
		res = append(res, o)
		return true
	})
	return res, nil
}

func (r *repository) Find(ctx context.Context, id string) *pb.Order {
	o, _ := r.storage.Find(ctx, id).(*pb.Order)
	return o
}

func (r *repository) Update(ctx context.Context, o *pb.Order) (*pb.Order, error) {
	order := r.Find(ctx, o.Id)
	if order == nil {
		return nil, &ErrNotFound{o.Id}
	}
	order.Lines = o.Lines
	order.Status = o.Status
	order.Total = o.Total
	return order, r.storage.Update(ctx, order.Id, order)
}

func (r *repository) Delete(ctx context.Context, id string) error {
	order := r.Find(ctx, id)
	if order == nil {
		return &ErrNotFound{id}
	}
	return r.storage.Delete(ctx, id)
}

func (r *repository) Filter(ctx context.Context, ids []string) []*pb.Order {
	result := []*pb.Order{}
	for _, id := range ids {
		o := r.Find(ctx, id)
		if o != nil {
			result = append(result, o)
		}
	}
	return result
}
