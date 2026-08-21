package http

import (
	"encoding/json"
	"net/http"

	"github.com/TheQuorix/Personal-Website/internal/domain/commentrequest"
)

type CommentRequestHandler struct {
	service *commentrequest.Service
}

func NewCommentRequestHandler(service *commentrequest.Service) *CommentRequestHandler {
	return &CommentRequestHandler{service: service}
}

// HandleCreate обрабатывает POST /api/v1/comments
func (h *CommentRequestHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	// 1. Проверяем HTTP-метод
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "Method not allowed"})
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	// 2. Декодируем JSON из тела запроса
	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON body"})
		return
	}

	// 3. Базовая валидация DTO
	if req.Message == "" || req.Author == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "text and author are required"})
		return
	}

	if len(req.Message) > 2000 {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "message too long"})
		return
	}

	request := commentrequest.CommentRequest{
		Author:  req.Author,
		Message: req.Message,
		Publish: req.Publish,
	}

	// 4. Вызов бизнес-логики (Domain Service)
	err := h.service.Create(r.Context(), request)

	if err != nil {
		// ToDo
	}

	resp := MessageResponse{
		Message: "Successfully created",
	}

	writeJSON(w, http.StatusCreated, resp)
}
