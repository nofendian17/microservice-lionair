package session_close

import (
	"encoding/xml"
	"fmt"
	"time"
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

func (r *SessionCloseXMLRequest) NewSessionClose(conversationID, binarySecurityToken string) (string, error) {
	timestamp := time.Now().Format("2006-01-02T15:04:05Z")
	messageID := fmt.Sprintf("mid:%d", time.Now().UnixNano())

	r.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
	r.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	r.Xsd = "http://www.w3.org/2001/XMLSchema"
	r.Header.MessageHeader.Xmlns = "http://www.ebxml.org/namespaces/messageHeader"
	r.Header.MessageHeader.CPAId = "JT"
	r.Header.MessageHeader.ConversationId = conversationID
	r.Header.MessageHeader.Service = "Logoff"
	r.Header.MessageHeader.Action = "SessionClose"
	r.Header.MessageHeader.MessageData.MessageId = messageID
	r.Header.MessageHeader.MessageData.Timestamp = timestamp
	r.Header.Security.Xmlns = "http://schemas.xmlsoap.org/ws/2002/12/secext"
	r.Header.Security.BinarySecurityToken = binarySecurityToken
	r.Body.Logoff.Xmlns = "http://www.vedaleon.com/webservices"

	xmlData, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}

	return string(xmlData), nil
}
