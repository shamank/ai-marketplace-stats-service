package aiservice

import (
	"context"
	"github.com/shamank/ai-marketplace-stats-service/internal/domain/models"
	"log/slog"
)

type AIServiceRepository interface {
	Create(ctx context.Context, service models.AIServiceCreate) (string, error)
}

type AIService struct {
	log  *slog.Logger
	repo AIServiceRepository
}

func NewAIService(log *slog.Logger, repo AIServiceRepository) *AIService {
	return &AIService{
		log:  log,
		repo: repo,
	}
}

func (s *AIService) Create(ctx context.Context, service models.AIServiceCreate) (string, error) {

	uid, err := s.repo.Create(ctx, service)
	if err != nil {
		return "", err
	}
	return uid, nil
}
