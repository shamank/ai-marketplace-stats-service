package statistic

import (
	"context"
	"fmt"
	"github.com/shamank/ai-marketplace-stats-service/internal/domain/models"
	"log/slog"
)

type StatisticRepository interface {
	GetStats(ctx context.Context, filter models.StatisticFilter) ([]models.StatisticRead, error)
	SetStat(ctx context.Context, userUID string, serviceUID string) (string, error)
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

func (s *StatisticService) GetStats(ctx context.Context, filter models.StatisticFilter) ([]models.StatisticRead, error) {

	stats, err := s.repo.GetStats(ctx, filter)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func (s *StatisticService) SetStat(ctx context.Context, userUID string, serviceUID string) (string, error) {

	uid, err := s.repo.SetStat(ctx, userUID, serviceUID)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return uid, nil
}
