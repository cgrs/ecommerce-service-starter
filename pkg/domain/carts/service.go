package carts

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context) (*Cart, error)
	FindById(ctx context.Context, id uuid.UUID) *Cart
	Update(ctx context.Context, id uuid.UUID, items ItemMap) error
	Empty(ctx context.Context, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Create(ctx context.Context) (*Cart, error) {
	c := &Cart{
		ID: uuid.NewString(),
	}
	return c, s.repository.Create(ctx, c)
}

func (s *service) FindById(ctx context.Context, id uuid.UUID) *Cart {
	return s.repository.FindById(ctx, id.String())
}

func (s *service) Update(ctx context.Context, id uuid.UUID, items ItemMap) error {
	return s.repository.UpdateItems(ctx, id.String(), items)
}

func (s *service) Empty(ctx context.Context, id uuid.UUID) error {
	return s.repository.UpdateItems(ctx, id.String(), ItemMap{})
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id.String())
}
