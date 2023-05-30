package lion

import (
	"lion/internal/shared/config"
	lionUseCase "lion/internal/usecase/lion"
)

type Handler struct {
	config      *config.Config
	lionUseCase lionUseCase.UseCase
}

func NewHandler(config *config.Config, lionUseCase lionUseCase.UseCase) *Handler {
	return &Handler{
		config:      config,
		lionUseCase: lionUseCase,
	}
}
