package service

import (
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/database/SqlQueries"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/database/bdService"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/wordsUtil"
	"strings"
)

var currentWord = ""
var currentTranslate = ""

func CheckWordInBase(text string) (string, bool) {

	answer := bdService.GetRequest(SqlQueries.CheckWord(text))
	mystr := ""
	if len(answer) != 0 {
		for _, s := range answer {
			mystr = wordsUtil.GetTranslatedString(strings.Split(s, ";"))
		}
		return mystr, true
	} else {
		translate := wordsUtil.Translate(text)
		mystr = "Cлова \"" + text + "\" с переводом \"" + translate + "\" нет в словаре! Добавить?"
		currentWord = text
		currentTranslate = translate
		return mystr, false
	}
}

func CheckTranslateInBase(text string) string {
	mystr := ""
	answer := bdService.GetRequest(SqlQueries.CheckTranslate(text))
	if len(answer) != 0 {
		for _, s := range answer {
			//translates := strings.Split(s, ";")
			mystr = mystr + wordsUtil.GetTranslatedString(strings.Split(s, ";")) + "\n"
		}
	}
	return mystr
}

func AddNewWordResult(username string, add bool) string {
	if currentWord == "" && currentTranslate == "" {
		return "Сначала введите слово"
	}
	if add {
		answer := bdService.GetRequestInsert(SqlQueries.AddNewWord(currentWord, currentTranslate, username))
		if answer && currentWord != "" && currentTranslate != "" {
			return "Слово успешно добавлено"
		} else {
			return "Что-то пошло не так, слово не добавлено."
		}
	} else {
		currentWord = ""
		currentTranslate = ""
		return "Хорошо, не будем добавлять."
	}
}
