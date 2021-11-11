package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/cmd/bot/handlers"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/config"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("в процесе выполнения возникла ошибка: %v", err)
	}
}

func run() error {

	// config
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("ошибка получения конфига: %w", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		return fmt.Errorf("не удалось подключиться к телеграм боту: %w", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return fmt.Errorf("не удалось получить канал с обновлениями: %w", err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			go handlers.CallbackHandler(bot, update.CallbackQuery)
		} else if update.Message != nil {
			go handlers.MessageHandler(bot, update.Message)
		}
	}

	return nil
}
