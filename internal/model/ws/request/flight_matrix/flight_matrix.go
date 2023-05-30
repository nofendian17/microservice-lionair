package flight_matrix

import (
	"encoding/xml"
	"errors"
	"time"
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

func (r *FlightMatrixXMLRequest) NewFlightMatrix(conversationID, binarySecurityToken, target, version, direction, departureDate, arrivalDate, departureAirportCode, arrivalAirportCode string, adt, cnn, inf int) (string, error) {
	if binarySecurityToken == "" {
		return "", errors.New("binary security token is empty")
	}

	timestamp := time.Now().Format("2006-01-02T15:04:05Z")

	r.SoapNS = "http://schemas.xmlsoap.org/soap/envelope/"
	r.Header.MessageHeader.Namespace = "http://www.ebxml.org/namespaces/messageHeader"
	r.Header.MessageHeader.CPAId = "DEFAULT"
	r.Header.MessageHeader.ConversationId = conversationID
	r.Header.MessageHeader.Service = "GetFlightMatrix"
	r.Header.MessageHeader.Action = "FlightMatrixRQ"
	r.Header.MessageHeader.MessageData.Timestamp = timestamp
	r.Header.Security.SecurityNamespace = "http://schemas.xmlsoap.org/ws/2002/12/secext"
	r.Header.Security.BinarySecurityToken = binarySecurityToken
	r.Body.FlightMatrixRequest = &FlightMatrixRequest{
		Namespace: "http://www.vedaleon.com/webservices",
		FlightMatrixRQ: FlightMatrixRQ{
			Target:  target,
			Version: version,
			AirItinerary: AirItinerary{
				DirectionInd: direction,
				OriginDestinationOptions: OriginDestinationOptions{
					OriginDestinationOption: OriginDestinationOption{
						FlightSegment: FlightSegment{
							DepartureDateTime: departureDate,
							ArrivalDateTime:   arrivalDate,
							RPH:               "1",
							DepartureAirport: DepartureAirport{
								LocationCode: departureAirportCode,
							},
							ArrivalAirport: ArrivalAirport{
								LocationCode: arrivalAirportCode,
							},
							MarketingAirline: MarketingAirline{
								Code: "JT",
							},
						},
					},
				},
			},
			TravelerInfoSummary: TravelerInfoSummary{},
		},
	}

	appendAirTravelerAvail := func(code string, quantity int) {
		passenger := PassengerTypeQuantity{
			Code:     code,
			Quantity: quantity,
		}
		airTraveler := AirTraveler{
			PassengerTypeQuantity: []PassengerTypeQuantity{passenger},
		}
		airTravelerAvail := AirTravelerAvail{
			AirTraveler: airTraveler,
		}
		r.Body.FlightMatrixRequest.FlightMatrixRQ.TravelerInfoSummary.AirTravelerAvail = append(r.Body.FlightMatrixRequest.FlightMatrixRQ.TravelerInfoSummary.AirTravelerAvail, airTravelerAvail)
	}

	if adt > 0 {
		appendAirTravelerAvail("ADT", adt)
	}

	if cnn > 0 {
		appendAirTravelerAvail("CNN", cnn)
	}

	if inf > 0 {
		appendAirTravelerAvail("INF", inf)
	}

	xmlData, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}

	return string(xmlData), nil
}

func (r *FlightMatrixXMLRequest) NewFlightMatrix2(conversationID, target, version, direction, departureDate, arrivalDate, departureAirportCode, arrivalAirportCode string, adt, cnn, inf int) (string, error) {
	timestamp := time.Now().Format("2006-01-02T15:04:05Z")

	r.SoapNS = "http://schemas.xmlsoap.org/soap/envelope/"
	r.Header.MessageHeader.Namespace = "http://www.ebxml.org/namespaces/messageHeader"
	r.Header.MessageHeader.CPAId = "DEFAULT"
	r.Header.MessageHeader.ConversationId = conversationID
	r.Header.MessageHeader.Service = "GetFlightMatrix"
	r.Header.MessageHeader.Action = "FlightMatrixRQ"
	r.Header.MessageHeader.MessageData.Timestamp = timestamp
	r.Header.Security.SecurityNamespace = "http://schemas.xmlsoap.org/ws/2002/12/secext"
	r.Header.Security.BinarySecurityToken = ""
	r.Body.FlightMatrixRequest2 = &FlightMatrixRequest2{
		Namespace: "http://www.vedaleon.com/webservices",
		FlightMatrixRQ: FlightMatrixRQ{
			Target:  target,
			Version: version,
			AirItinerary: AirItinerary{
				DirectionInd: direction,
				OriginDestinationOptions: OriginDestinationOptions{
					OriginDestinationOption: OriginDestinationOption{
						FlightSegment: FlightSegment{
							DepartureDateTime: departureDate,
							ArrivalDateTime:   arrivalDate,
							RPH:               "1",
							DepartureAirport: DepartureAirport{
								LocationCode: departureAirportCode,
							},
							ArrivalAirport: ArrivalAirport{
								LocationCode: arrivalAirportCode,
							},
							MarketingAirline: MarketingAirline{
								Code: "JT",
							},
						},
					},
				},
			},
			TravelerInfoSummary: TravelerInfoSummary{},
		},
	}

	appendAirTravelerAvail := func(code string, quantity int) {
		passenger := PassengerTypeQuantity{
			Code:     code,
			Quantity: quantity,
		}
		airTraveler := AirTraveler{
			PassengerTypeQuantity: []PassengerTypeQuantity{passenger},
		}
		airTravelerAvail := AirTravelerAvail{
			AirTraveler: airTraveler,
		}
		r.Body.FlightMatrixRequest2.FlightMatrixRQ.TravelerInfoSummary.AirTravelerAvail = append(r.Body.FlightMatrixRequest2.FlightMatrixRQ.TravelerInfoSummary.AirTravelerAvail, airTravelerAvail)
	}

	if adt > 0 {
		appendAirTravelerAvail("ADT", adt)
	}

	if cnn > 0 {
		appendAirTravelerAvail("CNN", cnn)
	}

	if inf > 0 {
		appendAirTravelerAvail("INF", inf)
	}

	xmlData, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}

	return string(xmlData), nil
}
