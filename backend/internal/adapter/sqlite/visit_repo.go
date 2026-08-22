package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/TheQuorix/Personal-Website/internal/domain/visit"
)

type VisitRepo struct {
	db *sql.DB
}

func NewVisitRepo(db *sql.DB) *VisitRepo {
	return &VisitRepo{db: db}
}

func (r *VisitRepo) Insert(ctx context.Context, path, visitorHash string) error {
	visitedAt := time.Now().Format("2006-01-02")
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO visits (path, visitor_hash, visited_at) VALUES (?, ?, ?)`,
		path, visitorHash, visitedAt,
	)
	return err
}

func (r *VisitRepo) CountTotal(ctx context.Context) (int, error) {
	var total int
	err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM visits`).Scan(&total)
	return total, err
}

func (r *VisitRepo) CountUnique(ctx context.Context) (int, error) {
	var unique int
	err := r.db.QueryRowContext(ctx, `SELECT COUNT(DISTINCT visitor_hash) FROM visits`).Scan(&unique)
	return unique, err
}

func (r *VisitRepo) DailyStats(ctx context.Context, limit int) ([]visit.DailyStat, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT visited_at, COUNT(*), COUNT(DISTINCT visitor_hash)
		FROM visits
		GROUP BY visited_at
		ORDER BY visited_at DESC
		LIMIT ?
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []visit.DailyStat
	for rows.Next() {
		var d visit.DailyStat
		if err := rows.Scan(&d.Date, &d.Total, &d.Unique); err != nil {
			return nil, err
		}
		result = append(result, d)
	}
	return result, rows.Err()
}
