package service

import (
	"github.com/shamank/ai-marketplace-stats-service/internal/repository"
	"github.com/shamank/ai-marketplace-stats-service/internal/service/statistic"
	"log/slog"
)

type Service struct {
	log *slog.Logger

	StatisticService *statistic.StatisticService
}

func NewService(log *slog.Logger, repos *repository.Repositories) *Service {
	return &Service{
		log:              log,
		StatisticService: statistic.NewStatisticService(log, repos.StatisticRepo),
	}
}
