package telegram

import (
	"context"
	"strings"

	"github.com/TheQuorix/Personal-Website/internal/domain/commentrequest"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CallbackHandler struct {
	ctx            context.Context
	requestService *commentrequest.Service
	fsmHandler     *FsmHandler
}

func NewCallbackHandler(ctx context.Context, requestService *commentrequest.Service, fsmHandler *FsmHandler) *CallbackHandler {
	return &CallbackHandler{
		ctx:            ctx,
		requestService: requestService,
		fsmHandler:     fsmHandler,
	}
}

func (h *CallbackHandler) HandleCallback(bot *tgbotapi.BotAPI, query *tgbotapi.CallbackQuery) {
	data := query.Data
	message := query.Message
	messageID := query.Message.MessageID
	chatID := query.Message.Chat.ID

	switch data {
	case "approve":
		err := h.requestService.Handle(h.ctx, messageID, true, "")
		if err != nil {
			return
		}

		EditMessage(bot, chatID, messageID, message.Text, "Комментарий принят!")
	case "approve-with-comment":
		bot.Send(tgbotapi.NewMessage(chatID, "Напишите сообщение, чтобы принять комментарий с ответом"))
		h.fsmHandler.StartReplyFSM(chatID, messageID, message.Text)
	case "decline":
		err := h.requestService.Handle(h.ctx, messageID, false, "")
		if err != nil {
			return
		}

		EditMessage(bot, chatID, messageID, message.Text, "Комментарий отвергнут")
	}
}

func EditMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int, originText, action string) {
	lines := strings.Split(originText, "\n")
	filtered := make([]string, 0, len(lines))
	for _, line := range lines {
		if line != "Новая заявка на комментарий!" {
			filtered = append(filtered, line)
		}
	}
	newText := strings.Join(filtered, "\n") + "\n\n" + action

	edit := tgbotapi.NewEditMessageText(chatID, messageID, newText)
	bot.Send(edit)
}
