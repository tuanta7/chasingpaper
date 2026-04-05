package invoice

import (
	"context"
)

type UseCase struct {
}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) CreatePaymentLink(ctx context.Context) error {
	return nil
}

func (u *UseCase) UpdatePaymentStatus(ctx context.Context) error {
	return nil
}
