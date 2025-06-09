package service

import (
	"context"

	"example.com/go-web-api-sample/ogen"
	"example.com/go-web-api-sample/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *ogen.ModelsUser) (*ogen.ModelsUser, error) {
	return s.repo.Create(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*ogen.ModelsUser, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]ogen.ModelsUser, error) {
	return s.repo.List(ctx)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}