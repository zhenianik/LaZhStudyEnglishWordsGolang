package SqlQueries

import (
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/wordsUtil"
	"strings"
)

func GetRandomVideo() string {
	return "SELECT tb1.word, tb1.translate1, tb1.translate2, tb1.translate3, tb1.translate4, tb1.context from words AS tb1 " +
		"INNER JOIN ( SELECT DISTINCT word, translate1, translate2, translate3, translate4, context from words " +
		"where context LIKE '%http%'  order by RAND() LIMIT 1) AS tb2 " +
		"ON tb1.context = tb2.context"
}

func GetRandomWords() string {
	return "SELECT word, translate1, translate2, translate3, translate4, context from words order by RAND() LIMIT 30"
}

func GetRandomWordsFromLast() string {
	return "SELECT word, translate1, translate2, translate3, translate4, context from " +
		"(SELECT word, translate1, translate2, translate3, translate4, context from words ORDER BY `id_word` DESC LIMIT 100) " +
		"as table1 order by RAND() LIMIT 30"
}

func GetLastWords() string {
	return "SELECT word, translate1, translate2, translate3, translate4, context from words ORDER BY `id_word` DESC LIMIT 30"
}

func CheckWord() string {
	return "SELECT word, translate1, translate2, translate3, translate4, context from words where TRIM(word) = ?"
}

func CheckTranslate() string {
	return "SELECT word, translate1, translate2, translate3, translate4, context from words " +
		"where TRIM(translate1) = ? OR TRIM(translate2) = ? OR TRIM(translate3) = ? OR TRIM(translate4) = ? "
}

func GetPhrasalVerbs() string {
	phrasalVerbsParticles := wordsUtil.GetParticles()
	if len(phrasalVerbsParticles) == 0 {
		return ""
	}
	query := "SELECT word, translate1, translate2, translate3, translate4, context from words where TRIM(word) like '% %'"

	query = query + " AND ("
	counter := 0
	for _, particle := range phrasalVerbsParticles {
		if counter != 0 {
			query = query + " OR"
		}
		query = query + " TRIM(word) like '% " + strings.ToLower(particle) + "%'"
		counter++
	}
	query = query + " ) order by RAND()"

	return query
}

func AddNewWord() string {
	return "INSERT INTO `words` (`id_word`, `word`, `translate1`, `translate2`, `translate3`, `translate4`, `context`) " +
		"VALUES (NULL, ?, ?, '', '', '', ?)"
}

func GetUser() string {
	return "SELECT name from users where TRIM(name) = ?"
}
