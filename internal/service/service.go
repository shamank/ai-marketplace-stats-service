package service

import (
	"github.com/shamank/ai-marketplace-stats-service/internal/repository"
	aiservice "github.com/shamank/ai-marketplace-stats-service/internal/service/ai-service"
	"github.com/shamank/ai-marketplace-stats-service/internal/service/statistic"
	"log/slog"
)

type Service struct {
	log *slog.Logger

	AIService        *aiservice.AIService
	StatisticService *statistic.StatisticService
}

func NewService(log *slog.Logger, repos *repository.Repositories) *Service {
	return &Service{
		log:              log,
		AIService:        aiservice.NewAIService(log, repos.AIServicesRepo),
		StatisticService: statistic.NewStatisticService(log, repos.StatsRepo),
	}
}
