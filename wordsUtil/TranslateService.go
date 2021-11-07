package wordsUtil

import (
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/database"
	"strings"
)

var currentWord = ""
var currentTranslate = ""

func CheckWordInBase(text string) string {

	answer := database.GetRequest(CheckWord(text))
	mystr := ""
	if len(answer) != 0 {
		for _, s := range answer {
			mystr = getTranslatedString(strings.Split(s, ";"))
		}
		return mystr
	} else {
		translate := Translate(text)
		mystr = "Cлова \"" + text + "\" с переводом \"" + translate + "\" нет в словаре! Добавить?"
		currentWord = text
		currentTranslate = translate
		return mystr
	}
}

func CheckTranslateInBase(text string) string {
	mystr := ""
	answer := database.GetRequest(CheckTranslate(text))
	if len(answer) != 0 {
		for _, s := range answer {
			//translates := strings.Split(s, ";")
			mystr = mystr + getTranslatedString(strings.Split(s, ";")) + "\n"
		}
	}
	return mystr
}

func AddNewWordResult(username string, add bool) string {
	if currentWord == "" && currentTranslate == "" {
		return "Сначала введите слово"
	}
	if add {
		answer := database.GetRequestInsert(AddNewWord(currentWord, currentTranslate, username))
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
