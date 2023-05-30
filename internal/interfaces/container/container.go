package container

import (
	"github.com/go-playground/validator/v10"
	"lion/internal/shared/config"
	"lion/internal/shared/infrastructure/httpclient"
	"lion/internal/shared/infrastructure/soapclient"
	"lion/internal/shared/intergration/lion"
	"lion/internal/usecase/health"
	lionUseCase "lion/internal/usecase/lion"
	"time"
)

type Container struct {
	Config     *config.Config
	Validator  *validator.Validate
	SoapClient soapclient.SOAPClient

	LionRequest lion.Request

	HealthUseCase health.UseCase
	LionUseCase   lionUseCase.UseCase
}

func Setup() *Container {
	cfg := config.New()
	validate := validator.New()
	httpClient := httpclient.NewDefaultClient(time.Duration(cfg.App.HttpRequestTimeout)*time.Second, false)
	soapClient := soapclient.NewSOAPClient(httpClient)

	lionRequest := lion.NewRequest(cfg, httpClient, soapClient)

	healthUseCase := health.NewHealthUseCase(cfg, lionRequest)
	lionUseCase := lionUseCase.NewLionUseCase(cfg, lionRequest)

	return &Container{
		Config:        cfg,
		Validator:     validate,
		SoapClient:    soapClient,
		LionRequest:   lionRequest,
		HealthUseCase: healthUseCase,
		LionUseCase:   lionUseCase,
	}
}
