package user

import (
	"context"

	"example.com/go-web-api-sample/ogen"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(ctx context.Context, req *ogen.ModelsUser) (*ogen.ModelsUser, error) {
	return s.repo.Create(ctx, req)
}

func (s *Service) GetUser(ctx context.Context, id int64) (*ogen.ModelsUser, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) ListUsers(ctx context.Context) ([]ogen.ModelsUser, error) {
	return s.repo.List(ctx)
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}