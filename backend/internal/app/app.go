package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TheQuorix/Personal-Website/internal/adapter/github"
	"github.com/TheQuorix/Personal-Website/internal/adapter/lastfm"
	"github.com/TheQuorix/Personal-Website/internal/adapter/openweather"
	"github.com/TheQuorix/Personal-Website/internal/adapter/sqlite"
	"github.com/TheQuorix/Personal-Website/internal/adapter/steam"
	"github.com/TheQuorix/Personal-Website/internal/adapter/telegram"
	"github.com/TheQuorix/Personal-Website/internal/config"
	"github.com/TheQuorix/Personal-Website/internal/domain/comment"
	"github.com/TheQuorix/Personal-Website/internal/domain/commentrequest"
	"github.com/TheQuorix/Personal-Website/internal/domain/info"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	httpDelivery "github.com/TheQuorix/Personal-Website/internal/delivery/http"
	tgDelivery "github.com/TheQuorix/Personal-Website/internal/delivery/telegram"

	_ "modernc.org/sqlite"
)

var (
	Database *sql.DB
	BotAPI   *tgbotapi.BotAPI
	Config   config.Config
	Context  context.Context
	Cancel   context.CancelFunc

	HttpClient        *http.Client
	OpenWeatherClient *openweather.Client
	LastFmClient      *lastfm.Client
	GithubClient      *github.Client
	SteamClient       *steam.Client

	CommentService        *comment.Service
	CommentRequestService *commentrequest.Service

	InfoPoller *info.Poller

	TgBot      *tgDelivery.Bot
	HttpServer *httpDelivery.Server
)

// Run — запуск основного кода
func Run() {
	Config = config.Load()

	Context, Cancel = context.WithCancel(context.Background())
	defer Cancel()

	initAdapters()
	initDomain()
	initDelivery()

	go func() {
		log.Println("Telegram bot started")
		TgBot.Start()
	}()

	go func() {
		log.Printf("HTTP-server started at %s", Config.Port)
		if err := HttpServer.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start HTTP-server: %v", err)
		}
	}()

	// Не позволяет выключаться программе без сигнала выключения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	shutdown()
}

func initAdapters() {
	var err error

	Database, err = sqlite.Init(Config.DatabasePath)
	if err != nil {
		log.Fatalf("init database: %v", err)
	}

	BotAPI, err = tgbotapi.NewBotAPI(Config.TelegramBotToken)
	if err != nil {
		log.Fatalf("init telegram bot: %v", err)
	}

	HttpClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	OpenWeatherClient = openweather.NewClient(HttpClient, Config)
	LastFmClient = lastfm.NewClient(HttpClient, Config)
	GithubClient = github.NewClient(HttpClient, Config)

	imageCache, err := steam.NewImageCache("./data/steam-icons", "/media/steam-icons", HttpClient)
	if err != nil {
		log.Fatalf("init steam image cache: %v", err)
	}
	SteamClient = steam.NewClient(HttpClient, Config, imageCache)
}

func initDomain() {
	commentRepo := sqlite.NewCommentRepo(Database)
	commentRequestRepo := sqlite.NewCommentRequestRepo(Database)
	notifier := telegram.NewNotifier(BotAPI, Config.TelegramChatID)

	CommentService = comment.NewService(commentRepo)
	CommentRequestService = commentrequest.NewService(commentRequestRepo, notifier, commentRepo)

	InfoPoller = info.NewPoller(OpenWeatherClient, LastFmClient, SteamClient, GithubClient)
	InfoPoller.StartPolling(Context)
}

func initDelivery() {
	// Telegram-деливери
	fsmHandler := tgDelivery.NewFsmHandler(Context, CommentRequestService)
	msgHandler := tgDelivery.NewMessageHandler(CommentService, fsmHandler)
	callbackHandler := tgDelivery.NewCallbackHandler(Context, CommentRequestService, fsmHandler)
	tgRouter := tgDelivery.NewRouter(msgHandler, callbackHandler)
	TgBot = tgDelivery.NewBot(BotAPI, tgRouter)

	// HTTP-деливери
	commentHandler := httpDelivery.NewCommentHandler(Context, CommentService)
	commentRequestHandler := httpDelivery.NewCommentRequestHandler(CommentRequestService)

	httpRouter := httpDelivery.NewRouter(InfoPoller, commentHandler, commentRequestHandler)
	HttpServer = httpDelivery.NewServer(Config.Port, httpRouter)
}

func shutdown() {
	log.Println("Shutting down...")

	Cancel()

	ctx, timeoutCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeoutCancel()

	if err := HttpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	BotAPI.StopReceivingUpdates()

	if err := Database.Close(); err != nil {
		log.Printf("database close error: %v", err)
	}

	log.Println("Exited")
}
