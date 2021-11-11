package bdService

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zhenianik/LaZhStudyEnglishWordsGolang/internal/database/SqlQueries"
	"testing"
)

func TestGetRequest(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "GetRandomWords",
			args: args{SqlQueries.GetRandomWords()},
		},
		{
			name: "GetRandomVideo",
			args: args{SqlQueries.GetRandomVideo()},
		},
		{
			name: "GetRandomWordsFromLast",
			args: args{SqlQueries.GetRandomWordsFromLast()},
		},
		{
			name: "GetLastWords",
			args: args{SqlQueries.GetLastWords()},
		},
		{
			name: "GetPhrasalVerbs",
			args: args{SqlQueries.GetPhrasalVerbs()},
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"word", "translate1", "translate2", "translate3", "translate4", "context"}).
		AddRow("take up", "начать", "занимать", "принимать", "браться за", "").
		AddRow("put over", "отложить", "переправить", "успешно осуществить", "", "")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery(tt.args.s).WillReturnRows(rows)
		})
	}

}
