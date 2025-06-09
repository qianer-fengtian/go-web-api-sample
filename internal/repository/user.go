package repository

import (
	"context"
	"errors"

	"example.com/go-web-api-sample/ogen"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	Create(ctx context.Context, user *ogen.ModelsUser) (*ogen.ModelsUser, error)
	GetByID(ctx context.Context, id int64) (*ogen.ModelsUser, error)
	List(ctx context.Context) ([]ogen.ModelsUser, error)
	Delete(ctx context.Context, id int64) error
}