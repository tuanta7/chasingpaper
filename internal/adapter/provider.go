package adapter

type Provider interface {
	PlanProvider
	CheckoutProvider
}
