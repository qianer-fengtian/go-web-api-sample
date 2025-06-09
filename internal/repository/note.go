package repository

import (
	"context"
	"errors"

	"example.com/go-web-api-sample/ogen"
)

var (
	ErrNoteNotFound = errors.New("note not found")
)

type NoteRepository interface {
	Create(ctx context.Context, note *ogen.ModelsNote) (*ogen.ModelsNote, error)
	GetByID(ctx context.Context, id int64) (*ogen.ModelsNote, error)
	List(ctx context.Context) ([]ogen.ModelsNote, error)
	Delete(ctx context.Context, id int64) error
}