package lion

import (
	"lion/internal/model/ws/request/flight_matrix"
)

type ScheduleRequest struct {
	ConversationID           string                      `json:"conversationID" validate:"required"`
	Direction                *flight_matrix.DirectionInd `json:"direction" validate:"required"`
	DepartureAirPort         string                      `json:"departureAirPort" validate:"required,min=3,max=3"`
	ArrivalAirPort           string                      `json:"arrivalAirPort" validate:"required,min=3,max=3"`
	DepartureDateTime        string                      `json:"departureDateTime" validate:"required,datetime=2006-01-02T15:04:05"`
	ArrivalDateTime          string                      `json:"arrivalDateTime" validate:"required,datetime=2006-01-02T15:04:05"`
	PassengerTypeQuantityADT int                         `json:"passengerTypeQuantityADT" validate:"required,numeric"`
	PassengerTypeQuantityCNN int                         `json:"passengerTypeQuantityCNN" validate:"numeric"`
	PassengerTypeQuantityINF int                         `json:"passengerTypeQuantityINF" validate:"numeric"`
}

type (
	ScheduleResponse struct {
		OneWay []FlightData `json:"oneWay"`
		Return []FlightData `json:"return"`
	}

	FlightData struct {
		Segments []FlightSegment `json:"segments"`
	}

	FlightSegment struct {
		DepartureAirPortCode string    `json:"departureAirPortCode"`
		ArrivalAirPortCode   string    `json:"arrivalAirPortCode"`
		DepartureDateTime    string    `json:"departureDateTime"`
		ArrivalDateTime      string    `json:"arrivalDateTime"`
		FlightNumber         string    `json:"flightNumber"`
		OperatingAirlineCode string    `json:"operatingAirlineCode"`
		MarketingAirlineCode string    `json:"marketingAirlineCode"`
		AirEquipType         string    `json:"airEquipType"`
		FlightDuration       string    `json:"flightDuration"`
		StopsInfo            *StopInfo `json:"stopsInfo"`
	}

	StopInfo struct {
		AirPortCode       string `json:"airPortCode"`
		ArrivalDateTime   string `json:"arrivalDateTime"`
		DepartureDateTime string `json:"departureDateTime"`
		StopTime          string `json:"stopTime"`
		ElapsedTime       string `json:"elapsedTime"`
		AirEquipType      string `json:"airEquipType"`
	}
)
