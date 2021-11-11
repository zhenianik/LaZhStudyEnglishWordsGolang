package buttons

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func MainCommands() *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Test", "/test"),
			tgbotapi.NewInlineKeyboardButtonData("Last", "/last"),
			tgbotapi.NewInlineKeyboardButtonData("Test last", "/testLast"),
			tgbotapi.NewInlineKeyboardButtonData("Video", "/video"),
			tgbotapi.NewInlineKeyboardButtonData("Ph. verbs", "/phV"),
		),
	)
	return &keyboard
}

func ControlCommands() *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yes", "/yes"),
			tgbotapi.NewInlineKeyboardButtonData("No", "/no"),
		),
	)
	return &keyboard
}
