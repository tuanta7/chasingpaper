package invoice

import (
	"context"

	"github.com/tuanta7/chasingpaper/internal/provider/stripe"
)

type UseCase struct {
	stripe *stripe.Client
}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) CreatePaymentLink(ctx context.Context) error {
	u.stripe.CreatePaymentLink(ctx)
	return nil
}

func (u *UseCase) UpdatePaymentStatus(ctx context.Context) error {
	return nil
}
