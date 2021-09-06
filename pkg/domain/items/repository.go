package items

import (
	"context"
	"fmt"
	"net/url"

	"github.com/cgrs/ecommerce-service-starter/pkg/internal/storage"
)

type Repository interface {
	Create(ctx context.Context, item *Item) error
	FindById(ctx context.Context, id string) *Item
	List(ctx context.Context) []*Item
	UpdateName(ctx context.Context, id, name string) error
	UpdateDescription(ctx context.Context, id, description string) error
	UpdateImage(ctx context.Context, id string, image *url.URL) error
	UpdatePrice(ctx context.Context, id string, price float64) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	storage storage.Storage
}

func NewRepository(s storage.Storage) Repository {
	return &repository{s}
}

func (r *repository) Create(ctx context.Context, i *Item) error {
	item := r.FindById(ctx, i.ID)
	if item != nil {
		return fmt.Errorf("item with id %s already exists", i.ID)
	}
	return r.storage.Insert(ctx, i.ID, i)
}

func (r *repository) FindById(ctx context.Context, id string) *Item {
	item, _ := r.storage.Find(ctx, id).(*Item)
	return item
}

func (r *repository) List(ctx context.Context) []*Item {
	result := r.storage.FetchAll(ctx)
	items := make([]*Item, len(result))
	i := 0
	for _, v := range result {
		items[i] = v.(*Item)
		i++
	}
	return items
}

func (r *repository) UpdateName(ctx context.Context, id, name string) error {
	item, err := r.findOrFail(ctx, id)
	if err != nil {
		return err
	}
	item.Name = name
	return r.storage.Update(ctx, id, item)
}

func (r *repository) UpdateDescription(ctx context.Context, id, description string) error {
	item, err := r.findOrFail(ctx, id)
	if err != nil {
		return err
	}
	item.Description = description
	return r.storage.Update(ctx, id, item)
}

func (r *repository) UpdateImage(ctx context.Context, id string, image *url.URL) error {
	item, err := r.findOrFail(ctx, id)
	if err != nil {
		return err
	}
	item.Image = image
	return r.storage.Update(ctx, id, item)
}

func (r *repository) UpdatePrice(ctx context.Context, id string, price float64) error {
	item, err := r.findOrFail(ctx, id)
	if err != nil {
		return err
	}
	item.UnitPrice = price
	return r.storage.Update(ctx, id, item)
}

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.findOrFail(ctx, id)
	if err != nil {
		return err
	}
	return r.storage.Delete(ctx, id)
}

func (r *repository) findOrFail(ctx context.Context, id string) (*Item, error) {
	item := r.FindById(ctx, id)
	if item == nil {
		return nil, fmt.Errorf("item with id %s not found", id)
	}
	return item, nil
}
