package payment

import "github.com/tuanta7/chasingpaper/pkg/postgres"

type Repository struct {
	pgPool *postgres.InstrumentedPool
}
