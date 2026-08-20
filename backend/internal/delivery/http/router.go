package http

import (
	"encoding/json"
	"net/http"
)

// Создание роутеров
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/test", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MessageResponse{
			Message: "Hello",
		})
	})

	return mux
}

// Ответ в виде JSON
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
