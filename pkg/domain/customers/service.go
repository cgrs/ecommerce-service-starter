package customers

import "context"

type Service interface {
	Create(ctx context.Context, username, password, email string, isAdmin bool) error
	Login(ctx context.Context, username, password string) error
	Update(ctx context.Context, username, password, email string) error
	BecomeAdmin(ctx context.Context, username string, isAdmin bool) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Create(ctx context.Context, username, password, email string, isAdmin bool) error {
	return s.repository.Create(ctx, &Customer{
		Username: username,
		Password: password,
		Email:    email,
		Admin:    isAdmin,
	})
}

func (s *service) Login(ctx context.Context, username, password string) error {
	return s.repository.ComparePassword(ctx, username, password)
}

func (s *service) Update(ctx context.Context, username, password, email string) (err error) {
	if password != "" {
		err = s.repository.UpdatePassword(ctx, username, password)
		if err != nil {
			return
		}
	}
	if email != "" {
		err = s.repository.UpdateEmail(ctx, username, email)
		if err != nil {
			return
		}
	}
	return
}

func (s *service) BecomeAdmin(ctx context.Context, username string, isAdmin bool) error {
	return s.repository.UpdateRole(ctx, username, isAdmin)
}
