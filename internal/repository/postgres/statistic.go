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

func (r *StatisticRepository) GetStats(ctx context.Context, filter models.StatisticFilter) ([]models.StatisticRead, error) {

	query := "select s.user_uid, s.aiservice_uid, count(s.aiservice_uid), sum(s.amount) from stats s where true"

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
			values = append(values, *filter.PageNumber)
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

func (r *StatisticRepository) SetStat(ctx context.Context, userUID string, serviceUID string) (string, error) {

	tx, err := r.db.Begin(ctx)
	if err != nil {

		return "", err
	}

	var amount int

	if err := tx.QueryRow(ctx, "select current_price from aiservices where uid = $1", serviceUID).Scan(&amount); err != nil {
		return "", err
	}

	query := "insert into stats(user_uid, aiservice_uid, amount) VALUES ($1, $2, $3) returning uid"

	var statUID string

	if err := tx.QueryRow(ctx, query, userUID, serviceUID, amount).Scan(&statUID); err != nil {
		fmt.Println(err)
		return "", err
	}

	return statUID, tx.Commit(ctx)
}
