package wordsUtil

import "testing"

func TestGetLang(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english test",
			args: args{s: "hello"},
			want: "en",
		},
		{
			name: "russian test",
			args: args{s: "привет"},
			want: "ru",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLang(tt.args.s); got != tt.want {
				t.Errorf("GetLang() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetResultStr(t *testing.T) {
	type args struct {
		s           []string
		showContext bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{
				[]string{
					"hi;jo;and;relax;bro",
					"hi;mark;give;your;money",
				},
				true,
			},
			want: "",
		},
		{
			name: "russian test",
			args: args{
				[]string{
					"hi;jo;and;relax;bro",
				},
				false,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetResultStr(tt.args.s, true); got == tt.want {
				t.Errorf("GetLang() = %v, want %v", got, tt.want)
			}
		})
	}
}
