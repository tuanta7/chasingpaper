package adapter

type PaymentProvider interface {
	PlanProvider
	CheckoutProvider
}

type PlanProvider interface {
}

type CheckoutProvider interface{}
