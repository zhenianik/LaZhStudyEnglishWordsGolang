package database

import (
	"database/sql"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/config"
)

var db *sql.DB

func CreateConnection() *sql.DB {
	var err error
	cfg, _ := config.GetConfig()
	db, err = sql.Open("mysql", cfg.Dbuser+":"+cfg.Dbpass+"@tcp("+cfg.Dbhost+":"+cfg.Dbport+")/"+cfg.Dbname+"?charset=utf8")

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func checkConnect() {
	if db == nil {
		CreateConnection()
	}
	err := db.Ping()
	if err != nil {
		err = db.Close()
		if err != nil {
			panic(err.Error()) // TODO proper error handling instead of panic in your app
		}
		CreateConnection()
	}
}

func GetRequest(sqlText string) []string {

	checkConnect()

	output := []string{}

	results, err := db.Query(sqlText)
	defer results.Close()

	if err != nil {
		panic(err.Error()) // TODO proper error handling instead of panic in your app
	}

	for results.Next() {
		var word, translate1, translate2, translate3, translate4, context string
		err = results.Scan(&word, &translate1, &translate2, &translate3, &translate4, &context)
		if err != nil {
			panic(err.Error()) // TODO proper error handling instead of panic in your app
		}
		wordAndTranslate := word + ";" + translate1 +
			";" + translate2 + ";" + translate3 +
			";" + translate4 + ";" + context

		output = append(output, wordAndTranslate)
	}
	return output
}

func GetRequestInsert(sqlText string) bool {

	checkConnect()

	insert, err := db.Query(sqlText)
	if err != nil {
		// TODO добавить лог ошибки?
		return false
	}
	defer insert.Close()

	return true
}
