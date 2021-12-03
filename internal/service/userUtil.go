package service

import (
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/database/SqlQueries"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/database/bdService"
)

func CheckUserPermission(u string) bool {
	queryString := SqlQueries.GetUser()
	param := []interface{}{u}
	return len(bdService.GetRequest(queryString, param)) > 0
}
