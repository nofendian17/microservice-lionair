package health

import (
	"lion/internal/shared/config"
	"lion/internal/usecase/health"
)

type Handler struct {
	config        *config.Config
	healthUseCase health.UseCase
}

func NewHandler(config *config.Config, healthUseCase health.UseCase) *Handler {
	return &Handler{
		config:        config,
		healthUseCase: healthUseCase,
	}
}
