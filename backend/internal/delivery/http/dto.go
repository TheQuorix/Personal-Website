package http

import (
	"time"

	"github.com/TheQuorix/Personal-Website/internal/domain/comment"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateCommentRequest struct {
	Author  string `json:"author"`
	Message string `json:"message"`
	Publish bool   `json:"publish"`
}

type CommentResponse struct {
	Author   string    `json:"author"`
	Message  string    `json:"message"`
	Response string    `json:"response"`
	Date     time.Time `json:"date"`
}

func ToCommentResponseList(comments []comment.Comment) []CommentResponse {
	responses := make([]CommentResponse, 0, len(comments))
	for _, c := range comments {
		responses = append(responses, CommentResponse{
			Author:   c.Author,
			Message:  c.Message,
			Response: c.Response,
			Date:     c.Date,
		})
	}
	return responses
}
