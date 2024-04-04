package statistic

import (
	"context"
	"github.com/shamank/ai-marketplace-stats-service/internal/domain/models"
	"log/slog"
)

type StatisticRepository interface {
	CreateService(ctx context.Context, input models.AIServiceCreate) (string, error)
	GetCalls(ctx context.Context, filter models.StatisticFilter) ([]models.StatisticRead, error)
	Call(ctx context.Context, statsWrite models.StatisticWrite) error
}

type StatisticService struct {
	log  *slog.Logger
	repo StatisticRepository
}

func NewStatisticService(log *slog.Logger, repo StatisticRepository) *StatisticService {
	return &StatisticService{
		log:  log,
		repo: repo,
	}
}

func (s *StatisticService) CreateService(ctx context.Context, input models.AIServiceCreate) (string, error) {

	return s.repo.CreateService(ctx, input)
}

func (s *StatisticService) GetCalls(ctx context.Context, filter models.StatisticFilter) ([]models.StatisticRead, error) {
	stats, err := s.repo.GetCalls(ctx, filter)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func (s *StatisticService) Call(ctx context.Context, userUID string, serviceUID string) error {
	statsWrite := models.StatisticWrite{
		UserUID:      userUID,
		AIServiceUID: serviceUID,
	}

	return s.repo.Call(ctx, statsWrite)
}
