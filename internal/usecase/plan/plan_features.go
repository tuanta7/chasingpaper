package plan

import "github.com/google/uuid"

type Features struct {
	PlanID     uuid.UUID `json:"plan_id"`
	FeatureKey string    `json:"feature_key"`
	Enabled    bool      `json:"enabled"`
}

type Feature struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Endpoint    string `json:"endpoint"`
}
