package plan

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tuanta7/chasingpaper/internal/adapter/paypal"
	"github.com/tuanta7/chasingpaper/internal/repository/store"
)

type Repository interface {
	ListPlans(ctx context.Context, arg store.ListPlansParams) ([]store.Plan, error)
	CreatePlan(ctx context.Context, arg store.CreatePlanParams) (store.Plan, error)
	GetPlan(ctx context.Context, id pgtype.UUID) (store.Plan, error)
}

type UseCase struct {
	provider *paypal.Client
	repo     Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) ListPlans(ctx context.Context, page, pageSize int32) ([]Plan, error) {
	dbPlans, err := u.repo.ListPlans(ctx, store.ListPlansParams{
		Offset: (page - 1) * pageSize,
		Limit:  pageSize,
	})
	if err != nil {
		return nil, err
	}

	plans := make([]Plan, len(dbPlans))
	for i, dbPlan := range dbPlans {
		plans[i] = Plan{
			ID:          dbPlan.ID.Bytes,
			Name:        dbPlan.Name,
			Description: dbPlan.Description,
		}
	}

	return plans, nil
}

func (u *UseCase) CreatePlan(ctx context.Context, plan Plan) error {
	var err error

	if plan.ID == uuid.Nil {
		plan.ID, err = uuid.NewV7()
		if err != nil {
			return err
		}
	}

	_, err = u.repo.CreatePlan(ctx, store.CreatePlanParams{
		ID:          pgtype.UUID{Bytes: plan.ID},
		Name:        plan.Name,
		Description: plan.Description,
		IsActive:    true,
	})

	return err
}

func (u *UseCase) ListPayPalPlans() error {
	return nil
}

func (u *UseCase) CreatePrice(ctx context.Context, provider, externalPriceID string) error {
	return nil
}

func (u *UseCase) GetPlan() (*Plan, error) {
	return nil, nil
}
