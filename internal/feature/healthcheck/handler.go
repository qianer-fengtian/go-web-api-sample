package healthcheck

import (
	"context"
	"time"

	"example.com/go-web-api-sample/ogen"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheckGetHealthCheck(ctx context.Context) (*ogen.RoutesHealthCheckResponse, error) {
	return &ogen.RoutesHealthCheckResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   ogen.NewOptString("1.0.0"),
	}, nil
}