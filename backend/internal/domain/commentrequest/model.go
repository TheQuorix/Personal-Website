package commentrequest

import "time"

// Структура заявок комментария
type CommentRequest struct {
	ID         int
	Author     string
	Message    string
	TelegramID int
	Publish    bool
	Date       time.Time
}
