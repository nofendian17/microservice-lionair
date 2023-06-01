package ota_air_sell

import (
	"encoding/xml"
)

type OTAAirSellXMLResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Header  Header   `xml:"Header"`
	Body    Body     `xml:"Body"`
}

type Header struct {
	Text          string        `xml:",chardata"`
	MessageHeader MessageHeader `xml:"MessageHeader"`
	Security      Security      `xml:"Security"`
}

type MessageHeader struct {
	Text  string `xml:",chardata"`
	Xmlns string `xml:"xmlns,attr"`
	CPAId struct {
		Text string `xml:",chardata"`
	} `xml:"CPAId"`
	ConversationId struct {
		Text string `xml:",chardata"`
	} `xml:"ConversationId"`
	Service struct {
		Text string `xml:",chardata"`
	} `xml:"Service"`
	Action struct {
		Text string `xml:",chardata"`
	} `xml:"Action"`
	MessageData struct {
		Text      string `xml:",chardata"`
		MessageId struct {
			Text string `xml:",chardata"`
		} `xml:"MessageId"`
		Timestamp struct {
			Text string `xml:",chardata"`
		} `xml:"Timestamp"`
	} `xml:"MessageData"`
}

type Security struct {
	Text                string `xml:",chardata"`
	Xmlns               string `xml:"xmlns,attr"`
	BinarySecurityToken struct {
		Text string `xml:",chardata"`
	} `xml:"BinarySecurityToken"`
}

type Body struct {
	Text              string `xml:",chardata"`
	AirSellRQResponse struct {
		Text          string `xml:",chardata"`
		Xmlns         string `xml:"xmlns,attr"`
		OTAAirPriceRS struct {
			Text      string `xml:",chardata"`
			TimeStamp string `xml:"TimeStamp,attr"`
			Target    string `xml:"Target,attr"`
			Version   string `xml:"Version,attr"`
			Success   struct {
				Text           string `xml:",chardata"`
				SuccessMessage struct {
					Text string `xml:",chardata"`
				} `xml:"SuccessMessage"`
			} `xml:"Success"`
			Warnings struct {
				Text string `xml:",chardata"`
			} `xml:"Warnings"`
			PricedItineraries struct {
				Text            string `xml:",chardata"`
				PricedItinerary struct {
					Text           string `xml:",chardata"`
					SequenceNumber string `xml:"SequenceNumber,attr"`
					AirItinerary   struct {
						Text                     string `xml:",chardata"`
						OriginDestinationOptions struct {
							Text                    string `xml:",chardata"`
							OriginDestinationOption []struct {
								Text          string `xml:",chardata"`
								FlightSegment struct {
									Text              string `xml:",chardata"`
									DepartureDateTime string `xml:"DepartureDateTime,attr"`
									ArrivalDateTime   string `xml:"ArrivalDateTime,attr"`
									FlightNumber      string `xml:"FlightNumber,attr"`
									ResBookDesigCode  string `xml:"ResBookDesigCode,attr"`
									NumberInParty     string `xml:"NumberInParty,attr"`
									Status            string `xml:"Status,attr"`
									DepartureAirport  struct {
										Text         string `xml:",chardata"`
										LocationCode string `xml:"LocationCode,attr"`
										Terminal     string `xml:"Terminal,attr"`
										Gate         string `xml:"Gate,attr"`
									} `xml:"DepartureAirport"`
									ArrivalAirport struct {
										Text         string `xml:",chardata"`
										LocationCode string `xml:"LocationCode,attr"`
										Terminal     string `xml:"Terminal,attr"`
									} `xml:"ArrivalAirport"`
									OperatingAirline struct {
										Text         string `xml:",chardata"`
										Code         string `xml:"Code,attr"`
										FlightNumber string `xml:"FlightNumber,attr"`
									} `xml:"OperatingAirline"`
									MarketingAirline struct {
										Text string `xml:",chardata"`
										Code string `xml:"Code,attr"`
									} `xml:"MarketingAirline"`
								} `xml:"FlightSegment"`
							} `xml:"OriginDestinationOption"`
						} `xml:"OriginDestinationOptions"`
					} `xml:"AirItinerary"`
					AirItineraryPricingInfo struct {
						Text          string `xml:",chardata"`
						ItinTotalFare struct {
							Text      string `xml:",chardata"`
							TotalFare struct {
								Text         string `xml:",chardata"`
								Amount       string `xml:"Amount,attr"`
								CurrencyCode string `xml:"CurrencyCode,attr"`
							} `xml:"TotalFare"`
						} `xml:"ItinTotalFare"`
						Discounts struct {
							Text string `xml:",chardata"`
						} `xml:"Discounts"`
						PTCFareBreakdowns struct {
							Text             string `xml:",chardata"`
							PTCFareBreakdown []struct {
								Text                  string `xml:",chardata"`
								PassengerTypeQuantity struct {
									Text     string `xml:",chardata"`
									Code     string `xml:"Code,attr"`
									Quantity string `xml:"Quantity,attr"`
								} `xml:"PassengerTypeQuantity"`
								FareBasisCodes struct {
									Text          string `xml:",chardata"`
									FareBasisCode []struct {
										Text string `xml:",chardata"`
									} `xml:"FareBasisCode"`
								} `xml:"FareBasisCodes"`
								PassengerFare struct {
									Text     string `xml:",chardata"`
									BaseFare struct {
										Text          string `xml:",chardata"`
										Amount        string `xml:"Amount,attr"`
										CurrencyCode  string `xml:"CurrencyCode,attr"`
										DecimalPlaces string `xml:"DecimalPlaces,attr"`
									} `xml:"BaseFare"`
									Taxes struct {
										Text string `xml:",chardata"`
										Tax  []struct {
											Text          string `xml:",chardata"`
											TaxCode       string `xml:"TaxCode,attr"`
											Amount        string `xml:"Amount,attr"`
											CurrencyCode  string `xml:"CurrencyCode,attr"`
											DecimalPlaces string `xml:"DecimalPlaces,attr"`
										} `xml:"Tax"`
									} `xml:"Taxes"`
									TotalFare struct {
										Text          string `xml:",chardata"`
										Amount        string `xml:"Amount,attr"`
										CurrencyCode  string `xml:"CurrencyCode,attr"`
										DecimalPlaces string `xml:"DecimalPlaces,attr"`
									} `xml:"TotalFare"`
									UnstructuredFareCalc struct {
										Text string `xml:",chardata"`
									} `xml:"UnstructuredFareCalc"`
								} `xml:"PassengerFare"`
							} `xml:"PTC_FareBreakdown"`
						} `xml:"PTC_FareBreakdowns"`
					} `xml:"AirItineraryPricingInfo"`
				} `xml:"PricedItinerary"`
			} `xml:"PricedItineraries"`
			Errors struct {
				Text  string `xml:",chardata"`
				Error []struct {
					Type      string `xml:"Type,attr"`
					ShortText string `xml:"ShortText,attr"`
					Text      string `xml:",chardata"`
				}
			} `xml:"Errors"`
			FreeBaggageList struct {
				Text        string `xml:",chardata"`
				FreeBaggage []struct {
					Text          string `xml:",chardata"`
					PassengerType struct {
						Text string `xml:",chardata"`
					} `xml:"PassengerType"`
					PassengerCount struct {
						Text string `xml:",chardata"`
					} `xml:"PassengerCount"`
					BaggageList struct {
						Text             string `xml:",chardata"`
						FreeBaggageRoute []struct {
							Text  string `xml:",chardata"`
							Route struct {
								Text string `xml:",chardata"`
							} `xml:"Route"`
							FlightOperator struct {
								Text string `xml:",chardata"`
							} `xml:"FlightOperator"`
							Quantity struct {
								Text string `xml:",chardata"`
							} `xml:"Quantity"`
						} `xml:"FreeBaggageRoute"`
					} `xml:"BaggageList"`
				} `xml:"FreeBaggage"`
			} `xml:"FreeBaggageList"`
			FreeTextType struct {
				Text string `xml:",chardata"`
			} `xml:"FreeTextType"`
		} `xml:"OTA_AirPriceRS"`
	} `xml:"AirSellRQResponse"`
}
