package plan

type Provider string

const (
	ProviderStripe Provider = "stripe"
	ProviderPayPal Provider = "paypal"
)

type Plan struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Provider    Provider `json:"provider"`    // denormalized
	ExternalID  string   `json:"external_id"` // provider-specific
	IsActive    bool     `json:"is_active"`
	CreatedAt   string   `json:"created_at"`
}
