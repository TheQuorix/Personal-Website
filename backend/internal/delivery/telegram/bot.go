package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	api    *tgbotapi.BotAPI
	router *Router
}

func NewBot(api *tgbotapi.BotAPI, router *Router) *Bot {
	return &Bot{
		api:    api,
		router: router,
	}
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.api.GetUpdatesChan(u)

	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		b.router.Route(b.api, update)
	}
}
