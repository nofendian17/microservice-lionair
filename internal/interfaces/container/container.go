package container

import (
	"github.com/go-playground/validator/v10"
	"lion/internal/shared/config"
	"lion/internal/shared/infrastructure/httpclient"
	"lion/internal/shared/infrastructure/redisclient"
	"lion/internal/shared/infrastructure/soapclient"
	"lion/internal/shared/intergration/lion"
	"lion/internal/usecase/health"
	lionUseCase "lion/internal/usecase/lion"
	"time"
)

type Container struct {
	Config      *config.Config
	Validator   *validator.Validate
	SoapClient  soapclient.SOAPClient
	RedisClient redisclient.Redis
	LionRequest lion.Request

	HealthUseCase health.UseCase
	LionUseCase   lionUseCase.UseCase
}

func Setup() *Container {
	cfg := config.New()
	validate := validator.New()
	httpClient := httpclient.NewDefaultClient(time.Duration(cfg.App.HttpRequestTimeout)*time.Second, false)
	soapClient := soapclient.NewSOAPClient(httpClient)
	redisClient := redisclient.NewDefaultClient(cfg)

	lionRequest := lion.NewRequest(cfg, httpClient, soapClient)

	healthUseCase := health.NewHealthUseCase(cfg, lionRequest)
	sessionCreate := lion.NewSessionCreate(lionRequest)
	sessionClose := lion.NewSessionClose(lionRequest)
	flightMatrix := lion.NewFlightMatrix(lionRequest)

	lionUseCase := lionUseCase.NewLionUseCase(
		cfg,
		redisClient,
		lionRequest,
		sessionCreate,
		sessionClose,
		flightMatrix,
	)

	return &Container{
		Config:        cfg,
		Validator:     validate,
		SoapClient:    soapClient,
		RedisClient:   redisClient,
		LionRequest:   lionRequest,
		HealthUseCase: healthUseCase,
		LionUseCase:   lionUseCase,
	}
}
