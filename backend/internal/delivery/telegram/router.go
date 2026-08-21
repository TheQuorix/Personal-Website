package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Router struct {
	msgHandler      *MessageHandler
	callbackHandler *CallbackHandler
}

func NewRouter(msgHandler *MessageHandler, callbackHandler *CallbackHandler) *Router {
	return &Router{msgHandler: msgHandler, callbackHandler: callbackHandler}
}

func (r *Router) Route(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		r.msgHandler.HandleMessage(bot, update.Message)
	} else if update.CallbackQuery != nil {
		r.callbackHandler.HandleCallback(bot, update.CallbackQuery)
	}
}
