package visit

import "context"

type Repository interface {
	Insert(ctx context.Context, path, visitorHash string) error
	CountTotal(ctx context.Context) (int, error)
	CountUnique(ctx context.Context) (int, error)
	DailyStats(ctx context.Context, limit int) ([]DailyStat, error)
}
