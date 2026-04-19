package paypal

import (
	"net/http"
)

const (
	baseSandboxURL = "https://api-m.sandbox.paypal.com"
	baseURL        = "https://api-m.paypal.com"

	pathV2Plans         = "/v2/plans"
	pathV2Subscriptions = "/v2/billing/subscriptions"
)

type Client struct {
	baseURL      string
	clientID     string
	clientSecret string
	httpClient   *http.Client
}

func NewClient(
	clientID, clientSecret string,
	httpClient *http.Client,
) *Client {
	return &Client{
		baseURL:      baseURL,
		clientID:     clientID,
		clientSecret: clientSecret,
		httpClient:   httpClient,
	}
}

func NewSandboxClient(
	clientID, clientSecret string,
	httpClient *http.Client,
) *Client {
	return &Client{
		baseURL:      baseSandboxURL,
		clientID:     clientID,
		clientSecret: clientSecret,
		httpClient:   httpClient,
	}
}
