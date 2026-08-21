package sqlite

import (
	"context"
	"database/sql"

	"github.com/TheQuorix/Personal-Website/internal/domain/commentrequest"
)

// Структура репозитория
type CommentRequestRepo struct {
	db *sql.DB
}

// Создание нового репозитория
func NewCommentRequestRepo(db *sql.DB) *CommentRequestRepo {
	return &CommentRequestRepo{db: db}
}

// Проверка реализации репозитории
var _ commentrequest.Repository = (*CommentRequestRepo)(nil)

// Создание заявки на комментарий
func (r *CommentRequestRepo) Create(ctx context.Context, req commentrequest.CommentRequest) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO comment_requests (author, message, telegram_id, date) VALUES (?, ?, ?, ?)",
		req.Author, req.Message, req.TelegramID, req.Date,
	)
	return err
}

// Получение заявки на комментарий по айди
func (r *CommentRequestRepo) GetByID(ctx context.Context, id int) (commentrequest.CommentRequest, error) {
	var req commentrequest.CommentRequest
	row := r.db.QueryRowContext(ctx,
		"SELECT id, author, message, telegram_id, date FROM comment_requests WHERE id = ?", id,
	)
	err := row.Scan(&req.ID, &req.Author, &req.Message, &req.TelegramID, &req.Date)
	return req, err
}

// Получение заявки на комментарий по айди сообщения
func (r *CommentRequestRepo) GetByTelegramID(ctx context.Context, telegramId int) (commentrequest.CommentRequest, error) {
	var req commentrequest.CommentRequest
	row := r.db.QueryRowContext(ctx,
		"SELECT id, author, message, telegram_id, date FROM comment_requests WHERE telegram_id = ?", telegramId,
	)
	err := row.Scan(&req.ID, &req.Author, &req.Message, &req.TelegramID, &req.Date)
	return req, err
}

// Удаление заявки на комментарий
func (r *CommentRequestRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM comment_requests WHERE id = ?", id)
	return err
}

// Удаление заявки на комментарий по айди сообщения
func (r *CommentRequestRepo) DeleteByTelegramID(ctx context.Context, telegramId int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM comment_requests WHERE telegram_id = ?", telegramId)
	return err
}
