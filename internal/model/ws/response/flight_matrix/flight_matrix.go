package flight_matrix

import (
	"encoding/xml"
)

type (
	FlightMatrixXMLResponse struct {
		XMLName xml.Name `xml:"Envelope"`
		Text    string   `xml:",chardata"`
		Soap    string   `xml:"soap,attr"`
		Xsi     string   `xml:"xsi,attr"`
		Xsd     string   `xml:"xsd,attr"`
		Header  Header   `xml:"Header"`
		Body    Body     `xml:"Body"`
		Fault
	}

	Fault struct {
		FaultCode   TextElement `xml:"faultcode"`
		FaultString TextElement `xml:"faultstring"`
	}

	Header struct {
		Text          string        `xml:",chardata"`
		MessageHeader MessageHeader `xml:"MessageHeader"`
		SessionHeader SessionHeader `xml:"SessionHeader"`
	}

	MessageHeader struct {
		Text           string      `xml:",chardata"`
		Xmlns          string      `xml:"xmlns,attr"`
		CPAId          TextElement `xml:"CPAId"`
		ConversationId TextElement `xml:"ConversationId"`
		Service        TextElement `xml:"Service"`
		Action         TextElement `xml:"Action"`
		MessageData    MessageData `xml:"MessageData"`
	}

	MessageData struct {
		Text      string      `xml:",chardata"`
		Timestamp TextElement `xml:"Timestamp"`
	}

	SessionHeader struct {
		Text            string      `xml:",chardata"`
		Xmlns           string      `xml:"xmlns,attr"`
		MatrixSessionId TextElement `xml:"MatrixSessionId"`
	}

	Body struct {
		Text                         string                       `xml:",chardata"`
		FlightMatrixRequestResponse  *FlightMatrixRequestResponse `xml:"FlightMatrixRequestResponse,omitempty"`
		FlightMatrixRequest2Response *FlightMatrixRequestResponse `xml:"FlightMatrixRequest2Response,omitempty"`
	}

	FlightMatrixRequestResponse struct {
		Text           string         `xml:",chardata"`
		Xmlns          string         `xml:"xmlns,attr"`
		FlightMatrixRS FlightMatrixRS `xml:"FlightMatrixRS"`
	}

	FlightMatrixRS struct {
		Text                string              `xml:",chardata"`
		TimeStamp           string              `xml:"TimeStamp,attr"`
		Target              string              `xml:"Target,attr"`
		Version             string              `xml:"Version,attr"`
		Success             Success             `xml:"Success"`
		Errors              Errors              `xml:"Errors"`
		FreeTextType        TextElement         `xml:"FreeTextType"`
		FamilyNames         FamilyNames         `xml:"FamilyNames"`
		Currency            TextElement         `xml:"Currency"`
		FlightMatrices      FlightMatrices      `xml:"FlightMatrices"`
		FlightMatrixOptions FlightMatrixOptions `xml:"FlightMatrixOptions"`
	}

	Errors struct {
		XMLName xml.Name `xml:"Errors"`
		Error   Error    `xml:"Error"`
	}

	Error struct {
		Type      string `xml:"Type,attr"`
		ShortText string `xml:"ShortText,attr"`
	}

	Success struct {
		Text           string      `xml:",chardata"`
		SuccessMessage TextElement `xml:"SuccessMessage"`
	}

	FamilyNames struct {
		Text   string        `xml:",chardata"`
		String []TextElement `xml:"string"`
	}

	FlightMatrices struct {
		Text         string         `xml:",chardata"`
		FlightMatrix []FlightMatrix `xml:"FlightMatrix"`
	}

	FlightMatrix struct {
		Text               string           `xml:",chardata"`
		FlightSearchResult TextElement      `xml:"FlightSearchResult"`
		FlightMatrixRows   FlightMatrixRows `xml:"FlightMatrixRows"`
	}

	FlightMatrixRows struct {
		Text            string            `xml:",chardata"`
		FlightMatrixRow []FlightMatrixRow `xml:"FlightMatrixRow"`
	}

	FlightMatrixRow struct {
		Text                        string                      `xml:",chardata"`
		OriginDestinationOptionType OriginDestinationOptionType `xml:"OriginDestinationOptionType"`
		FlightMatrixCells           FlightMatrixCells           `xml:"FlightMatrixCells"`
		RPH                         TextElement                 `xml:"RPH"`
	}

	OriginDestinationOptionType struct {
		Text          string          `xml:",chardata"`
		FlightSegment []FlightSegment `xml:"FlightSegment"`
	}

	FlightSegment struct {
		Text               string             `xml:",chardata"`
		DepartureDateTime  string             `xml:"DepartureDateTime,attr"`
		ArrivalDateTime    string             `xml:"ArrivalDateTime,attr"`
		StopQuantity       string             `xml:"StopQuantity,attr"`
		RPH                string             `xml:"RPH,attr"`
		FlightNumber       string             `xml:"FlightNumber,attr"`
		DepartureAirport   DepartureAirport   `xml:"DepartureAirport"`
		ArrivalAirport     ArrivalAirport     `xml:"ArrivalAirport"`
		OperatingAirline   OperatingAirline   `xml:"OperatingAirline"`
		Equipment          Equipment          `xml:"Equipment"`
		MarketingAirline   MarketingAirline   `xml:"MarketingAirline"`
		BookingClassAvails BookingClassAvails `xml:"BookingClassAvails"`
		Duration           TextElement        `xml:"Duration"`
		StopAirports       StopAirports       `xml:"StopAirports"`
	}

	DepartureAirport struct {
		Text         string `xml:",chardata"`
		LocationCode string `xml:"LocationCode,attr"`
	}

	ArrivalAirport struct {
		Text         string `xml:",chardata"`
		LocationCode string `xml:"LocationCode,attr"`
	}

	OperatingAirline struct {
		Text             string `xml:",chardata"`
		CompanyShortName string `xml:"CompanyShortName,attr"`
		Code             string `xml:"Code,attr"`
		CodeContext      string `xml:"CodeContext,attr"`
		FlightNumber     string `xml:"FlightNumber,attr"`
	}

	Equipment struct {
		Text         string `xml:",chardata"`
		AirEquipType string `xml:"AirEquipType,attr"`
	}

	MarketingAirline struct {
		Text        string `xml:",chardata"`
		Code        string `xml:"Code,attr"`
		CodeContext string `xml:"CodeContext,attr"`
	}

	BookingClassAvails struct {
		Text              string              `xml:",chardata"`
		BookingClassAvail []BookingClassAvail `xml:"BookingClassAvail"`
	}

	BookingClassAvail struct {
		Text                 string `xml:",chardata"`
		ResBookDesigCode     string `xml:"ResBookDesigCode,attr"`
		ResBookDesigQuantity string `xml:"ResBookDesigQuantity,attr"`
	}

	StopAirports struct {
		Text            string          `xml:",chardata"`
		StopAirportType StopAirportType `xml:"StopAirportType"`
	}

	StopAirportType struct {
		Text              string `xml:",chardata"`
		PortCode          string `xml:"PortCode,attr"`
		ArrivalDateTime   string `xml:"ArrivalDateTime,attr"`
		DepartureDateTime string `xml:"DepartureDateTime,attr"`
		StopTime          string `xml:"StopTime,attr"`
		FlightTime        string `xml:"FlightTime,attr"`
		ElapsedTime       string `xml:"ElapsedTime,attr"`
		Equipment         string `xml:"Equipment,attr"`
	}

	FlightMatrixCells struct {
		Text             string             `xml:",chardata"`
		FlightMatrixCell []FlightMatrixCell `xml:"FlightMatrixCell"`
	}

	FlightMatrixCell struct {
		Text       string `xml:",chardata"`
		FamilyName string `xml:"FamilyName,attr"`
		FamilyRPH  string `xml:"FamilyRPH,attr"`
		FlightRPH  string `xml:"FlightRPH,attr"`
		Status     string `xml:"Status,attr"`
	}

	FlightMatrixOptions struct {
		Text               string               `xml:",chardata"`
		FlightMatrixOption []FlightMatrixOption `xml:"FlightMatrixOption"`
	}

	FlightMatrixOption struct {
		Text                string              `xml:",chardata"`
		PTCFareBreakdown    []PTCFareBreakdown  `xml:"PTC_FareBreakdown"`
		FlightMatrixIndices FlightMatrixIndices `xml:"FlightMatrixIndices"`
	}

	PTCFareBreakdown struct {
		Text                  string                `xml:",chardata"`
		PassengerTypeQuantity PassengerTypeQuantity `xml:"PassengerTypeQuantity"`
		FareBasisCodes        FareBasisCodes        `xml:"FareBasisCodes"`
		PassengerFare         PassengerFare         `xml:"PassengerFare"`
		TravelerRefNumber     TravelerRefNumber     `xml:"TravelerRefNumber"`
		FareInfo              []FareInfo            `xml:"FareInfo"`
	}

	PassengerTypeQuantity struct {
		Text     string `xml:",chardata"`
		Code     string `xml:"Code,attr"`
		Quantity string `xml:"Quantity,attr"`
	}

	FareBasisCodes struct {
		Text          string          `xml:",chardata"`
		FareBasisCode []FareBasisCode `xml:"FareBasisCode"`
	}

	FareBasisCode struct {
		Text string `xml:",chardata"`
	}

	PassengerFare struct {
		Text      string    `xml:",chardata"`
		BaseFare  BaseFare  `xml:"BaseFare"`
		Taxes     Taxes     `xml:"Taxes"`
		Vouchers  Vouchers  `xml:"Vouchers"`
		TotalFare TotalFare `xml:"TotalFare"`
	}

	BaseFare struct {
		Text          string `xml:",chardata"`
		Amount        string `xml:"Amount,attr"`
		CurrencyCode  string `xml:"CurrencyCode,attr"`
		DecimalPlaces string `xml:"DecimalPlaces,attr"`
	}

	Taxes struct {
		Text string `xml:",chardata"`
		Tax  []Tax  `xml:"Tax"`
	}

	Tax struct {
		Text          string `xml:",chardata"`
		TaxCode       string `xml:"TaxCode,attr"`
		Amount        string `xml:"Amount,attr"`
		CurrencyCode  string `xml:"CurrencyCode,attr"`
		DecimalPlaces string `xml:"DecimalPlaces,attr"`
	}

	Vouchers struct {
		Text string `xml:",chardata"`
	}

	TotalFare struct {
		Text          string `xml:",chardata"`
		Amount        string `xml:"Amount,attr"`
		CurrencyCode  string `xml:"CurrencyCode,attr"`
		DecimalPlaces string `xml:"DecimalPlaces,attr"`
	}

	TravelerRefNumber struct {
		Text string `xml:",chardata"`
		RPH  string `xml:"RPH,attr"`
	}

	FareInfo struct {
		Text          string          `xml:",chardata"`
		FareReference []FareReference `xml:"FareReference"`
	}

	FareReference struct {
		Text             string `xml:",chardata"`
		ResBookDesigCode string `xml:"ResBookDesigCode,attr"`
	}

	FlightMatrixIndices struct {
		Text              string              `xml:",chardata"`
		FlightMatrixIndex []FlightMatrixIndex `xml:"FlightMatrixIndex"`
	}

	FlightMatrixIndex struct {
		Text     string `xml:",chardata"`
		FlightOB struct {
			Text string `xml:",chardata"`
		} `xml:"FlightOB"`
		FamilyOB struct {
			Text string `xml:",chardata"`
		} `xml:"FamilyOB"`
		FlightIB struct {
			Text string `xml:",chardata"`
		} `xml:"FlightIB"`
		FamilyIB struct {
			Text string `xml:",chardata"`
		} `xml:"FamilyIB"`
	}

	TextElement struct {
		Text string `xml:",chardata"`
	}
)
