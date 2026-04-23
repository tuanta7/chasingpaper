package adapter

import "context"

type CheckoutProvider interface {
	CreateSubscription(context.Context, *Subscription) error
}

type Subscription struct{}
