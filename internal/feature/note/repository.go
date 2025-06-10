package note

import (
	"context"

	"example.com/go-web-api-sample/ogen"
)

type Repository interface {
	Create(ctx context.Context, note *ogen.ModelsNote) (*ogen.ModelsNote, error)
	GetByID(ctx context.Context, id int64) (*ogen.ModelsNote, error)
	List(ctx context.Context) ([]ogen.ModelsNote, error)
	Delete(ctx context.Context, id int64) error
}