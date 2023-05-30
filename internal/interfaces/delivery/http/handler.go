package http

import (
	"lion/internal/interfaces/container"
	"lion/internal/interfaces/delivery/http/health"
	"lion/internal/interfaces/delivery/http/lion"
)

type Handler struct {
	health *health.Handler
	lion   *lion.Handler
}

func NewHandler(container *container.Container) *Handler {
	return &Handler{
		health: health.NewHandler(container.Config, container.HealthUseCase),
		lion:   lion.NewHandler(container.Config, container.LionUseCase),
	}
}
