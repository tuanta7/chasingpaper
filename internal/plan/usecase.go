package plan

type Repository interface {
}

type UseCase struct {
	repo *PostgresRepository
}

func NewUseCase(repo *PostgresRepository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
