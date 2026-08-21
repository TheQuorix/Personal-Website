package sqlite

import (
	"context"
	"database/sql"

	"github.com/TheQuorix/Personal-Website/internal/domain/comment"
)

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

var _ comment.Repository = (*CommentRepo)(nil)

func (r *CommentRepo) Create(ctx context.Context, c comment.Comment) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO comments (author, message, response, date) VALUES (?, ?, ?, ?)",
		c.Author, c.Message, c.Response, c.Date,
	)
	return err
}

func (r *CommentRepo) List(ctx context.Context) ([]comment.Comment, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, author, message, response, date FROM comments",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []comment.Comment
	for rows.Next() {
		var c comment.Comment
		if err := rows.Scan(&c.ID, &c.Author, &c.Message, &c.Response, &c.Date); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, rows.Err()
}
