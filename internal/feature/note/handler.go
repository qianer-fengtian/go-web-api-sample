package note

import (
	"context"
	"errors"

	"example.com/go-web-api-sample/ogen"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateNote(ctx context.Context, req *ogen.ModelsNote) (*ogen.ModelsNote, error) {
	return h.service.CreateNote(ctx, req)
}

func (h *Handler) GetNote(ctx context.Context, params ogen.GetNoteParams) (ogen.GetNoteRes, error) {
	note, err := h.service.GetNote(ctx, params.ID)
	if err != nil {
		if errors.Is(err, ErrNoteNotFound) {
			return &ogen.GetNoteNotFound{}, nil
		}
		return nil, err
	}
	return note, nil
}

func (h *Handler) ListNotes(ctx context.Context) ([]ogen.ModelsNote, error) {
	return h.service.ListNotes(ctx)
}

func (h *Handler) DeleteNote(ctx context.Context, params ogen.DeleteNoteParams) (ogen.DeleteNoteRes, error) {
	err := h.service.DeleteNote(ctx, params.ID)
	if err != nil {
		if errors.Is(err, ErrNoteNotFound) {
			return &ogen.DeleteNoteNotFound{}, nil
		}
		return nil, err
	}
	return &ogen.DeleteNoteNoContent{}, nil
}