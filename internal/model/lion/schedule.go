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
	CabinClass               string                      `json:"cabinClass" validate:"required,oneof=PROMO ECONOMY BUSINESS"`
}

type (
	ScheduleResponse struct {
		Currency string       `json:"currency"`
		OneWay   []FlightData `json:"oneWay"`
		Return   []FlightData `json:"return"`
	}

	FlightData struct {
		CabinClass  string          `json:"cabinClass"`
		NetPrice    int             `json:"netPrice"`
		Segments    []FlightSegment `json:"segments"`
		FareDetails []FareDetail    `json:"fareDetails"`
	}

	FlightSegment struct {
		DepartureAirPortCode string           `json:"departureAirPortCode"`
		ArrivalAirPortCode   string           `json:"arrivalAirPortCode"`
		DepartureDateTime    string           `json:"departureDateTime"`
		ArrivalDateTime      string           `json:"arrivalDateTime"`
		FlightNumber         string           `json:"flightNumber"`
		OperatingAirlineCode string           `json:"operatingAirlineCode"`
		MarketingAirlineCode string           `json:"marketingAirlineCode"`
		AirEquipType         string           `json:"airEquipType"`
		FlightDuration       string           `json:"flightDuration"`
		StopsInfo            *StopInfo        `json:"stopsInfo"`
		Class                string           `json:"class"`
		AvailableClasses     []AvailableClass `json:"availableClasses"`
	}

	StopInfo struct {
		AirPortCode       string `json:"airPortCode"`
		ArrivalDateTime   string `json:"arrivalDateTime"`
		DepartureDateTime string `json:"departureDateTime"`
		StopTime          string `json:"stopTime"`
		ElapsedTime       string `json:"elapsedTime"`
		AirEquipType      string `json:"airEquipType"`
	}

	FareDetail struct {
		Code       string      `json:"code"`
		Quantity   int         `json:"quantity"`
		BaseFare   int         `json:"baseFare"`
		TotalFare  int         `json:"totalFare"`
		ClassCodes []ClassCode `json:"-"`
		Taxes      []Tax       `json:"taxes"`
	}

	ClassCode struct {
		Code string `json:"-"`
	}

	Tax struct {
		Code   string `json:"code"`
		Amount int    `json:"amount"`
	}

	AvailableClass struct {
		Code     string `json:"code"`
		Quantity int    `json:"quantity"`
	}
)
