package session_close

import (
	"encoding/xml"
)

type SessionCloseXMLRequest struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Header  struct {
		Text          string `xml:",chardata"`
		MessageHeader struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			CPAId          string `xml:"CPAId"`          // JT
			ConversationId string `xml:"ConversationId"` // {{$guid}}
			Service        string `xml:"Service"`        // Logoff
			Action         string `xml:"Action"`         // SessionClose
			MessageData    struct {
				Text      string `xml:",chardata"`
				MessageId string `xml:"MessageId"` // mid:13:30:03.161@vedaleon...
				Timestamp string `xml:"Timestamp"` // {{$timestamp}}
			} `xml:"MessageData"`
		} `xml:"MessageHeader"`
		Security struct {
			Text                string `xml:",chardata"`
			Xmlns               string `xml:"xmlns,attr"`
			BinarySecurityToken string `xml:"BinarySecurityToken"` // ryfQk06ZpFIWRHXyis9QVIddo...
		} `xml:"Security"`
	} `xml:"soap:Header"`
	Body struct {
		Text   string `xml:",chardata"`
		Logoff struct {
			Text  string `xml:",chardata"`
			Xmlns string `xml:"xmlns,attr"`
		} `xml:"Logoff"`
	} `xml:"soap:Body"`
}
