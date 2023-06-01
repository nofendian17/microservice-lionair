package ota_air_sell

import (
	"encoding/xml"
)

type OTAAirSellXMLRequest struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Header  Header   `xml:"soap:Header"`
	Body    Body     `xml:"soap:Body"`
}

type Header struct {
	Text          string        `xml:",chardata"`
	MessageHeader MessageHeader `xml:"MessageHeader"`
	Security      Security      `xml:"Security"`
}

type MessageHeader struct {
	Text           string         `xml:",chardata"`
	Xmlns          string         `xml:"xmlns,attr"`
	CPAId          CPAId          `xml:"CPAId"`
	ConversationId ConversationId `xml:"ConversationId"`
	Service        Service        `xml:"Service"`
	Action         Action         `xml:"Action"`
	MessageData    MessageData    `xml:"MessageData"`
}

type CPAId struct {
	Text string `xml:",chardata"`
}

type ConversationId struct {
	Text string `xml:",chardata"`
}

type Service struct {
	Text string `xml:",chardata"`
}

type Action struct {
	Text string `xml:",chardata"`
}

type MessageData struct {
	Text      string    `xml:",chardata"`
	MessageId MessageId `xml:"MessageId"`
	Timestamp Timestamp `xml:"Timestamp"`
}

type MessageId struct {
	Text string `xml:",chardata"`
}

type Timestamp struct {
	Text string `xml:",chardata"`
}

type Security struct {
	Text                string              `xml:",chardata"`
	Xmlns               string              `xml:"xmlns,attr"`
	BinarySecurityToken BinarySecurityToken `xml:"BinarySecurityToken"`
}

type BinarySecurityToken struct {
	Text string `xml:",chardata"`
}

type Body struct {
	Text      string    `xml:",chardata"`
	AirSellRQ AirSellRQ `xml:"AirSellRQ"`
}

type AirSellRQ struct {
	Text          string        `xml:",chardata"`
	Xmlns         string        `xml:"xmlns,attr"`
	OTAAirPriceRQ OTAAirPriceRQ `xml:"OTA_AirPriceRQ"`
}

type OTAAirPriceRQ struct {
	Text                         string              `xml:",chardata"`
	Target                       string              `xml:"Target,attr"`
	Version                      string              `xml:"Version,attr"`
	ReturnTerminalInfoInResponse string              `xml:"ReturnTerminalInfoInResponse,attr"`
	POS                          POS                 `xml:"POS"`
	AirItinerary                 AirItinerary        `xml:"AirItinerary"`
	TravelerInfoSummary          TravelerInfoSummary `xml:"TravelerInfoSummary"`
}

type POS struct {
	Text   string `xml:",chardata"`
	Source Source `xml:"Source"`
}

type Source struct {
	Text       string `xml:",chardata"`
	ISOCountry string `xml:"ISOCountry,attr"`
}

type AirItinerary struct {
	Text                     string                   `xml:",chardata"`
	DirectionInd             string                   `xml:"DirectionInd,attr"`
	OriginDestinationOptions OriginDestinationOptions `xml:"OriginDestinationOptions"`
}

type OriginDestinationOptions struct {
	Text                    string                    `xml:",chardata"`
	OriginDestinationOption []OriginDestinationOption `xml:"OriginDestinationOption"`
}

type OriginDestinationOption struct {
	Text          string        `xml:",chardata"`
	FlightSegment FlightSegment `xml:"FlightSegment"`
}

type FlightSegment struct {
	Text              string           `xml:",chardata"`
	DepartureDateTime string           `xml:"DepartureDateTime,attr"`
	ArrivalDateTime   string           `xml:"ArrivalDateTime,attr"`
	StopQuantity      string           `xml:"StopQuantity,attr"`
	RPH               string           `xml:"RPH,attr"`
	FlightNumber      string           `xml:"FlightNumber,attr"`
	ResBookDesigCode  string           `xml:"ResBookDesigCode,attr"`
	NumberInParty     string           `xml:"NumberInParty,attr"`
	Status            string           `xml:"Status,attr"`
	DepartureAirport  DepartureAirport `xml:"DepartureAirport"`
	ArrivalAirport    ArrivalAirport   `xml:"ArrivalAirport"`
	OperatingAirline  OperatingAirline `xml:"OperatingAirline"`
	MarketingAirline  MarketingAirline `xml:"MarketingAirline"`
}

type DepartureAirport struct {
	Text         string `xml:",chardata"`
	LocationCode string `xml:"LocationCode,attr"`
}

type ArrivalAirport struct {
	Text         string `xml:",chardata"`
	LocationCode string `xml:"LocationCode,attr"`
}

type OperatingAirline struct {
	Text        string `xml:",chardata"`
	Code        string `xml:"Code,attr"`
	CodeContext string `xml:"CodeContext,attr"`
}

type MarketingAirline struct {
	Text        string `xml:",chardata"`
	Code        string `xml:"Code,attr"`
	CodeContext string `xml:"CodeContext,attr"`
}

type TravelerInfoSummary struct {
	Text             string             `xml:",chardata"`
	AirTravelerAvail []AirTravelerAvail `xml:"AirTravelerAvail"`
}

type AirTravelerAvail struct {
	Text        string      `xml:",chardata"`
	AirTraveler AirTraveler `xml:"AirTraveler"`
}

type AirTraveler struct {
	Text                  string                `xml:",chardata"`
	PassengerTypeQuantity PassengerTypeQuantity `xml:"PassengerTypeQuantity"`
}

type PassengerTypeQuantity struct {
	Text     string `xml:",chardata"`
	Code     string `xml:"Code,attr"`
	Quantity string `xml:"Quantity,attr"`
}
