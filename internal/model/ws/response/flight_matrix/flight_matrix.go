package flight_matrix

import (
	"encoding/xml"
	"fmt"
	"github.com/spf13/cast"
	"lion/internal/model/lion"
)

type (
	FlightMatrixXMLResponse struct {
		XMLName         xml.Name          `xml:"Envelope"`
		Text            string            `xml:",chardata"`
		Soap            string            `xml:"soap,attr"`
		Xsi             string            `xml:"xsi,attr"`
		Xsd             string            `xml:"xsd,attr"`
		Header          Header            `xml:"Header"`
		Body            Body              `xml:"Body"`
		AdditionalField map[string]string `xml:"-"`
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
		Text                         string                        `xml:",chardata"`
		FlightMatrixRequestResponse  *FlightMatrixRequestResponse  `xml:"FlightMatrixRequestResponse,omitempty"`
		FlightMatrixRequest2Response *FlightMatrixRequest2Response `xml:"FlightMatrixRequest2Response,omitempty"`
	}

	FlightMatrixRequestResponse struct {
		Text           string         `xml:",chardata"`
		Xmlns          string         `xml:"xmlns,attr"`
		FlightMatrixRS FlightMatrixRS `xml:"FlightMatrixRS"`
	}

	FlightMatrixRequest2Response struct {
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

// SetAdditionalField sets the additional fields for the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) SetAdditionalField(fields map[string]string) {
	f.AdditionalField = fields
}

// GetAdditionalField returns the additional fields of the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) GetAdditionalField() map[string]string {
	return f.AdditionalField
}

// getRequestedClass returns the requested flight class based on the additional fields.
func (f *FlightMatrixXMLResponse) getRequestedClass() string {
	additionalField := f.GetAdditionalField()
	switch additionalField["FLIGHT_CLASS"] {
	case "PROMO":
		return "0"
	case "ECONOMY":
		return "1"
	case "BUSINESS":
		return "2"
	default:
		return "-1"
	}
}

// CheckAPIError checks for API errors in the FlightMatrixXMLResponse and returns an error if present.
func (f *FlightMatrixXMLResponse) CheckAPIError() error {
	switch f.Body.FlightMatrixRequestResponse {
	case nil:
		if errorText := f.Body.FlightMatrixRequest2Response.FlightMatrixRS.Errors.Error.ShortText; errorText != "" {
			return fmt.Errorf("API response error: %s", errorText)
		}
		return nil
	default:
		if errorText := f.Body.FlightMatrixRequestResponse.FlightMatrixRS.Errors.Error.ShortText; errorText != "" {
			return fmt.Errorf("API response error: %s", errorText)
		}
		return nil
	}
}

// getCurrency returns the currency from the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) getCurrency() string {
	if f.Body.FlightMatrixRequestResponse == nil {
		return f.Body.FlightMatrixRequest2Response.FlightMatrixRS.Currency.Text
	}
	return f.Body.FlightMatrixRequestResponse.FlightMatrixRS.Currency.Text
}

// getFlightMatrices returns the flight matrices from the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) getFlightMatrices() FlightMatrices {
	if f.Body.FlightMatrixRequestResponse == nil {
		return f.Body.FlightMatrixRequest2Response.FlightMatrixRS.FlightMatrices
	}
	return f.Body.FlightMatrixRequestResponse.FlightMatrixRS.FlightMatrices
}

// getFlightMatrixOptions returns the flight matrix options from the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) getFlightMatrixOptions() FlightMatrixOptions {
	if f.Body.FlightMatrixRequestResponse == nil {
		return f.Body.FlightMatrixRequest2Response.FlightMatrixRS.FlightMatrixOptions
	}
	return f.Body.FlightMatrixRequestResponse.FlightMatrixRS.FlightMatrixOptions
}

// getAvailableClass returns the available classes for a given flight segment in the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) getAvailableClass(flightSegment FlightSegment) []lion.AvailableClass {
	classes := make([]lion.AvailableClass, len(flightSegment.BookingClassAvails.BookingClassAvail))
	for i, v := range flightSegment.BookingClassAvails.BookingClassAvail {
		class := lion.AvailableClass{
			Code:     v.ResBookDesigCode,
			Quantity: cast.ToInt(v.ResBookDesigQuantity),
		}
		classes[i] = class
	}
	return classes
}

// getFareDetails returns the fare details for a given index and RPH in the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) getFareDetails(index int, rph string) []lion.FareDetail {
	requestedClass := f.getRequestedClass()
	options := f.getFlightMatrixOptions().FlightMatrixOption

	var fareDetails []lion.FareDetail
	fareAdded := false

	for _, option := range options {
		for _, indices := range option.FlightMatrixIndices.FlightMatrixIndex {
			if (index == 0 && rph == indices.FlightOB.Text && requestedClass == indices.FamilyOB.Text && indices.FlightIB.Text == "-1" && indices.FamilyIB.Text == "-1") ||
				(index == 1 && rph == indices.FlightIB.Text && requestedClass == indices.FamilyIB.Text && indices.FlightOB.Text == "-1" && indices.FamilyOB.Text == "-1") {

				if fareAdded {
					// Fare details already added for selected flight, skip
					continue
				}

				fareDetails = append(fareDetails, f.appendFareDetails(option.PTCFareBreakdown)...)
				fareAdded = true
			}
		}
	}

	return fareDetails
}

// appendFareDetails appends fare details to an existing slice in the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) appendFareDetails(breakdowns []PTCFareBreakdown) []lion.FareDetail {
	var fareDetails []lion.FareDetail

	for _, breakdown := range breakdowns {
		taxes := make([]lion.Tax, len(breakdown.PassengerFare.Taxes.Tax))

		for i, t := range breakdown.PassengerFare.Taxes.Tax {
			tax := lion.Tax{
				Code:   t.TaxCode,
				Amount: cast.ToInt(t.Amount),
			}
			taxes[i] = tax
		}

		var classCodes []lion.ClassCode
		for _, fareInfo := range breakdown.FareInfo {
			for _, fareRef := range fareInfo.FareReference {
				classCode := lion.ClassCode{
					Code: fareRef.ResBookDesigCode,
				}
				classCodes = append(classCodes, classCode)
			}
		}

		fareDetail := lion.FareDetail{
			Code:       breakdown.PassengerTypeQuantity.Code,
			Quantity:   cast.ToInt(breakdown.PassengerTypeQuantity.Quantity),
			BaseFare:   cast.ToInt(breakdown.PassengerFare.BaseFare.Amount),
			TotalFare:  cast.ToInt(breakdown.PassengerFare.TotalFare.Amount),
			ClassCodes: classCodes,
			Taxes:      taxes,
		}
		fareDetails = append(fareDetails, fareDetail)
	}

	return fareDetails
}

// iterateSchedule iterates through the flight matrix and returns flight data in the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) iterateSchedule(index int, flightMatrix FlightMatrix) []lion.FlightData {
	flightsData := make([]lion.FlightData, 0, len(flightMatrix.FlightMatrixRows.FlightMatrixRow))

	for _, row := range flightMatrix.FlightMatrixRows.FlightMatrixRow {
		fareDetails := f.getFareDetails(index, row.RPH.Text)
		if len(fareDetails) == 0 {
			continue
		}

		segments := make([]lion.FlightSegment, len(row.OriginDestinationOptionType.FlightSegment))
		for i, v := range row.OriginDestinationOptionType.FlightSegment {
			var stopInfo *lion.StopInfo
			if v.StopAirports.StopAirportType.ArrivalDateTime != "" {
				stopInfo = &lion.StopInfo{
					AirPortCode:       v.StopAirports.StopAirportType.PortCode,
					ArrivalDateTime:   v.StopAirports.StopAirportType.ArrivalDateTime,
					DepartureDateTime: v.StopAirports.StopAirportType.DepartureDateTime,
					StopTime:          v.StopAirports.StopAirportType.StopTime,
					ElapsedTime:       v.StopAirports.StopAirportType.ElapsedTime,
					AirEquipType:      v.StopAirports.StopAirportType.Equipment,
				}
			}

			classCode := ""
			if i < len(fareDetails[0].ClassCodes) {
				classCode = fareDetails[0].ClassCodes[i].Code
			}

			segments[i] = lion.FlightSegment{
				DepartureAirPortCode: v.DepartureAirport.LocationCode,
				ArrivalAirPortCode:   v.ArrivalAirport.LocationCode,
				DepartureDateTime:    v.DepartureDateTime,
				ArrivalDateTime:      v.ArrivalDateTime,
				FlightNumber:         v.FlightNumber,
				OperatingAirlineCode: v.OperatingAirline.Code,
				MarketingAirlineCode: v.MarketingAirline.Code,
				AirEquipType:         v.Equipment.AirEquipType,
				FlightDuration:       v.Duration.Text,
				StopsInfo:            stopInfo,
				Class:                classCode,
				AvailableClasses:     f.getAvailableClass(v),
			}
		}

		additionalField := f.GetAdditionalField()

		// Calculate NetPrice from TotalFare
		var netPrice int
		for _, fareDetail := range fareDetails {
			netPrice += fareDetail.TotalFare
		}

		flightData := lion.FlightData{
			Segments:    segments,
			CabinClass:  additionalField["FLIGHT_CLASS"],
			NetPrice:    netPrice,
			FareDetails: fareDetails,
		}

		flightsData = append(flightsData, flightData)
	}

	return flightsData
}

// getFlightSchedule returns the flight schedules from the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) getFlightSchedule() map[string][]lion.FlightData {
	schedules := make(map[string][]lion.FlightData)
	flightMatrices := f.getFlightMatrices().FlightMatrix

	// Initialize the schedules with empty slices
	schedules["ONEWAY"] = []lion.FlightData{}
	schedules["RETURN"] = []lion.FlightData{}

	for index, flightMatrix := range flightMatrices {
		switch index {
		case 0:
			schedules["ONEWAY"] = f.iterateSchedule(index, flightMatrix)
		case 1:
			schedules["RETURN"] = f.iterateSchedule(index, flightMatrix)
		}
	}
	return schedules
}

// getFlightData returns the flight data from the FlightMatrixXMLResponse.
func (f *FlightMatrixXMLResponse) getFlightData() *lion.ScheduleResponse {
	schedules := f.getFlightSchedule()
	return &lion.ScheduleResponse{
		Currency: f.getCurrency(),
		OneWay:   schedules["ONEWAY"],
		Return:   schedules["RETURN"],
	}
}

// ToJSON converts the FlightMatrixXMLResponse to a JSON format and returns a ScheduleResponse.
// It checks for API errors and returns an error if present.
func (f *FlightMatrixXMLResponse) ToJSON() (*lion.ScheduleResponse, error) {
	if err := f.CheckAPIError(); err != nil {
		return nil, err
	}
	return f.getFlightData(), nil
}
