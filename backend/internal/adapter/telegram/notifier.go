package telegram

import (
	"context"
	"fmt"

	"github.com/TheQuorix/Personal-Website/internal/domain/commentrequest"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Структура уведомлений в телеграм
type Notifier struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

// Создание сервиса уведомлений
func NewNotifier(bot *tgbotapi.BotAPI, chatID int64) *Notifier {
	return &Notifier{bot: bot, chatID: chatID}
}

// Проверка совпадает ли сервис уведомлений с интерфейсом
var _ commentrequest.Notifier = (*Notifier)(nil)

// Отправить комментарий без публикации
func (n *Notifier) NotifyWithoutPublish(ctx context.Context, r commentrequest.CommentRequest) error {
	msg := tgbotapi.NewMessage(n.chatID,
		fmt.Sprintf("Новая комментарий!\n\n%s:\n- %s", r.Author, r.Message))

	_, err := n.bot.Send(msg)
	return err
}

// Отправить комментарий с заявкой на публикацию
func (n *Notifier) NotifyWithPublish(ctx context.Context, r commentrequest.CommentRequest) (int, error) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Принять", "approve"),
			tgbotapi.NewInlineKeyboardButtonData("Отказать", "decline"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ответить и принять", "approve-with-comment"),
		),
	)

	msg := tgbotapi.NewMessage(n.chatID,
		fmt.Sprintf("Новая заявка на комментарий!\n\n%s:\n- %s", r.Author, r.Message))
	msg.ReplyMarkup = keyboard

	message, err := n.bot.Send(msg)

	return message.MessageID, err
}
