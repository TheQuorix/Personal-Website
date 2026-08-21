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

	"github.com/TheQuorix/Personal-Website/internal/adapter/github"
	"github.com/TheQuorix/Personal-Website/internal/adapter/lastfm"
	"github.com/TheQuorix/Personal-Website/internal/adapter/openweather"
	"github.com/TheQuorix/Personal-Website/internal/adapter/steam"
	"github.com/TheQuorix/Personal-Website/internal/config"
	httpDelivery "github.com/TheQuorix/Personal-Website/internal/delivery/http"
	"github.com/TheQuorix/Personal-Website/internal/domain/info"
)

var Config config.Config
var Context context.Context

var HttpClient *http.Client
var OpenWeatherClient *openweather.Client
var LastFmClient *lastfm.Client
var GithubClient *github.Client
var SteamClient *steam.Client

var InfoPoller *info.Poller

// Запуск основного кода
func Run() {
	Config = config.Load()

	var cancel context.CancelFunc
	Context, cancel = context.WithCancel(context.Background())
	defer cancel()

	initAdapters()
	initDomain()
	initDelivery()

	// Не позволяет выключаться программе без сигнала выключения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Exiting...")
}

func initAdapters() {
	// Создание http клиента
	HttpClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	// Подключение и тест OpenWeather клиента
	OpenWeatherClient = openweather.NewClient(HttpClient, Config)

	// Подключение и тест LastFm клиента
	LastFmClient = lastfm.NewClient(HttpClient, Config)

	// Подключение и тест Github
	GithubClient = github.NewClient(HttpClient, Config)

	// Подключение Steam
	imageCache, err := steam.NewImageCache("./data/steam-icons", "/media/steam-icons", HttpClient)
	if err != nil {
		log.Fatal(err)
	}

	SteamClient = steam.NewClient(HttpClient, Config, imageCache)
}

func initDomain() {
	// Подключение и старт опроса информации
	InfoPoller = info.NewPoller(*OpenWeatherClient, *LastFmClient, *SteamClient, *GithubClient)
	InfoPoller.StartPolling(Context)
}

func initDelivery() {
	httpRouter := httpDelivery.NewRouter(InfoPoller)
	httpServer := httpDelivery.NewServer(Config.Port, httpRouter)

	go func() {
		log.Printf("HTTP-server started at %s", Config.Port)
		if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("start HTTP-server: %w", err))
		}
	}()
}
