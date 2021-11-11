package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/buttons"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/database/SqlQueries"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/database/bdService"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/service"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/wordsUtil"
)

func CallbackHandler(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	message := callback.Message
	sndMsg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")
	switch callback.Data {
	case "/test":
		queryString := SqlQueries.GetRandomWords()
		sndMsg.Text = wordsUtil.GetResultStr(bdService.GetRequest(queryString), false)
		sndMsg.ReplyMarkup = buttons.MainCommands()
	case "/last":
		queryString := SqlQueries.GetLastWords()
		sndMsg.Text = wordsUtil.GetResultStr(bdService.GetRequest(queryString), false)
		sndMsg.ReplyMarkup = buttons.MainCommands()
	case "/testLast":
		queryString := SqlQueries.GetRandomWordsFromLast()
		sndMsg.Text = wordsUtil.GetResultStr(bdService.GetRequest(queryString), false)
		sndMsg.ReplyMarkup = buttons.MainCommands()
	case "/video":
		queryString := SqlQueries.GetRandomVideo()
		sndMsg.Text = wordsUtil.GetResultStr(bdService.GetRequest(queryString), true)
		sndMsg.ReplyMarkup = buttons.MainCommands()
	case "/phV":
		queryString := SqlQueries.GetPhrasalVerbs()
		sndMsg.Text = wordsUtil.GetResultStr(bdService.GetRequest(queryString), false)
		sndMsg.ReplyMarkup = buttons.MainCommands()
	case "/yes":
		sndMsg.Text = service.AddNewWordResult(callback.From.UserName, true)
		sndMsg.ReplyMarkup = buttons.MainCommands()
	case "/no":
		sndMsg.Text = service.AddNewWordResult(callback.From.UserName, false)
		sndMsg.ReplyMarkup = buttons.MainCommands()
	default:
		sndMsg.Text = "Нет такой команды!"
		sndMsg.ReplyMarkup = buttons.MainCommands()
	}

	go bot.Send(sndMsg)

}
