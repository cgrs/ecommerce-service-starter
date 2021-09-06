package carts

import (
	"context"
	"fmt"

	"github.com/cgrs/ecommerce-service-starter/pkg/internal/storage"
)

type Repository interface {
	Create(ctx context.Context, c *Cart) error
	FindById(ctx context.Context, id string) *Cart
	UpdateItems(ctx context.Context, id string, items ItemMap) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	storage storage.Storage
}

func NewRepository(s storage.Storage) Repository {
	return &repository{s}
}

func (r *repository) Create(ctx context.Context, c *Cart) error {
	cart := r.FindById(ctx, c.ID)
	if cart != nil {
		return fmt.Errorf("cart with id %s already exists", c.ID)
	}
	return r.storage.Insert(ctx, c.ID, c)
}

func (r *repository) FindById(ctx context.Context, id string) *Cart {
	cart, _ := r.storage.Find(ctx, id).(*Cart)
	return cart
}

func (r *repository) UpdateItems(ctx context.Context, id string, items ItemMap) error {
	cart, err := r.findOrFail(ctx, id)
	if err != nil {
		return err
	}
	cart.Items = items
	return r.storage.Update(ctx, id, cart)
}

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.findOrFail(ctx, id)
	if err != nil {
		return err
	}
	return r.storage.Delete(ctx, id)
}

func (r *repository) findOrFail(ctx context.Context, id string) (*Cart, error) {
	c := r.FindById(ctx, id)
	if c == nil {
		return nil, fmt.Errorf("cart with id %s not found", id)
	}
	return c, nil
}
