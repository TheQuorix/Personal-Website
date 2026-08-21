package commentrequest

import "context"

// Интерфейс репозитория
type Repository interface {
	Create(ctx context.Context, r CommentRequest) error
	GetByID(ctx context.Context, id int) (CommentRequest, error)
	GetByTelegramID(ctx context.Context, telegramId int) (CommentRequest, error)
	Delete(ctx context.Context, id int) error
	DeleteByTelegramID(ctx context.Context, telegramId int) error
}
