package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheQuorix/Personal-Website/internal/config"
	httpDelivery "github.com/TheQuorix/Personal-Website/internal/delivery/http"
)

// Запуск основного кода
func Run() {
	cfg := config.Load()

	// Создание http сервера
	httpRouter := httpDelivery.NewRouter()
	httpServer := httpDelivery.NewServer(cfg.Port, httpRouter)

	go func() {
		log.Printf("HTTP-server started at %s", cfg.Port)
		if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("start HTTP-server: %w", err))
		}
	}()

	// Не позволяет выключаться программе без сигнала выключения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Exiting...")
}
