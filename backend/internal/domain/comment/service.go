package comment

import (
	"context"
	"fmt"
)

// Структура сервиса
type Service struct {
	repo Repository
}

// Создаем новый сервис
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Получение списка комментариев
func (s *Service) List(ctx context.Context) ([]Comment, error) {
	comments, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list comments: %w", err)
	}
	return comments, nil
}

// Создание нового комментария
func (s *Service) Create(ctx context.Context, c Comment) error {
	if err := s.repo.Create(ctx, c); err != nil {
		return fmt.Errorf("create comment: %w", err)
	}
	return nil
}
