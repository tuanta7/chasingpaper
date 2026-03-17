package plan

import (
	"github.com/tuanta7/chasingpaper/pkg/postgres"
)

type PostgresRepository struct {
	pg postgres.Client
}

func NewPostgresRepository(pool postgres.Client) *PostgresRepository {
	return &PostgresRepository{
		pg: pool,
	}
}

func (r *PostgresRepository) GetPlanList() ([]*Plan, error) {

	return nil, nil
}
