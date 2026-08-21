package telegram

import (
	"github.com/TheQuorix/Personal-Website/internal/domain/comment"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MessageHandler struct {
	service    *comment.Service
	fsmHandler *FsmHandler
}

func NewMessageHandler(service *comment.Service, fsmHandler *FsmHandler) *MessageHandler {
	return &MessageHandler{
		service:    service,
		fsmHandler: fsmHandler,
	}
}

func (h *MessageHandler) HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if message == nil {
		return
	}

	chatID := message.Chat.ID

	if state, active := h.fsmHandler.GetActiveState(chatID); active {
		h.fsmHandler.HandleFSM(bot, message, state)
		return
	}

	if message.IsCommand() && message.Command() == "start" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Привет!")
		bot.Send(msg)
	}
}
