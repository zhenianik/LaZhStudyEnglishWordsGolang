package bdService

import (
	"database/sql"
	"fmt"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/config"
	"log"
)

var db *sql.DB

func CreateConnection() error {
	var err error

	cfg, _ := config.GetConfig()
	if err != nil {
		return fmt.Errorf("ошибка получения конфига: %w", err)
	}

	db, err = sql.Open("mysql", cfg.Dbuser+":"+cfg.Dbpass+"@tcp("+cfg.Dbhost+":"+cfg.Dbport+")/"+cfg.Dbname+"?charset=utf8")

	if err != nil {
		return fmt.Errorf("не удалось с базой данных: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("отсутствует пинг: %w", err)
	}
	return err
}

func checkConnect() error {
	if db == nil {
		err := CreateConnection()
		if err != nil {
			return fmt.Errorf("не получилось создать соединение: %w", err)
		}
	}
	err := db.Ping()
	if err != nil {
		err = db.Close()
		if err != nil {
			return fmt.Errorf("ошибка закрытия соединения с базой: %w", err)
		}
		err = CreateConnection()
		if err != nil {
			return fmt.Errorf("не получилось создать соединение: %w", err)
		}
	}
	return nil
}

func GetRequest(sqlText string) []string {

	err := checkConnect()
	if err != nil {
		log.Printf("ошибка соединения с бд: %v", err)
	}

	output := []string{}

	results, err := db.Query(sqlText)
	defer results.Close()

	if err != nil {
		log.Printf("ошибка выполнения запроса выборки данных: %v", err)
	}

	for results.Next() {
		var word, translate1, translate2, translate3, translate4, context string
		err = results.Scan(&word, &translate1, &translate2, &translate3, &translate4, &context)
		if err != nil {
			log.Printf("ошибка получения результатов выборки: %v", err)
		}
		wordAndTranslate := word + ";" + translate1 +
			";" + translate2 + ";" + translate3 +
			";" + translate4 + ";" + context

		output = append(output, wordAndTranslate)
	}
	return output
}

func GetRequestInsert(sqlText string) bool {

	err := checkConnect()
	if err != nil {
		log.Printf("ошибка соединения с бд: %v", err)
	}

	insert, err := db.Query(sqlText)
	if err != nil {
		log.Printf("ошибка выполнения запроса вставки данных: %v", err)
	}
	defer insert.Close()

	return true
}
