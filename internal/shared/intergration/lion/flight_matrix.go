package lion

import (
	"context"
	"encoding/xml"
	"errors"
	"github.com/spf13/cast"
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

	return &lion.ScheduleResponse{}, nil
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

	originDestinationOptions, err := f.originDestinationOptions(flightMatrixResponse)
	if err != nil {
		return nil, err
	}
	return originDestinationOptions, nil
}

func (f *flightMatrix) originDestinationOptions(response string) (*lion.ScheduleResponse, error) {
	flightMatrixXMLResponse := &wsFlightMatrixResponse.FlightMatrixXMLResponse{}
	err := xml.Unmarshal([]byte(response), flightMatrixXMLResponse)
	if err != nil {
		return nil, err
	}

	flightMatrixResponse := flightMatrixXMLResponse.Body.FlightMatrixRequest2Response
	if flightMatrixResponse == nil {
		flightMatrixResponse = flightMatrixXMLResponse.Body.FlightMatrixRequestResponse
	}

	flightMatrices := flightMatrixResponse.FlightMatrixRS.FlightMatrices.FlightMatrix
	schedule := make(map[int][]lion.OriginDestinationOption, len(flightMatrices))

	for i, fm := range flightMatrices {
		originDestOpts := make([]lion.OriginDestinationOption, len(fm.FlightMatrixRows.FlightMatrixRow))

		for j, row := range fm.FlightMatrixRows.FlightMatrixRow {
			originDestOpts[j] = lion.OriginDestinationOption{
				RPH:           cast.ToUint(row.RPH.Text),
				FlightSegment: f.extractFlightSegment(row),
			}
		}

		schedule[i] = originDestOpts
	}

	scheduleResponse := &lion.ScheduleResponse{
		OW: schedule[0],
		RT: schedule[1],
	}

	if scheduleResponse.OW == nil {
		scheduleResponse.OW = make([]lion.OriginDestinationOption, 0)
	}

	if scheduleResponse.RT == nil {
		scheduleResponse.RT = make([]lion.OriginDestinationOption, 0)
	}

	return scheduleResponse, nil
}

func (f *flightMatrix) extractFlightSegment(rows wsFlightMatrixResponse.FlightMatrixRow) []lion.FlightSegment {
	var flightSegments []lion.FlightSegment
	for _, segment := range rows.OriginDestinationOptionType.FlightSegment {
		flightSegment := lion.FlightSegment{
			ArrivalDateTime:     segment.ArrivalDateTime,
			DepartureDateTime:   segment.DepartureDateTime,
			FlightNumber:        segment.FlightNumber,
			StopQuantity:        cast.ToUint(segment.StopQuantity),
			OriginLocation:      segment.DepartureAirport.LocationCode,
			DestinationLocation: segment.ArrivalAirport.LocationCode,
			Equipment:           segment.Equipment.AirEquipType,
			MarketingAirline:    segment.MarketingAirline.Code,
			OperatingAirline:    segment.OperatingAirline.Code,
			Duration:            segment.Duration.Text,
			Meal:                "",
			BookingClassAvails:  f.extractBookingClassAvails(segment),
		}
		flightSegments = append(flightSegments, flightSegment)
	}
	return flightSegments
}

func (f *flightMatrix) extractBookingClassAvails(segment wsFlightMatrixResponse.FlightSegment) []lion.BookingClassAvail {
	var classes []lion.BookingClassAvail
	for _, v := range segment.BookingClassAvails.BookingClassAvail {
		class := lion.BookingClassAvail{
			Class:    v.ResBookDesigCode,
			Quantity: cast.ToUint(v.ResBookDesigQuantity),
		}
		classes = append(classes, class)
	}
	return classes
}

func NewFlightMatrix(lionRequest Request) FlightMatrix {
	return &flightMatrix{
		lionRequest: lionRequest,
	}
}
