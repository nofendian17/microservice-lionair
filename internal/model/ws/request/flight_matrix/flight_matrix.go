package flight_matrix

import (
	"encoding/xml"
)

type DirectionInd int

const (
	OneWay DirectionInd = iota
	Return
)

func (d DirectionInd) String() string {
	switch d {
	case OneWay:
		return "OneWay"
	case Return:
		return "Return"
	default:
		return "OneWay"
	}
}

type (
	FlightMatrixXMLRequest struct {
		XMLName xml.Name `xml:"soap:Envelope"`
		SoapNS  string   `xml:"xmlns:soap,attr"`
		Header  Header   `xml:"soap:Header"`
		Body    Body     `xml:"soap:Body"`
	}
	Header struct {
		MessageHeader MessageHeader `xml:"MessageHeader"`
		Security      Security      `xml:"Security"`
	}

	MessageHeader struct {
		XMLName        xml.Name `xml:"MessageHeader"`
		Namespace      string   `xml:"xmlns,attr"`
		CPAId          string   `xml:"CPAId"`
		ConversationId string   `xml:"ConversationId"`
		Service        string   `xml:"Service"`
		Action         string   `xml:"Action"`
		MessageData    MessageData
	}

	MessageData struct {
		Timestamp string `xml:"Timestamp"`
	}

	Security struct {
		XMLName             xml.Name `xml:"Security"`
		SecurityNamespace   string   `xml:"xmlns,attr"`
		BinarySecurityToken string   `xml:"BinarySecurityToken"`
	}

	Body struct {
		FlightMatrixRequest  *FlightMatrixRequest  `xml:"FlightMatrixRequest,omitempty"`
		FlightMatrixRequest2 *FlightMatrixRequest2 `xml:"FlightMatrixRequest2,omitempty"`
	}

	FlightMatrixRequest struct {
		XMLName        xml.Name       `xml:"FlightMatrixRequest"`
		Namespace      string         `xml:"xmlns,attr"`
		FlightMatrixRQ FlightMatrixRQ `xml:"flightMatrixRQ"`
	}

	FlightMatrixRequest2 struct {
		XMLName        xml.Name       `xml:"FlightMatrixRequest2"`
		Namespace      string         `xml:"xmlns,attr"`
		FlightMatrixRQ FlightMatrixRQ `xml:"flightMatrixRQ"`
	}

	FlightMatrixRQ struct {
		XMLName             xml.Name            `xml:"flightMatrixRQ"`
		Target              string              `xml:"Target,attr"`
		Version             string              `xml:"Version,attr"`
		AirItinerary        AirItinerary        `xml:"AirItinerary"`
		TravelerInfoSummary TravelerInfoSummary `xml:"TravelerInfoSummary"`
	}

	AirItinerary struct {
		XMLName                  xml.Name                 `xml:"AirItinerary"`
		DirectionInd             string                   `xml:"DirectionInd,attr"`
		OriginDestinationOptions OriginDestinationOptions `xml:"OriginDestinationOptions"`
	}

	OriginDestinationOptions struct {
		OriginDestinationOption OriginDestinationOption `xml:"OriginDestinationOption"`
	}

	OriginDestinationOption struct {
		FlightSegment FlightSegment `xml:"FlightSegment"`
	}

	FlightSegment struct {
		XMLName           xml.Name         `xml:"FlightSegment"`
		DepartureDateTime string           `xml:"DepartureDateTime,attr"`
		ArrivalDateTime   string           `xml:"ArrivalDateTime,attr"`
		RPH               string           `xml:"RPH,attr"`
		DepartureAirport  DepartureAirport `xml:"DepartureAirport"`
		ArrivalAirport    ArrivalAirport   `xml:"ArrivalAirport"`
		MarketingAirline  MarketingAirline `xml:"MarketingAirline"`
	}

	DepartureAirport struct {
		XMLName      xml.Name `xml:"DepartureAirport"`
		LocationCode string   `xml:"LocationCode,attr"`
	}

	ArrivalAirport struct {
		XMLName      xml.Name `xml:"ArrivalAirport"`
		LocationCode string   `xml:"LocationCode,attr"`
	}

	MarketingAirline struct {
		XMLName xml.Name `xml:"MarketingAirline"`
		Code    string   `xml:"Code,attr"`
	}

	TravelerInfoSummary struct {
		XMLName          xml.Name           `xml:"TravelerInfoSummary"`
		AirTravelerAvail []AirTravelerAvail `xml:"AirTravelerAvail"`
	}

	AirTravelerAvail struct {
		XMLName     xml.Name    `xml:"AirTravelerAvail"`
		AirTraveler AirTraveler `xml:"AirTraveler"`
	}

	AirTraveler struct {
		XMLName               xml.Name                `xml:"AirTraveler"`
		PassengerTypeQuantity []PassengerTypeQuantity `xml:"PassengerTypeQuantity"`
	}

	PassengerTypeQuantity struct {
		XMLName  xml.Name `xml:"PassengerTypeQuantity"`
		Code     string   `xml:"Code,attr"`
		Quantity int      `xml:"Quantity,attr"`
	}
)
