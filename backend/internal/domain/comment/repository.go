package comment

import "context"

// Структура репозитория
type Repository interface {
	Create(ctx context.Context, c Comment) error
	List(ctx context.Context) ([]Comment, error)
}
