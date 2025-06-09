package handler

import (
	"context"
	"errors"

	"example.com/go-web-api-sample/ogen"
	"example.com/go-web-api-sample/internal/repository"
	"example.com/go-web-api-sample/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *ogen.ModelsUser) (*ogen.ModelsUser, error) {
	return h.userService.CreateUser(ctx, req)
}

func (h *UserHandler) GetUser(ctx context.Context, params ogen.GetUserParams) (ogen.GetUserRes, error) {
	user, err := h.userService.GetUser(ctx, params.ID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return &ogen.GetUserNotFound{}, nil
		}
		return nil, err
	}
	return user, nil
}

func (h *UserHandler) ListUsers(ctx context.Context) ([]ogen.ModelsUser, error) {
	return h.userService.ListUsers(ctx)
}

func (h *UserHandler) DeleteUser(ctx context.Context, params ogen.DeleteUserParams) (ogen.DeleteUserRes, error) {
	err := h.userService.DeleteUser(ctx, params.ID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return &ogen.DeleteUserNotFound{}, nil
		}
		return nil, err
	}
	return &ogen.DeleteUserNoContent{}, nil
}