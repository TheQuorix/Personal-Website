package telegram

import (
	"context"
	"sync"

	"github.com/TheQuorix/Personal-Website/internal/domain/commentrequest"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type UserState struct {
	MessageID  int
	ChatID     int64
	OriginText string
	Waiting    bool
}

var states = make(map[int64]*UserState)
var statesMu sync.Mutex

type FsmHandler struct {
	ctx            context.Context
	requestService *commentrequest.Service
}

func NewFsmHandler(ctx context.Context, requestService *commentrequest.Service) *FsmHandler {
	return &FsmHandler{
		ctx:            ctx,
		requestService: requestService,
	}
}

func (h *FsmHandler) GetActiveState(chatID int64) (*UserState, bool) {
	statesMu.Lock()
	defer statesMu.Unlock()

	state, exists := states[chatID]
	return state, exists
}

func (h *FsmHandler) HandleFSM(bot *tgbotapi.BotAPI, message *tgbotapi.Message, state *UserState) {
	chatID := message.Chat.ID

	statesMu.Lock()
	delete(states, chatID)
	statesMu.Unlock()

	if message.Text == "" {
		bot.Send(tgbotapi.NewMessage(chatID, "Ответ должен быть текстом. Попробуйте начать заново."))
		return
	}

	err := h.requestService.Handle(h.ctx, state.MessageID, true, message.Text)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "Не удалось сохранить ответ: "+err.Error()))
		return
	}

	EditMessage(bot, chatID, state.MessageID, state.OriginText, "Комментарий принят!\n\nОтвет: "+message.Text)

	bot.Send(tgbotapi.NewMessage(chatID, "Ответ сохранён."))
}

func (h *FsmHandler) StartReplyFSM(chatID int64, messageID int, originText string) {
	statesMu.Lock()
	defer statesMu.Unlock()

	states[chatID] = &UserState{
		MessageID:  messageID,
		OriginText: originText,
		Waiting:    true,
	}
}
