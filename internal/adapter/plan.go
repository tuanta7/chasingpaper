package adapter

import (
	"context"
	"time"
)

type PlanProvider interface {
	ListPlans(context.Context) ([]Plan, error)
	CreatePlan(context.Context, *Plan) error
	GetPlan(context.Context, string) (*Plan, error)
	UpdatePlanPrice(context.Context, *Price) error
	DeletePlan(context.Context, string) error
}

type Plan struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreateAt    time.Time `json:"create_at"`
}

type Price struct {
	Interval string `json:"interval"`
}
