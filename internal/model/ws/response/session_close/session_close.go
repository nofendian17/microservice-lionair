package session_close

import "encoding/xml"

type SessionCloseXMLResponse struct {
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
			ConversationId string `xml:"ConversationId"` // 4890ab30-bcf8-43fc-9f55-8...
			Service        string `xml:"Service"`        // Logoff
			Action         string `xml:"Action"`         // SessionClose
			MessageData    struct {
				Text      string `xml:",chardata"`
				MessageId string `xml:"MessageId"` // mid:13:30:03.161@vedaleon...
				Timestamp string `xml:"Timestamp"` // 2023-05-30T10:34:07Z
			} `xml:"MessageData"`
		} `xml:"MessageHeader"`
		Security struct {
			Text                string `xml:",chardata"`
			Xmlns               string `xml:"xmlns,attr"`
			BinarySecurityToken string `xml:"BinarySecurityToken"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		Text           string `xml:",chardata"`
		LogoffResponse struct {
			Text  string `xml:",chardata"`
			Xmlns string `xml:"xmlns,attr"`
		} `xml:"LogoffResponse"`
	} `xml:"Body"`
}
