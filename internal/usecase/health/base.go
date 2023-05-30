package health

import (
	"context"
	"lion/internal/model/health"
	"lion/internal/shared/config"
	"lion/internal/shared/intergration/lion"
	"net/http"
)

type UseCase interface {
	Readiness(ctx context.Context) interface{}
}

type healthUseCase struct {
	config      *config.Config
	lionRequest lion.Request
}

func NewHealthUseCase(cfg *config.Config, lionRequest lion.Request) UseCase {
	return &healthUseCase{
		config:      cfg,
		lionRequest: lionRequest,
	}
}

func (h *healthUseCase) Readiness(ctx context.Context) interface{} {
	lionAPICheck := h.checkIntegrationConnection(ctx, "lion")
	return map[string]*health.Readiness{
		"lion_connection": lionAPICheck,
	}
}

func (h *healthUseCase) checkIntegrationConnection(ctx context.Context, serviceName string) *health.Readiness {
	status := &health.Readiness{
		Status: http.StatusText(http.StatusOK),
	}
	switch serviceName {
	case "lion":
		err := h.lionRequest.Ping(ctx)
		if err != nil {
			status.Status = http.StatusText(http.StatusInternalServerError)
			status.Message = err.Error()
		}
	}
	return status
}
