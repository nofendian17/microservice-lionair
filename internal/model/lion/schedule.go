package lion

import (
	"lion/internal/model/ws/request/flight_matrix"
)

type ScheduleRequest struct {
	ConversationID    string                      `json:"conversationID" validate:"required"`
	SessionID         string                      `json:"sessionID" validate:"required,uuid"`
	Direction         *flight_matrix.DirectionInd `json:"direction" validate:"required"`
	DepartureAirPort  string                      `json:"departureAirPort" validate:"required,min=3,max=3"`
	ArrivalAirPort    string                      `json:"arrivalAirPort" validate:"required,min=3,max=3"`
	DepartureDateTime string                      `json:"departureDateTime" validate:"required,datetime=2006-01-02T15:04:05"`
	ArrivalDateTime   string                      `json:"arrivalDateTime" validate:"required,datetime=2006-01-02T15:04:05"`
	QuantityADT       int                         `json:"adt" validate:"required,numeric"`
	QuantityCNN       int                         `json:"cnn" validate:"numeric"`
	QuantityINF       int                         `json:"inf" validate:"numeric"`
}

type ScheduleResponse struct {
	OW []OriginDestinationOption `json:"ow"`
	RT []OriginDestinationOption `json:"rt"`
}

type OriginDestinationOption struct {
	RPH           uint            `json:"rph"`
	FlightSegment []FlightSegment `json:"flightSegment"`
}

type FlightSegment struct {
	ArrivalDateTime     string              `json:"arrivalDateTime"`
	DepartureDateTime   string              `json:"departureDateTime"`
	FlightNumber        string              `json:"flightNumber"`
	StopQuantity        uint                `json:"stopQuantity"`
	OriginLocation      string              `json:"originLocation"`
	DestinationLocation string              `json:"destinationLocation"`
	Equipment           string              `json:"equipment"`
	MarketingAirline    string              `json:"marketingAirline"`
	OperatingAirline    string              `json:"operatingAirline"`
	Duration            string              `json:"duration"`
	Meal                string              `json:"meal"`
	BookingClassAvails  []BookingClassAvail `json:"bookingClassAvails"`
}

type BookingClassAvail struct {
	Class    string `json:"class"`
	Quantity uint   `json:"quantity"`
}
