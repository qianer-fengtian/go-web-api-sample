package handler

import (
	"context"
	"errors"

	"example.com/go-web-api-sample/ogen"
	"example.com/go-web-api-sample/internal/repository"
	"example.com/go-web-api-sample/internal/service"
)

type NoteHandler struct {
	noteService *service.NoteService
}

func NewNoteHandler(noteService *service.NoteService) *NoteHandler {
	return &NoteHandler{
		noteService: noteService,
	}
}

func (h *NoteHandler) CreateNote(ctx context.Context, req *ogen.ModelsNote) (*ogen.ModelsNote, error) {
	return h.noteService.CreateNote(ctx, req)
}

func (h *NoteHandler) GetNote(ctx context.Context, params ogen.GetNoteParams) (ogen.GetNoteRes, error) {
	note, err := h.noteService.GetNote(ctx, params.ID)
	if err != nil {
		if errors.Is(err, repository.ErrNoteNotFound) {
			return &ogen.GetNoteNotFound{}, nil
		}
		return nil, err
	}
	return note, nil
}

func (h *NoteHandler) ListNotes(ctx context.Context) ([]ogen.ModelsNote, error) {
	return h.noteService.ListNotes(ctx)
}

func (h *NoteHandler) DeleteNote(ctx context.Context, params ogen.DeleteNoteParams) (ogen.DeleteNoteRes, error) {
	err := h.noteService.DeleteNote(ctx, params.ID)
	if err != nil {
		if errors.Is(err, repository.ErrNoteNotFound) {
			return &ogen.DeleteNoteNotFound{}, nil
		}
		return nil, err
	}
	return &ogen.DeleteNoteNoContent{}, nil
}