package lion

import (
	"context"
	model "lion/internal/model/lion"
	"lion/internal/model/response"
	"net/http"
)

func (l *lionUseCase) ScheduleWithSession(ctx context.Context, request model.ScheduleRequest) (*response.Response, error) {
	binarySecurityToken, err := l.sessionCreate.GetBinarySecurityToken(ctx, request.ConversationID, l.config.Integration.Credentials.Username, l.config.Integration.Credentials.Password, l.config.Integration.Credentials.Organization)
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
		request.PassengerTypeQuantityADT,
		request.PassengerTypeQuantityCNN,
		request.PassengerTypeQuantityINF,
	)
	if err != nil {
		return nil, err
	}

	err = l.sessionClose.Logout(ctx, request.ConversationID, binarySecurityToken)
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
		request.PassengerTypeQuantityADT,
		request.PassengerTypeQuantityCNN,
		request.PassengerTypeQuantityINF,
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
