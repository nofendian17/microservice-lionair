package httpclient

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

// Client defines the interface for making HTTP requests.
type Client interface {
	Get(ctx context.Context, url string, headers Headers) (*http.Response, error)
	Post(ctx context.Context, url string, headers Headers, body io.Reader) (*http.Response, error)
	Put(ctx context.Context, url string, headers Headers, body io.Reader) (*http.Response, error)
	Patch(ctx context.Context, url string, headers Headers, body io.Reader) (*http.Response, error)
	Delete(ctx context.Context, url string, headers Headers) (*http.Response, error)
}

// Headers represents the HTTP headers.
type Headers map[string]string

// DefaultClient is the response implementation of the Client interface.
type DefaultClient struct {
	client *http.Client
}

// NewDefaultClient creates a new instance of DefaultClient.
func NewDefaultClient(timeout time.Duration, insecureSkipVerify bool) Client {
	return &DefaultClient{
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecureSkipVerify,
				},
			},
		},
	}
}

// Get performs a GET request with custom headers.
func (c *DefaultClient) Get(ctx context.Context, url string, headers Headers) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

// Post performs a POST request with custom headers and a request body.
func (c *DefaultClient) Post(ctx context.Context, url string, headers Headers, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

// Put performs a PUT request with custom headers and a request body.
func (c *DefaultClient) Put(ctx context.Context, url string, headers Headers, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

// Patch performs a PATCH request with custom headers and a request body.
func (c *DefaultClient) Patch(ctx context.Context, url string, headers Headers, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

// Delete performs a DELETE request with custom headers.
func (c *DefaultClient) Delete(ctx context.Context, url string, headers Headers) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}
