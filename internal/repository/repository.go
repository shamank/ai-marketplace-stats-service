package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/shamank/ai-marketplace-stats-service/internal/repository/postgres"
	aiservice "github.com/shamank/ai-marketplace-stats-service/internal/service/ai-service"
	"github.com/shamank/ai-marketplace-stats-service/internal/service/statistic"
)

type Repositories struct {
	db *pgx.Conn

	AIServicesRepo aiservice.AIServiceRepository
	StatsRepo      statistic.StatisticRepository
}

func NewRepositories(db *pgx.Conn) *Repositories {
	return &Repositories{
		db: db,

		AIServicesRepo: postgres.NewAIServiceRepository(db),
		StatsRepo:      postgres.NewStatisticRepository(db),
	}
}
