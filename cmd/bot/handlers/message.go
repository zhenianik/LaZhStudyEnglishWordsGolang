package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/buttons"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/service"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/wordsUtil"
)

var Word string

func MessageHandler(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
	sndMsg.ReplyToMessageID = message.MessageID
	sndMsg.ParseMode = "Markdown"
	sndMsg.DisableNotification = true
	sndMsg.ReplyMarkup = buttons.MainCommands()

	if wordsUtil.GetLang(message.Text) == "en" {
		inBase := false
		Word, inBase = service.CheckWordInBase(message.Text)
		sndMsg.Text = Word
		if inBase {
			sndMsg.ReplyMarkup = buttons.MainCommands()
		} else {
			sndMsg.ReplyMarkup = buttons.ControlCommands()
		}
	}
	if wordsUtil.GetLang(message.Text) == "ru" {
		translate := service.CheckTranslateInBase(message.Text)
		sndMsg.Text = translate
	}

	go bot.Send(sndMsg)
}
