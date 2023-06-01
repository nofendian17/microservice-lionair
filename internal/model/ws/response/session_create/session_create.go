package session_create

import (
	"encoding/xml"
)

type SessionCreateXMLResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Header  struct {
		Text          string `xml:",chardata"`
		MessageHeader struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			CPAId          string `xml:"CPAId"`          // JT
			ConversationId string `xml:"ConversationId"` // 1685442563
			Service        string `xml:"Service"`        // Create
			Action         string `xml:"Action"`         // CreateSession
			MessageData    struct {
				Text      string `xml:",chardata"`
				Timestamp string `xml:"Timestamp"` // 2023-05-30T10:29:22Z
			} `xml:"MessageData"`
		} `xml:"MessageHeader"`
		Security struct {
			Text                string `xml:",chardata"`
			Xmlns               string `xml:"xmlns,attr"`
			BinarySecurityToken string `xml:"BinarySecurityToken"` // ryfQk06ZpFIWRHXyis9QVIddo...
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		Text          string `xml:",chardata"`
		LogonResponse struct {
			Text        string `xml:",chardata"`
			Xmlns       string `xml:"xmlns,attr"`
			LogonResult string `xml:"LogonResult"` // OK
		} `xml:"LogonResponse"`
	} `xml:"Body"`
}
