package lion

import (
	"context"
	"encoding/xml"
	"fmt"
	wsSessionCloseRequest "lion/internal/model/ws/request/session_close"
	"time"
)

type SessionClose interface {
	Logout(ctx context.Context, conversationID, binarySecurityToken string) error
}

type sessionClose struct {
	lionRequest Request
}

func (s *sessionClose) Logout(ctx context.Context, conversationID, binarySecurityToken string) error {
	timestamp := time.Now().Format("2006-01-02T15:04:05Z")
	messageID := fmt.Sprintf("mid:%d", time.Now().UnixNano())
	sessionCloseRequest := &wsSessionCloseRequest.SessionCloseXMLRequest{}

	sessionCloseRequest.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
	sessionCloseRequest.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	sessionCloseRequest.Xsd = "http://www.w3.org/2001/XMLSchema"
	sessionCloseRequest.Header.MessageHeader.Xmlns = "http://www.ebxml.org/namespaces/messageHeader"
	sessionCloseRequest.Header.MessageHeader.CPAId = "JT"
	sessionCloseRequest.Header.MessageHeader.ConversationId = conversationID
	sessionCloseRequest.Header.MessageHeader.Service = "Logoff"
	sessionCloseRequest.Header.MessageHeader.Action = "SessionClose"
	sessionCloseRequest.Header.MessageHeader.MessageData.MessageId = messageID
	sessionCloseRequest.Header.MessageHeader.MessageData.Timestamp = timestamp
	sessionCloseRequest.Header.Security.Xmlns = "http://schemas.xmlsoap.org/ws/2002/12/secext"
	sessionCloseRequest.Header.Security.BinarySecurityToken = binarySecurityToken
	sessionCloseRequest.Body.Logoff.Xmlns = "http://www.vedaleon.com/webservices"

	sessionCloseXMLRequest, err := xml.MarshalIndent(sessionCloseRequest, "", "    ")
	if err != nil {
		return err
	}

	_, err = s.lionRequest.SessionClose(ctx, string(sessionCloseXMLRequest))
	if err != nil {
		return err
	}
	return nil
}

func NewSessionClose(lionRequest Request) SessionClose {
	return &sessionClose{
		lionRequest: lionRequest,
	}
}
