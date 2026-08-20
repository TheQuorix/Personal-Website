package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Структура конфига
type Config struct {
	Port            string
	DatabasePath    string
	VisitorHashSalt string

	WeatherCity   string
	WeatherApiKey string

	LastFmUser   string
	LastFmApiKey string

	GithubUsername string
	GuthibToken    string

	SteamID           string
	SteamApiKey       string
	SteamGridDbApiKey string

	TelegramChatID   int64
	TelegramBotToken string
}

// Получение данных конфига
func Load() Config {
	godotenv.Load()

	chatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
	if err != nil {
		panic(fmt.Errorf("parse CHAT_ID: %w", err))
	}

	return Config{
		Port:            os.Getenv("PORT"),
		DatabasePath:    os.Getenv("DATABASE_PATH"),
		VisitorHashSalt: os.Getenv("VISITOR_HASH_SALT"),

		WeatherCity:   os.Getenv("WEATHER_CITY"),
		WeatherApiKey: os.Getenv("WEATHER_API_KEY"),

		LastFmUser:   os.Getenv("LAST_FM_USER"),
		LastFmApiKey: os.Getenv("LAST_FM_API_KEY"),

		GithubUsername: os.Getenv("GITHUB_USERNAME"),
		GuthibToken:    os.Getenv("GITHUB_TOKEN"),

		SteamID:           os.Getenv("STEAM_ID"),
		SteamApiKey:       os.Getenv("STEAM_API_KEY"),
		SteamGridDbApiKey: os.Getenv("STEAM_GRID_DB_API_KEY"),

		TelegramChatID:   chatID,
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}
}
