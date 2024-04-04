package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/shamank/ai-marketplace-stats-service/internal/repository/postgres"
	"github.com/shamank/ai-marketplace-stats-service/internal/service/statistic"
)

type Repositories struct {
	db *pgx.Conn

	StatisticRepo statistic.StatisticRepository
}

func NewRepositories(db *pgx.Conn) *Repositories {
	return &Repositories{
		db: db,

		StatisticRepo: postgres.NewStatisticRepository(db),
	}
}
