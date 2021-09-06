package customers

import (
	"context"
	"fmt"

	"github.com/cgrs/ecommerce-service-starter/pkg/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Create(ctx context.Context, customer *Customer) error
	FindByUsername(ctx context.Context, username string) *Customer
	ComparePassword(ctx context.Context, username, password string) error
	UpdateEmail(ctx context.Context, username, newEmail string) error
	UpdatePassword(ctx context.Context, username, newPassword string) error
	UpdateRole(ctx context.Context, username string, becomeAdmin bool) error
}

type repository struct {
	storage storage.Storage
}

func NewRepository(s storage.Storage) Repository {
	return &repository{s}
}

func (r *repository) Create(ctx context.Context, cust *Customer) error {
	customer := r.FindByUsername(ctx, cust.Username)
	if customer != nil {
		return fmt.Errorf("customer with username %s already exists", cust.Username)
	}
	hashedPass, _ := computeHash(cust.Password)
	cust.Password = hashedPass
	return r.storage.Insert(ctx, cust.Username, cust)
}

func (r *repository) FindByUsername(ctx context.Context, username string) *Customer {
	customer, _ := r.storage.Find(ctx, username).(*Customer)
	return customer
}

func (r *repository) ComparePassword(ctx context.Context, username, password string) error {
	customer, err := r.findOrFail(ctx, username)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
}

func (r *repository) UpdateEmail(ctx context.Context, username, newEmail string) error {
	customer, err := r.findOrFail(ctx, username)
	if err != nil {
		return err
	}
	customer.Email = newEmail
	return r.storage.Update(ctx, username, customer)
}

func (r *repository) UpdatePassword(ctx context.Context, username, newPassword string) error {
	customer, err := r.findOrFail(ctx, username)
	if err != nil {
		return err
	}
	hashedPass, _ := computeHash(newPassword)
	customer.Password = hashedPass
	return r.storage.Update(ctx, username, customer)
}

func (r *repository) UpdateRole(ctx context.Context, username string, becomeAdmin bool) error {
	customer, err := r.findOrFail(ctx, username)
	if err != nil {
		return err
	}
	customer.Admin = becomeAdmin
	return r.storage.Update(ctx, username, customer)
}

func (r *repository) findOrFail(ctx context.Context, username string) (*Customer, error) {
	customer := r.FindByUsername(ctx, username)
	if customer == nil {
		return nil, fmt.Errorf("customer with username %s not found", username)
	}
	return customer, nil
}

func computeHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
