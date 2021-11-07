package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/config"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/database"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/wordsUtil"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var word string

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//ctx := context.Background()

	// config
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	b, err := tb.NewBot(tb.Settings{
		URL:    "",
		Token:  cfg.TelegramBotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return err
	}

	process(b)

	return nil
}

func process(b *tb.Bot) {
	b.Handle("/test", func(m *tb.Message) {
		b.Send(m.Sender, wordsUtil.GetResultStr(database.GetRequest(wordsUtil.GetRandomWords()), false))
	})
	b.Handle("/last", func(m *tb.Message) {
		b.Send(m.Sender, wordsUtil.GetResultStr(database.GetRequest(wordsUtil.GetLastWords()), false))
	})
	b.Handle("/testLast", func(m *tb.Message) {
		b.Send(m.Sender, wordsUtil.GetResultStr(database.GetRequest(wordsUtil.GetRandomWordsFromLast()), false))
	})
	b.Handle("/video", func(m *tb.Message) {
		b.Send(m.Sender, wordsUtil.GetResultStr(database.GetRequest(wordsUtil.GetRandomVideo()), true))
	})
	b.Handle("/phV", func(m *tb.Message) {
		b.Send(m.Sender, wordsUtil.GetResultStr(database.GetRequest(wordsUtil.GetPhrasalVerbs()), false))
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		// если ввели слово (не команда), тогда проверим его наличие в бд и выведем соотв.результат
		if wordsUtil.GetLang(m.Text) == "en" {
			word = wordsUtil.CheckWordInBase(m.Text)
			b.Send(m.Sender, word)
		}
		if wordsUtil.GetLang(m.Text) == "ru" {
			translate := wordsUtil.CheckTranslateInBase(m.Text)
			b.Send(m.Sender, translate)
		}
	})

	b.Handle("/yes", func(m *tb.Message) {
		b.Send(m.Sender, wordsUtil.AddNewWordResult(m.Sender.Username, true))
	})
	b.Handle("/no", func(m *tb.Message) {
		b.Send(m.Sender, wordsUtil.AddNewWordResult(m.Sender.Username, false))
	})

	b.Start()
}
