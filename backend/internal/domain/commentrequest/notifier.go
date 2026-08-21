package commentrequest

import "context"

// Интерфейс уведомлений
type Notifier interface {
	NotifyWithoutPublish(ctx context.Context, r CommentRequest) error
	NotifyWithPublish(ctx context.Context, r CommentRequest) (int, error)
}
