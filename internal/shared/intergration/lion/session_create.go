package lion

import (
	"context"
	"encoding/xml"
	"fmt"
	wsSessionCreateRequest "lion/internal/model/ws/request/session_create"
	wsSessionCreateResponse "lion/internal/model/ws/response/session_create"
	"time"
)

type SessionCreate interface {
	GetBinarySecurityToken(ctx context.Context, conversationID, username, password, organization string) (string, error)
}

type sessionCreate struct {
	lionRequest Request
}

func (s *sessionCreate) GetBinarySecurityToken(ctx context.Context, conversationID, username, password, organization string) (string, error) {
	sessionCreateRequest := &wsSessionCreateRequest.SessionCreateXMLRequest{}
	timestamp := time.Now().Format("2006-01-02T15:04:05Z")

	sessionCreateRequest.Nm1 = "http://www.vedaleon.com/webservices"
	sessionCreateRequest.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
	sessionCreateRequest.Header.MessageHeader.Xmlns = "http://www.ebxml.org/namespaces/messageHeader"
	sessionCreateRequest.Header.MessageHeader.CPAId = "JT"
	sessionCreateRequest.Header.MessageHeader.ConversationId = conversationID
	sessionCreateRequest.Header.MessageHeader.Service = "Create"
	sessionCreateRequest.Header.MessageHeader.Action = "CreateSession"
	sessionCreateRequest.Header.MessageHeader.MessageData.Timestamp = timestamp
	sessionCreateRequest.Header.Security.Xmlns = "http://schemas.xmlsoap.org/ws/2002/12/secext"
	sessionCreateRequest.Header.Security.UsernameToken.Username = username
	sessionCreateRequest.Header.Security.UsernameToken.Password = password
	sessionCreateRequest.Header.Security.UsernameToken.Organization.Text = organization
	sessionCreateRequest.Body.Logon.Xmlns = "http://www.vedaleon.com/webservices"

	xmlData, err := xml.MarshalIndent(sessionCreateRequest, "", "    ")
	if err != nil {
		return "", err
	}

	sessionCreateResponse, err := s.lionRequest.SessionCreate(ctx, string(xmlData))
	if err != nil {
		return "", err
	}
	sessionCreateXMLResponse := &wsSessionCreateResponse.SessionCreateXMLResponse{}
	err = xml.Unmarshal([]byte(sessionCreateResponse), sessionCreateXMLResponse)
	if err != nil {
		return "", err
	}
	if errorText := sessionCreateXMLResponse.Body.LogonResponse.LogonResult; errorText != "OK" {
		return "", fmt.Errorf("%s", errorText)
	}
	return sessionCreateXMLResponse.Header.Security.BinarySecurityToken, nil
}

func NewSessionCreate(lionRequest Request) SessionCreate {
	return &sessionCreate{
		lionRequest: lionRequest,
	}
}
