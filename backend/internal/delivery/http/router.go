package http

import (
	"encoding/json"
	"net/http"

	"github.com/TheQuorix/Personal-Website/internal/domain/info"
	"github.com/rs/cors"
)

// Создание роутеров
func NewRouter(infoPoller *info.Poller, commentHandler *CommentHandler, commentRequestHandler *CommentRequestHandler, visitHandler *VisitHandler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/media/steam-icons/", http.StripPrefix("/media/steam-icons/", http.FileServer(http.Dir("./data/steam-icons"))))

	mux.HandleFunc("GET /api/v1/test", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MessageResponse{
			Message: "Hello",
		})
	})

	mux.HandleFunc("GET /api/v1/info", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, infoPoller.GetCached())
	})

	mux.HandleFunc("POST /api/v1/comments", commentRequestHandler.HandleCreate)
	mux.HandleFunc("GET /api/v1/comments", commentHandler.HandleGet)

	mux.HandleFunc("POST /api/v1/visits", visitHandler.Track)
	mux.HandleFunc("GET /api/v1/visits/stats", visitHandler.Stats)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	return c.Handler(mux)
}

// Ответ в виде JSON
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
