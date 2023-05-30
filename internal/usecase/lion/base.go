package lion

import (
	"context"
	"encoding/xml"
	model "lion/internal/model/lion"
	"lion/internal/model/response"
	wsFlightMatrixRequest "lion/internal/model/ws/request/flight_matrix"
	wsSessionCloseRequest "lion/internal/model/ws/request/session_close"
	wsSessionCreateRequest "lion/internal/model/ws/request/session_create"
	wsFlightMatrixResponse "lion/internal/model/ws/response/flight_matrix"
	wsSessionCreateResponse "lion/internal/model/ws/response/session_create"
	"lion/internal/shared/config"
	"lion/internal/shared/intergration/lion"
	"net/http"
)

type UseCase interface {
	ScheduleWithoutSession(ctx context.Context, request model.ScheduleRequest) (*response.Response, error)
	ScheduleWithSession(ctx context.Context, request model.ScheduleRequest) (*response.Response, error)
}

type lionUseCase struct {
	config      *config.Config
	lionRequest lion.Request
}

func NewLionUseCase(cfg *config.Config, lionRequest lion.Request) UseCase {
	return &lionUseCase{
		config:      cfg,
		lionRequest: lionRequest,
	}
}

func (l *lionUseCase) ScheduleWithSession(ctx context.Context, request model.ScheduleRequest) (*response.Response, error) {
	sessionCreateRequest := &wsSessionCreateRequest.SessionCreateXMLRequest{}
	sessionCreateXMLRequest, err := sessionCreateRequest.NewSessionCreate(request.ConversationID, l.config.Integration.Credentials.Username, l.config.Integration.Credentials.Password, l.config.Integration.Credentials.Organization)
	if err != nil {
		return nil, err
	}

	sessionCreateResponse, err := l.lionRequest.SessionCreate(ctx, sessionCreateXMLRequest)
	if err != nil {
		return nil, err
	}

	sessionCreateXMLResponse := &wsSessionCreateResponse.SessionCreateXMLResponse{}
	err = xml.Unmarshal([]byte(sessionCreateResponse), sessionCreateXMLResponse)
	if err != nil {
		return nil, err
	}

	binarySecurityToken, err := sessionCreateXMLResponse.GetBinarySecurityToken()
	if err != nil {
		return nil, err
	}

	flightMatrixRequest := &wsFlightMatrixRequest.FlightMatrixXMLRequest{}
	flightMatrixXMLRequest, err := flightMatrixRequest.NewFlightMatrix(request.ConversationID, binarySecurityToken, l.config.Integration.Target, l.config.Integration.Service.FlightMatrix.Version, request.Direction.String(), request.DepartureDateTime, request.ArrivalDateTime, request.DepartureAirPort, request.ArrivalAirPort, request.PassengerTypeQuantityADT, request.PassengerTypeQuantityCNN, request.PassengerTypeQuantityINF)
	if err != nil {
		return nil, err
	}

	flightMatrixResponse, err := l.lionRequest.FlightMatrix(ctx, flightMatrixXMLRequest)
	if err != nil {
		return nil, err
	}

	flightMatrixXMLResponse := &wsFlightMatrixResponse.FlightMatrixXMLResponse{}
	err = xml.Unmarshal([]byte(flightMatrixResponse), flightMatrixXMLResponse)
	if err != nil {
		return nil, err
	}

	additionalFields := map[string]string{
		"FLIGHT_CLASS": request.CabinClass,
	}
	flightMatrixXMLResponse.SetAdditionalField(additionalFields)

	data, err := flightMatrixXMLResponse.ToJSON()
	if err != nil {
		return nil, err
	}

	sessionCloseRequest := &wsSessionCloseRequest.SessionCloseXMLRequest{}
	sessionCloseXMLRequest, err := sessionCloseRequest.NewSessionClose(request.ConversationID, binarySecurityToken)
	if err != nil {
		return nil, err
	}

	_, err = l.lionRequest.SessionClose(ctx, sessionCloseXMLRequest)
	if err != nil {
		return nil, err
	}

	return &response.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}, nil
}

func (l *lionUseCase) ScheduleWithoutSession(ctx context.Context, request model.ScheduleRequest) (*response.Response, error) {
	flightMatrixRequest := &wsFlightMatrixRequest.FlightMatrixXMLRequest{}
	flightMatrixXMLRequest, err := flightMatrixRequest.NewFlightMatrix2(request.ConversationID, l.config.Integration.Target, l.config.Integration.Service.FlightMatrix.Version, request.Direction.String(), request.DepartureDateTime, request.ArrivalDateTime, request.DepartureAirPort, request.ArrivalAirPort, request.PassengerTypeQuantityADT, request.PassengerTypeQuantityCNN, request.PassengerTypeQuantityINF)
	if err != nil {
		return nil, err
	}

	flightMatrixResponse, err := l.lionRequest.FlightMatrix(ctx, flightMatrixXMLRequest)
	if err != nil {
		return nil, err
	}

	flightMatrixXMLResponse := &wsFlightMatrixResponse.FlightMatrixXMLResponse{}
	err = xml.Unmarshal([]byte(flightMatrixResponse), flightMatrixXMLResponse)
	if err != nil {
		return nil, err
	}

	additionalFields := map[string]string{
		"FLIGHT_CLASS": request.CabinClass,
	}
	flightMatrixXMLResponse.SetAdditionalField(additionalFields)

	data, err := flightMatrixXMLResponse.ToJSON()
	if err != nil {
		return nil, err
	}

	return &response.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}, nil
}
