package main

import (
	"context"

	"example.com/go-web-api-sample/ogen"
	"example.com/go-web-api-sample/internal/feature/note"
	"example.com/go-web-api-sample/internal/feature/user"
	"example.com/go-web-api-sample/internal/feature/healthcheck"
)

// CombinedHandler implements all ogen Handler interface methods by delegating to specialized handlers
type CombinedHandler struct {
	noteHandler *note.Handler
	userHandler *user.Handler
	healthCheckHandler *healthcheck.Handler
}

// Note operations - delegated to NoteHandler
func (h *CombinedHandler) CreateNote(ctx context.Context, req *ogen.ModelsNote) (*ogen.ModelsNote, error) {
	return h.noteHandler.CreateNote(ctx, req)
}

func (h *CombinedHandler) GetNote(ctx context.Context, params ogen.GetNoteParams) (ogen.GetNoteRes, error) {
	return h.noteHandler.GetNote(ctx, params)
}

func (h *CombinedHandler) ListNotes(ctx context.Context) ([]ogen.ModelsNote, error) {
	return h.noteHandler.ListNotes(ctx)
}

func (h *CombinedHandler) DeleteNote(ctx context.Context, params ogen.DeleteNoteParams) (ogen.DeleteNoteRes, error) {
	return h.noteHandler.DeleteNote(ctx, params)
}

// User operations - delegated to UserHandler
func (h *CombinedHandler) CreateUser(ctx context.Context, req *ogen.ModelsUser) (*ogen.ModelsUser, error) {
	return h.userHandler.CreateUser(ctx, req)
}

func (h *CombinedHandler) GetUser(ctx context.Context, params ogen.GetUserParams) (ogen.GetUserRes, error) {
	return h.userHandler.GetUser(ctx, params)
}

func (h *CombinedHandler) ListUsers(ctx context.Context) ([]ogen.ModelsUser, error) {
	return h.userHandler.ListUsers(ctx)
}

func (h *CombinedHandler) DeleteUser(ctx context.Context, params ogen.DeleteUserParams) (ogen.DeleteUserRes, error) {
	return h.userHandler.DeleteUser(ctx, params)
}

// Health check operations - delegated to HealthCheckHandler
func (h *CombinedHandler) HealthCheckGetHealthCheck(ctx context.Context) (*ogen.RoutesHealthCheckResponse, error) {
	return h.healthCheckHandler.HealthCheckGetHealthCheck(ctx)
}