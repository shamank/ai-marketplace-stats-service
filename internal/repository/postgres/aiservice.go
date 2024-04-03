package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/shamank/ai-marketplace-stats-service/internal/domain/models"
	aiservice "github.com/shamank/ai-marketplace-stats-service/internal/service/ai-service"
)

var _ aiservice.AIServiceRepository = (*AIServiceRepository)(nil)

type AIServiceRepository struct {
	db *pgx.Conn
}

func NewAIServiceRepository(db *pgx.Conn) *AIServiceRepository {
	return &AIServiceRepository{
		db: db,
	}
}

func (r *AIServiceRepository) Create(ctx context.Context, service models.AIServiceCreate) (string, error) {
	query := "insert into aiservices(title, description, current_price) VALUES ($1, $2, $3) returning uid"

	var serviceUID string

	if err := r.db.QueryRow(ctx, query, service.Title, service.Description, service.Price).Scan(&serviceUID); err != nil {
		return "", err
	}
	return serviceUID, nil
}
