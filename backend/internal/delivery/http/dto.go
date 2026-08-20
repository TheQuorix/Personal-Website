package http

import "time"

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
