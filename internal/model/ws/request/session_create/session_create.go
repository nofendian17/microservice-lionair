package session_create

import (
	"encoding/xml"
	"time"
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

func (r *SessionCreateXMLRequest) NewSessionCreate(conversationID, username, password, organization string) (string, error) {
	timestamp := time.Now().Format("2006-01-02T15:04:05Z")

	r.Nm1 = "http://www.vedaleon.com/webservices"
	r.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
	r.Header.MessageHeader.Xmlns = "http://www.ebxml.org/namespaces/messageHeader"
	r.Header.MessageHeader.CPAId = "JT"
	r.Header.MessageHeader.ConversationId = conversationID
	r.Header.MessageHeader.Service = "Create"
	r.Header.MessageHeader.Action = "CreateSession"
	r.Header.MessageHeader.MessageData.Timestamp = timestamp
	r.Header.Security.Xmlns = "http://schemas.xmlsoap.org/ws/2002/12/secext"
	r.Header.Security.UsernameToken.Username = username
	r.Header.Security.UsernameToken.Password = password
	r.Header.Security.UsernameToken.Organization.Text = organization
	r.Body.Logon.Xmlns = "http://www.vedaleon.com/webservices"

	xmlData, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}

	return string(xmlData), nil
}
