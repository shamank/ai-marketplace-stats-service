package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/shamank/ai-marketplace-stats-service/internal/domain/models"
	"github.com/shamank/ai-marketplace-stats-service/internal/service/statistic"
)

var _ statistic.StatisticRepository = (*StatisticRepository)(nil)

type StatisticRepository struct {
	db *pgx.Conn
}

func NewStatisticRepository(db *pgx.Conn) *StatisticRepository {
	return &StatisticRepository{
		db: db,
	}
}

func (r *StatisticRepository) CreateService(ctx context.Context, service models.AIServiceCreate) (string, error) {
	query := "insert into aiservices(title, description, price) VALUES ($1, $2, $3) returning uid"

	var serviceUID string

	if err := r.db.QueryRow(ctx, query, service.Title, service.Description, service.Price).Scan(&serviceUID); err != nil {
		return "", err
	}
	return serviceUID, nil
}

func (r *StatisticRepository) GetCalls(ctx context.Context, filter models.StatisticFilter) ([]models.StatisticRead, error) {

	query := "select s.user_uid, s.aiservice_uid, count(s.aiservice_uid), sum(s.amount) from statistics s where true"

	values := []interface{}{}
	argNumber := 0

	if filter.UserUID != nil {
		argNumber++
		values = append(values, *filter.UserUID)
		query += fmt.Sprintf(" and s.user_uid = $%d", argNumber)
	}

	if filter.AIServiceUID != nil {
		argNumber++
		values = append(values, *filter.AIServiceUID)
		query += fmt.Sprintf(" and s.aiservice_uid = $%d", argNumber)
	}

	query += " group by s.user_uid, s.aiservice_uid"

	if filter.Order != nil {
		if *filter.Order == "asc" || *filter.Order == "desc" {
			argNumber++
			values = append(values, *filter.Order)
			query += fmt.Sprintf(" order by s.created_at %d ", argNumber)
		}
	}

	if filter.PageSize != nil {
		argNumber++
		values = append(values, *filter.PageSize)
		query += fmt.Sprintf(" limit $%d ", argNumber)

		if filter.PageNumber != nil {
			argNumber++
			offset := (*filter.PageNumber - 1) * (*filter.PageSize)
			values = append(values, offset)
			query += fmt.Sprintf(" offset $%d ", argNumber)
		}
	}

	rows, err := r.db.Query(ctx, query, values...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]models.StatisticRead, 0)

	for rows.Next() {
		var stat models.StatisticRead
		if err := rows.Scan(&stat.UserUID, &stat.AIServiceUID, &stat.Count, &stat.FullAmount); err != nil {
			return nil, err
		}
		result = append(result, stat)
	}

	return result, nil
}

func (r *StatisticRepository) Call(ctx context.Context, statsWrite models.StatisticWrite) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {

		return err
	}

	var price float64

	if err := tx.QueryRow(ctx, "select price from aiservices where uid = $1", statsWrite.AIServiceUID).Scan(&price); err != nil {
		return err
	}

	query := "insert into statistics(user_uid, aiservice_uid, amount) VALUES ($1, $2, $3) returning uid"

	var statUID string

	if err := tx.QueryRow(ctx, query, statsWrite.UserUID, statsWrite.AIServiceUID, price).Scan(&statUID); err != nil {
		fmt.Println(err)
		return err
	}

	return tx.Commit(ctx)
}
