package lion

import (
	"context"
	"encoding/xml"
	"errors"
	"lion/internal/model/lion"
	wsFlightMatrixRequest "lion/internal/model/ws/request/flight_matrix"
	wsFlightMatrixResponse "lion/internal/model/ws/response/flight_matrix"
	"time"
)

type FlightMatrix interface {
	GetFlightScheduleWithSession(
		ctx context.Context,
		conversationID string,
		binarySecurityToken string,
		apiEnv string,
		apiVersion string,
		direction string,
		departureDateTime string,
		arrivalDateTime string,
		departureAirport string,
		arrivalAirport string,
		qtyAdt int,
		qtyCnn int,
		qtyInf int,
	) (*lion.ScheduleResponse, error)
	GetFlightScheduleWithOutSession(
		ctx context.Context,
		conversationID string,
		apiEnv string,
		apiVersion string,
		direction string,
		departureDateTime string,
		arrivalDateTime string,
		departureAirport string,
		arrivalAirport string,
		qtyAdt int,
		qtyCnn int,
		qtyInf int,
	) (*lion.ScheduleResponse, error)
}

type flightMatrix struct {
	lionRequest Request
}

func (f *flightMatrix) GetFlightScheduleWithSession(
	ctx context.Context,
	conversationID string,
	binarySecurityToken string,
	apiEnv string,
	apiVersion string,
	direction string,
	departureDateTime string,
	arrivalDateTime string,
	departureAirport string,
	arrivalAirport string,
	qtyAdt int,
	qtyCnn int,
	qtyInf int,
) (*lion.ScheduleResponse, error) {
	if binarySecurityToken == "" {
		return nil, errors.New("binary security token is empty")
	}

	timestamp := time.Now().Format("2006-01-02T15:04:05Z")
	flightMatrixRequest := &wsFlightMatrixRequest.FlightMatrixXMLRequest{}
	flightMatrixRequest.SoapNS = "http://schemas.xmlsoap.org/soap/envelope/"
	flightMatrixRequest.Header.MessageHeader.Namespace = "http://www.ebxml.org/namespaces/messageHeader"
	flightMatrixRequest.Header.MessageHeader.CPAId = "DEFAULT"
	flightMatrixRequest.Header.MessageHeader.ConversationId = conversationID
	flightMatrixRequest.Header.MessageHeader.Service = "GetFlightMatrix"
	flightMatrixRequest.Header.MessageHeader.Action = "FlightMatrixRQ"
	flightMatrixRequest.Header.MessageHeader.MessageData.Timestamp = timestamp
	flightMatrixRequest.Header.Security.SecurityNamespace = "http://schemas.xmlsoap.org/ws/2002/12/secext"
	flightMatrixRequest.Header.Security.BinarySecurityToken = binarySecurityToken
	flightMatrixRequest.Body.FlightMatrixRequest = &wsFlightMatrixRequest.FlightMatrixRequest{
		Namespace: "http://www.vedaleon.com/webservices",
		FlightMatrixRQ: wsFlightMatrixRequest.FlightMatrixRQ{
			Target:  apiEnv,
			Version: apiVersion,
			AirItinerary: wsFlightMatrixRequest.AirItinerary{
				DirectionInd: direction,
				OriginDestinationOptions: wsFlightMatrixRequest.OriginDestinationOptions{
					OriginDestinationOption: wsFlightMatrixRequest.OriginDestinationOption{
						FlightSegment: wsFlightMatrixRequest.FlightSegment{
							DepartureDateTime: departureDateTime,
							ArrivalDateTime:   arrivalDateTime,
							RPH:               "1",
							DepartureAirport: wsFlightMatrixRequest.DepartureAirport{
								LocationCode: departureAirport,
							},
							ArrivalAirport: wsFlightMatrixRequest.ArrivalAirport{
								LocationCode: arrivalAirport,
							},
							MarketingAirline: wsFlightMatrixRequest.MarketingAirline{
								Code: "JT",
							},
						},
					},
				},
			},
			TravelerInfoSummary: wsFlightMatrixRequest.TravelerInfoSummary{},
		},
	}

	appendAirTravelerAvail := func(code string, quantity int) {
		passenger := wsFlightMatrixRequest.PassengerTypeQuantity{
			Code:     code,
			Quantity: quantity,
		}
		airTraveler := wsFlightMatrixRequest.AirTraveler{
			PassengerTypeQuantity: []wsFlightMatrixRequest.PassengerTypeQuantity{passenger},
		}
		airTravelerAvail := wsFlightMatrixRequest.AirTravelerAvail{
			AirTraveler: airTraveler,
		}
		flightMatrixRequest.Body.FlightMatrixRequest.FlightMatrixRQ.TravelerInfoSummary.AirTravelerAvail = append(flightMatrixRequest.Body.FlightMatrixRequest.FlightMatrixRQ.TravelerInfoSummary.AirTravelerAvail, airTravelerAvail)
	}

	if qtyAdt > 0 {
		appendAirTravelerAvail("ADT", qtyAdt)
	}

	if qtyCnn > 0 {
		appendAirTravelerAvail("CNN", qtyCnn)
	}

	if qtyInf > 0 {
		appendAirTravelerAvail("INF", qtyInf)
	}

	flightMatrixXMLRequest, err := xml.MarshalIndent(flightMatrixRequest, "", "    ")
	if err != nil {
		return nil, err
	}

	flightMatrixResponse, err := f.lionRequest.FlightMatrix(ctx, string(flightMatrixXMLRequest))
	if err != nil {
		return nil, err
	}
	flightMatrixXMLResponse := &wsFlightMatrixResponse.FlightMatrixXMLResponse{}
	err = xml.Unmarshal([]byte(flightMatrixResponse), flightMatrixXMLResponse)
	if err != nil {
		return nil, err
	}

	// todo : need parsing schedule

	return &lion.ScheduleResponse{
		OneWay: nil,
		Return: nil,
	}, nil
}

func (f *flightMatrix) GetFlightScheduleWithOutSession(
	ctx context.Context,
	conversationID string,
	apiEnv string,
	apiVersion string,
	direction string,
	departureDateTime string,
	arrivalDateTime string,
	departureAirport string,
	arrivalAirport string,
	qtyAdt int,
	qtyCnn int,
	qtyInf int,
) (*lion.ScheduleResponse, error) {
	timestamp := time.Now().Format("2006-01-02T15:04:05Z")
	flightMatrixRequest := &wsFlightMatrixRequest.FlightMatrixXMLRequest{}
	flightMatrixRequest.SoapNS = "http://schemas.xmlsoap.org/soap/envelope/"
	flightMatrixRequest.Header.MessageHeader.Namespace = "http://www.ebxml.org/namespaces/messageHeader"
	flightMatrixRequest.Header.MessageHeader.CPAId = "DEFAULT"
	flightMatrixRequest.Header.MessageHeader.ConversationId = conversationID
	flightMatrixRequest.Header.MessageHeader.Service = "GetFlightMatrix"
	flightMatrixRequest.Header.MessageHeader.Action = "FlightMatrixRQ"
	flightMatrixRequest.Header.MessageHeader.MessageData.Timestamp = timestamp
	flightMatrixRequest.Header.Security.SecurityNamespace = "http://schemas.xmlsoap.org/ws/2002/12/secext"
	flightMatrixRequest.Header.Security.BinarySecurityToken = ""
	flightMatrixRequest.Body.FlightMatrixRequest2 = &wsFlightMatrixRequest.FlightMatrixRequest2{
		Namespace: "http://www.vedaleon.com/webservices",
		FlightMatrixRQ: wsFlightMatrixRequest.FlightMatrixRQ{
			Target:  apiEnv,
			Version: apiVersion,
			AirItinerary: wsFlightMatrixRequest.AirItinerary{
				DirectionInd: direction,
				OriginDestinationOptions: wsFlightMatrixRequest.OriginDestinationOptions{
					OriginDestinationOption: wsFlightMatrixRequest.OriginDestinationOption{
						FlightSegment: wsFlightMatrixRequest.FlightSegment{
							DepartureDateTime: departureDateTime,
							ArrivalDateTime:   arrivalDateTime,
							RPH:               "1",
							DepartureAirport: wsFlightMatrixRequest.DepartureAirport{
								LocationCode: departureAirport,
							},
							ArrivalAirport: wsFlightMatrixRequest.ArrivalAirport{
								LocationCode: arrivalAirport,
							},
							MarketingAirline: wsFlightMatrixRequest.MarketingAirline{
								Code: "JT",
							},
						},
					},
				},
			},
			TravelerInfoSummary: wsFlightMatrixRequest.TravelerInfoSummary{},
		},
	}

	appendAirTravelerAvail := func(code string, quantity int) {
		passenger := wsFlightMatrixRequest.PassengerTypeQuantity{
			Code:     code,
			Quantity: quantity,
		}
		airTraveler := wsFlightMatrixRequest.AirTraveler{
			PassengerTypeQuantity: []wsFlightMatrixRequest.PassengerTypeQuantity{passenger},
		}
		airTravelerAvail := wsFlightMatrixRequest.AirTravelerAvail{
			AirTraveler: airTraveler,
		}
		flightMatrixRequest.Body.FlightMatrixRequest2.FlightMatrixRQ.TravelerInfoSummary.AirTravelerAvail = append(flightMatrixRequest.Body.FlightMatrixRequest2.FlightMatrixRQ.TravelerInfoSummary.AirTravelerAvail, airTravelerAvail)
	}

	if qtyAdt > 0 {
		appendAirTravelerAvail("ADT", qtyAdt)
	}

	if qtyCnn > 0 {
		appendAirTravelerAvail("CNN", qtyCnn)
	}

	if qtyInf > 0 {
		appendAirTravelerAvail("INF", qtyInf)
	}

	flightMatrixXMLRequest, err := xml.MarshalIndent(flightMatrixRequest, "", "    ")
	if err != nil {
		return nil, err
	}

	flightMatrixResponse, err := f.lionRequest.FlightMatrix(ctx, string(flightMatrixXMLRequest))
	if err != nil {
		return nil, err
	}
	flightMatrixXMLResponse := &wsFlightMatrixResponse.FlightMatrixXMLResponse{}
	err = xml.Unmarshal([]byte(flightMatrixResponse), flightMatrixXMLResponse)
	if err != nil {
		return nil, err
	}

	// todo : need parsing schedule
	return &lion.ScheduleResponse{
		OneWay: nil,
		Return: nil,
	}, nil
}

func NewFlightMatrix(lionRequest Request) FlightMatrix {
	return &flightMatrix{
		lionRequest: lionRequest,
	}
}
