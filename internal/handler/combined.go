package handler

import (
	"context"

	"example.com/go-web-api-sample/ogen"
	"example.com/go-web-api-sample/internal/service"
)

// CombinedHandler implements all ogen Handler interface methods by delegating to specialized handlers
type CombinedHandler struct {
	services *service.ServiceContainer
	// 専用ハンドラーは遅延初期化または直接サービス呼び出し
	noteHandler *NoteHandler
	userHandler *UserHandler
}

func NewCombinedHandler(services *service.ServiceContainer) *CombinedHandler {
	return &CombinedHandler{
		services:    services,
		noteHandler: NewNoteHandler(services.NoteService),
		userHandler: NewUserHandler(services.UserService),
	}
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