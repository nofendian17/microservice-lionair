package soapclient

import (
	"context"
	"lion/internal/shared/infrastructure/httpclient"
	"net/http"
	"strings"
)

// SOAPClient defines the interface for making SOAP calls.
type SOAPClient interface {
	MakeSOAPCall(ctx context.Context, url, payload string) (*http.Response, error)
}

// DefaultSOAPClient is the response implementation of the SOAPClient interface.
type DefaultSOAPClient struct {
	httpClient httpclient.Client
}

// NewSOAPClient creates a new instance of DefaultSOAPClient.
func NewSOAPClient(httpClient httpclient.Client) SOAPClient {
	return &DefaultSOAPClient{
		httpClient: httpClient,
	}
}

// MakeSOAPCall makes a SOAP call to the specified URL with the given action and payload.
func (c *DefaultSOAPClient) MakeSOAPCall(ctx context.Context, url, payload string) (*http.Response, error) {
	headers := httpclient.Headers{
		"Content-Type": "text/xml",
	}
	reader := strings.NewReader(payload)
	response, err := c.httpClient.Post(ctx, url, headers, reader)
	if err != nil {
		return nil, err
	}

	return response, nil
}
