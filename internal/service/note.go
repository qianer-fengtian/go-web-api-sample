package service

import (
	"context"

	"example.com/go-web-api-sample/ogen"
	"example.com/go-web-api-sample/internal/repository"
)

type NoteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

func (s *NoteService) CreateNote(ctx context.Context, req *ogen.ModelsNote) (*ogen.ModelsNote, error) {
	return s.repo.Create(ctx, req)
}

func (s *NoteService) GetNote(ctx context.Context, id int64) (*ogen.ModelsNote, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *NoteService) ListNotes(ctx context.Context) ([]ogen.ModelsNote, error) {
	return s.repo.List(ctx)
}

func (s *NoteService) DeleteNote(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}