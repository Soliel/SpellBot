package data

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInitDB(t *testing.T) {
	type args struct {
		loginString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Connection Test",
			args: args{
				loginString: "remote:testing@tcp(192.168.56.4:3306)/SpellBot",
			},
			wantErr: false,
		},
		{
			name: "Invalid Credential Test",
			args: args{
				loginString: "blah:testing@tcp(192.168.56.4:3306)/SpellBot",
			},
			wantErr: true,
		},
		{
			name: "Unable to connect test",
			args: args{
				loginString: "remote:testing@tcp(192.168.56.5:3306)/SpellBot",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitDB(tt.args.loginString); (err != nil) != tt.wantErr {
				t.Errorf("InitDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
