package commentrequest

import (
	"context"
	"fmt"

	"github.com/TheQuorix/Personal-Website/internal/domain/comment"
)

// Структура сервиса
type Service struct {
	repo        Repository
	notifier    Notifier
	commentRepo comment.Repository
}

// Создание нового сервиса
func NewService(repo Repository, notifier Notifier, commentRepo comment.Repository) *Service {
	return &Service{repo: repo, notifier: notifier, commentRepo: commentRepo}
}

// Создание заявки на комментарий
func (s *Service) Create(ctx context.Context, r CommentRequest) error {
	if r.Publish {
		id, err := s.notifier.NotifyWithPublish(ctx, r)
		if err != nil {
			return fmt.Errorf("notify new comment: %w", err)
		}

		r.TelegramID = id
		if err := s.repo.Create(ctx, r); err != nil {
			return fmt.Errorf("create comment request: %w", err)
		}
	} else {
		if err := s.notifier.NotifyWithoutPublish(ctx, r); err != nil {
			return fmt.Errorf("notify new comment request: %w", err)
		}
	}

	return nil
}

// Обработка заявки на комментарий (принять/отклонить)
func (s *Service) Handle(ctx context.Context, telegramId int, approve bool, response string) error {
	req, err := s.repo.GetByTelegramID(ctx, telegramId)
	if err != nil {
		return fmt.Errorf("get comment request %d: %w", telegramId, err)
	}

	if err := s.repo.DeleteByTelegramID(ctx, telegramId); err != nil {
		return fmt.Errorf("delete comment request %d: %w", telegramId, err)
	}

	if !approve {
		return nil
	}

	c := comment.Comment{
		Author:   req.Author,
		Message:  req.Message,
		Response: response,
		Date:     req.Date,
	}
	if err := s.commentRepo.Create(ctx, c); err != nil {
		return fmt.Errorf("save approved comment: %w", err)
	}

	return nil
}
