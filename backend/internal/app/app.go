package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TheQuorix/Personal-Website/internal/adapter/openweather"
	"github.com/TheQuorix/Personal-Website/internal/config"
	httpDelivery "github.com/TheQuorix/Personal-Website/internal/delivery/http"
)

// Запуск основного кода
func Run() {
	cfg := config.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создание http клиента
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Подключение и тест OpenWeather клиента
	openweatherClient := openweather.NewClient(httpClient, cfg)

	weather, err := openweatherClient.Fetch(ctx)
	if err != nil {
		panic(fmt.Errorf("get weather: %w", err))
	}

	fmt.Printf("Temp: %v\nFeels like: %v\n", weather.Temp, weather.FeelsLike)

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
