package lion

import (
	"context"
	"errors"
	"fmt"
	"io"
	"lion/internal/shared/config"
	"lion/internal/shared/infrastructure/httpclient"
	"lion/internal/shared/infrastructure/soapclient"
	"net/http"
)

// Request is the interface that defines the available methods.
type Request interface {
	Ping(ctx context.Context) error
	SessionCreate(ctx context.Context, request string) (string, error)
	SessionClose(ctx context.Context, request string) (string, error)
	FlightMatrix(ctx context.Context, request string) (string, error)
	OTAAirSell(ctx context.Context, request string) (string, error)
}

type request struct {
	config     *config.Config
	httpClient httpclient.Client
	soapClient soapclient.SOAPClient
}

// NewRequest creates a new instance of the Request interface.
func NewRequest(config *config.Config, httpClient httpclient.Client, soapClient soapclient.SOAPClient) Request {
	return &request{
		config:     config,
		httpClient: httpClient,
		soapClient: soapClient,
	}
}

// Ping sends a Ping request to the specified URL.
func (r *request) Ping(ctx context.Context) error {
	url := r.config.Integration.Url
	headers := map[string]string{
		"Accept": "*",
	}

	_, err := r.httpClient.Get(ctx, url, headers)
	return err
}

func (r *request) SessionCreate(ctx context.Context, request string) (string, error) {
	if r.config.Integration.Url == "" {
		return "", errors.New("the provided URL is empty")
	}

	url := fmt.Sprintf("%s%s", r.config.Integration.Url, r.config.Integration.Service.SessionCreate.Path)
	call, err := r.soapClient.MakeSOAPCall(ctx, url, request)
	if err != nil {
		return "", fmt.Errorf("failed to make SOAP call: %w", err)
	}
	body, err := r.getBody(call)
	if err != nil {
		return "", err
	}

	return body, nil

}

func (r *request) SessionClose(ctx context.Context, request string) (string, error) {
	if r.config.Integration.Url == "" {
		return "", errors.New("the provided URL is empty")
	}

	url := fmt.Sprintf("%s%s", r.config.Integration.Url, r.config.Integration.Service.SessionClose.Path)
	call, err := r.soapClient.MakeSOAPCall(ctx, url, request)
	if err != nil {
		return "", fmt.Errorf("failed to make SOAP call: %w", err)
	}
	body, err := r.getBody(call)
	if err != nil {
		return "", err
	}

	return body, nil
}

// FlightMatrix search a new flight schedule without session using specified XML template.
func (r *request) FlightMatrix(ctx context.Context, request string) (string, error) {
	if r.config.Integration.Url == "" {
		return "", errors.New("the provided URL is empty")
	}

	url := fmt.Sprintf("%s%s", r.config.Integration.Url, r.config.Integration.Service.FlightMatrix.Path)
	call, err := r.soapClient.MakeSOAPCall(ctx, url, request)
	if err != nil {
		return "", fmt.Errorf("failed to make SOAP call: %w", err)
	}
	body, err := r.getBody(call)
	if err != nil {
		return "", err
	}

	return body, nil
}

func (r *request) OTAAirSell(ctx context.Context, request string) (string, error) {
	if r.config.Integration.Url == "" {
		return "", errors.New("the provided URL is empty")
	}

	url := fmt.Sprintf("%s%s", r.config.Integration.Url, r.config.Integration.Service.OTAAirSell.Path)
	fmt.Println(request)
	call, err := r.soapClient.MakeSOAPCall(ctx, url, request)
	if err != nil {
		return "", fmt.Errorf("failed to make SOAP call: %w", err)
	}
	body, err := r.getBody(call)
	if err != nil {
		return "", err
	}
	return body, nil
}

// getBody extract raw body response
func (r *request) getBody(response *http.Response) (string, error) {
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	fmt.Println(string(body))
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return string(body), nil
}
