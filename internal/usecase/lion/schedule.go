package lion

import (
	"context"
	model "lion/internal/model/lion"
	"lion/internal/model/response"
	"net/http"
)

func (l *lionUseCase) ScheduleWithSession(ctx context.Context, request model.ScheduleRequest) (*response.Response, error) {
	binarySecurityToken, err := l.redisClient.Get(ctx, request.SessionID)
	if err != nil {
		return nil, err
	}
	schedule, err := l.flightMatrix.GetFlightScheduleWithSession(
		ctx,
		request.ConversationID,
		binarySecurityToken,
		l.config.Integration.Target,
		l.config.Integration.Service.FlightMatrix.Version,
		request.Direction.String(),
		request.DepartureDateTime,
		request.ArrivalDateTime,
		request.DepartureAirPort,
		request.ArrivalAirPort,
		request.QuantityADT,
		request.QuantityCNN,
		request.QuantityINF,
	)
	if err != nil {
		return nil, err
	}
	return &response.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    schedule,
	}, nil
}

func (l *lionUseCase) ScheduleWithoutSession(ctx context.Context, request model.ScheduleRequest) (*response.Response, error) {
	schedule, err := l.flightMatrix.GetFlightScheduleWithOutSession(
		ctx,
		request.ConversationID,
		l.config.Integration.Target,
		l.config.Integration.Service.FlightMatrix.Version,
		request.Direction.String(),
		request.DepartureDateTime,
		request.ArrivalDateTime,
		request.DepartureAirPort,
		request.ArrivalAirPort,
		request.QuantityADT,
		request.QuantityCNN,
		request.QuantityINF,
	)
	if err != nil {
		return nil, err
	}

	return &response.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    schedule,
	}, nil
}
