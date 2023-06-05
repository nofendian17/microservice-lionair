package lion

import (
	"context"
	model "lion/internal/model/lion"
	"lion/internal/model/response"
	"lion/internal/shared/config"
	"lion/internal/shared/infrastructure/redisclient"
	"lion/internal/shared/intergration/lion"
)

type UseCase interface {
	SessionCreate(ctx context.Context, request model.SessionCreateRequest) (*response.Response, error)
	SessionClose(ctx context.Context, request model.SessionCloseRequest) (*response.Response, error)
	ScheduleWithoutSession(ctx context.Context, request model.ScheduleRequest) (*response.Response, error)
	ScheduleWithSession(ctx context.Context, request model.ScheduleRequest) (*response.Response, error)
	Pricing(ctx context.Context, request model.PricingRequest) (*response.Response, error)
}

type lionUseCase struct {
	config        *config.Config
	redisClient   redisclient.Redis
	lionRequest   lion.Request
	sessionCreate lion.SessionCreate
	sessionClose  lion.SessionClose
	flightMatrix  lion.FlightMatrix
}

func NewLionUseCase(
	cfg *config.Config,
	redisClient redisclient.Redis,
	lionRequest lion.Request,
	sessionCreate lion.SessionCreate,
	sessionClose lion.SessionClose,
	flightMatrix lion.FlightMatrix,
) UseCase {
	return &lionUseCase{
		config:        cfg,
		redisClient:   redisClient,
		lionRequest:   lionRequest,
		sessionCreate: sessionCreate,
		sessionClose:  sessionClose,
		flightMatrix:  flightMatrix,
	}
}
