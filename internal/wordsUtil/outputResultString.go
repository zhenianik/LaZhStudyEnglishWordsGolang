package wordsUtil

import (
	"strings"
)

func GetResultStr(answer []string, showContext bool) string {
	resultStr := ""
	if len(answer) == 0 {
		return ""
	}
	for _, s := range answer {

		resultStr += strings.ToLower(GetTranslatedString(strings.Split(s, ";")) + "\n")
	}
	if showContext {
		arr := strings.Split(answer[len(answer)-1], ";")
		resultStr = resultStr + "\n" + arr[len(arr)-1]
	}
	return resultStr
}

func GetTranslatedString(arr []string) string {
	mystr := ""
	if len(arr) == 0 {
		return mystr
	}
	for i := 0; i <= 4; i++ {
		mystr += arr[i] + ", "
		if arr[i] == "" {
			break
		}
	}

	mystr = strings.TrimRight(mystr, ", ")
	mystr = strings.Replace(mystr, ", ", " - ", 1)

	return mystr
}

func GetLang(s string) string {
	s = strings.TrimSpace(s)
	r := []rune(s)[0]
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return "en"
	} else if (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я') {
		return "ru"
	} else {
		//TODO обработать ошибку
		return "строка начинается с символа неизвестного языка"
	}
}
