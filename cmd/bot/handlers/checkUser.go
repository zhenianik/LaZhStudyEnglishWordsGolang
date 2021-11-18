package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/buttons"
)

func UserHandler(bot *tgbotapi.BotAPI, message *tgbotapi.Message, callback *tgbotapi.CallbackQuery, s string) {
	if message == nil {
		message = callback.Message
	}
	sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
	sndMsg.ParseMode = "Markdown"
	sndMsg.DisableNotification = true
	sndMsg.ReplyMarkup = buttons.MainCommands()
	sndMsg.Text = s

	go bot.Send(sndMsg)
}
