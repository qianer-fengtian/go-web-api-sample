package note

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

func (s *Service) CreateNote(ctx context.Context, req *ogen.ModelsNote) (*ogen.ModelsNote, error) {
	return s.repo.Create(ctx, req)
}

func (s *Service) GetNote(ctx context.Context, id int64) (*ogen.ModelsNote, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) ListNotes(ctx context.Context) ([]ogen.ModelsNote, error) {
	return s.repo.List(ctx)
}

func (s *Service) DeleteNote(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}