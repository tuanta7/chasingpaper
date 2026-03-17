package paypal

import "net/http"

const (
	baseSandboxURL = "https://api-m.sandbox.paypal.com"
	baseURL        = "https://api-m.paypal.com"

	plansPath         = "/v2/plans"
	subscriptionsPath = "/v1/billing/subscriptions"
)

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}
