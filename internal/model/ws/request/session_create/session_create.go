package session_create

import (
	"encoding/xml"
)

type SessionCreateXMLRequest struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Text    string   `xml:",chardata"`
	Nm1     string   `xml:"xmlns:nm1,attr"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Header  struct {
		Text          string `xml:",chardata"`
		MessageHeader struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			CPAId          string `xml:"CPAId"`          // JT
			ConversationId string `xml:"ConversationId"` // {{$timestamp}}
			Service        string `xml:"Service"`        // Create
			Action         string `xml:"Action"`         // CreateSession
			MessageData    struct {
				Text      string `xml:",chardata"`
				MessageId string `xml:"MessageId"`
				Timestamp string `xml:"Timestamp"`
			} `xml:"MessageData"`
		} `xml:"MessageHeader"`
		Security struct {
			Text                string `xml:",chardata"`
			Xmlns               string `xml:"xmlns,attr"`
			BinarySecurityToken string `xml:"BinarySecurityToken"`
			UsernameToken       struct {
				Text         string `xml:",chardata"`
				Username     string `xml:"Username"`
				Password     string `xml:"Password"`
				Organization struct {
					Text  string `xml:",chardata"` // JT
					Xmlns string `xml:"xmlns,attr"`
				} `xml:"Organization"`
			} `xml:"UsernameToken"`
		} `xml:"Security"`
	} `xml:"soap:Header"`
	Body struct {
		Text  string `xml:",chardata"`
		Logon struct {
			Text  string `xml:",chardata"`
			Xmlns string `xml:"xmlns,attr"`
		} `xml:"Logon"`
	} `xml:"soap:Body"`
}
