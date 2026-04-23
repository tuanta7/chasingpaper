package plan

import (
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	IsActive       bool      `json:"is_active"`
	Prices         []Price   `json:"prices,omitempty"`
	Provider       string    `json:"provider"`
	ProviderPlanID string    `json:"provider_plan_id"` // provider-specific
	CreatedAt      time.Time `json:"created_at"`
}

type Price struct {
	PlanID         uuid.UUID  `json:"plan_id"`
	CachedResponse []byte     `json:"cached_response"` // JSONB
	LastSyncedAt   *time.Time `json:"last_synced_at"`
	CreatedAt      time.Time  `json:"created_at"`
}
