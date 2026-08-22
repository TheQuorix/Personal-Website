package http

import (
	"log"
	"net/http"

	"github.com/TheQuorix/Personal-Website/internal/domain/visit"
)

type VisitHandler struct {
	service *visit.Service
}

func NewVisitHandler(service *visit.Service) *VisitHandler {
	return &VisitHandler{service: service}
}

func (h *VisitHandler) Track(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.RemoteAddr
	}
	ua := r.UserAgent()
	path := r.URL.Query().Get("p")

	if err := h.service.Track(r.Context(), path, ip, ua); err != nil {
		log.Printf("track visit error: %v", err)
		writeJSON(w, http.StatusInternalServerError, MessageResponse{Message: "internal error"})
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

func (h *VisitHandler) Stats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.service.GetStats(r.Context())
	if err != nil {
		log.Printf("get stats error: %v", err)
		writeJSON(w, http.StatusInternalServerError, MessageResponse{Message: "internal error"})
		return
	}
	writeJSON(w, http.StatusOK, stats)
}
