package stripe

import (
	"context"

	"github.com/stripe/stripe-go/v84"
)

type Client struct {
	client *stripe.Client
}

func (c *Client) CreatePaymentLink(ctx context.Context) (string, error) {
	link, err := c.client.V1PaymentLinks.Create(ctx, &stripe.PaymentLinkCreateParams{})
	if err != nil {
		return "", err
	}

	return link.ID, nil
}
