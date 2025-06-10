package user

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

func (h *Handler) CreateUser(ctx context.Context, req *ogen.ModelsUser) (*ogen.ModelsUser, error) {
	return h.service.CreateUser(ctx, req)
}

func (h *Handler) GetUser(ctx context.Context, params ogen.GetUserParams) (ogen.GetUserRes, error) {
	user, err := h.service.GetUser(ctx, params.ID)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return &ogen.GetUserNotFound{}, nil
		}
		return nil, err
	}
	return user, nil
}

func (h *Handler) ListUsers(ctx context.Context) ([]ogen.ModelsUser, error) {
	return h.service.ListUsers(ctx)
}

func (h *Handler) DeleteUser(ctx context.Context, params ogen.DeleteUserParams) (ogen.DeleteUserRes, error) {
	err := h.service.DeleteUser(ctx, params.ID)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return &ogen.DeleteUserNotFound{}, nil
		}
		return nil, err
	}
	return &ogen.DeleteUserNoContent{}, nil
}