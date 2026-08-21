package http

import (
	"context"
	"net/http"

	"github.com/TheQuorix/Personal-Website/internal/domain/comment"
)

type CommentHandler struct {
	ctx     context.Context
	service *comment.Service
}

func NewCommentHandler(ctx context.Context, service *comment.Service) *CommentHandler {
	return &CommentHandler{
		ctx:     ctx,
		service: service,
	}
}

func (h *CommentHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "Method not allowed"})
		return
	}

	comments, err := h.service.List(h.ctx)

	if err != nil {
		// ToDo
	}

	// 2. Преобразуем слайс домена в слайс DTO
	responseList := ToCommentResponseList(comments)

	// 3. Отправляем JSON-массив
	writeJSON(w, http.StatusOK, responseList)
}
