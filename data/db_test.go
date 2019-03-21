package data

import (
	"testing"
	"github.com/soliel/SpellBot/config"
	_ "github.com/go-sql-driver/mysql"
)

func TestInitDB(t *testing.T) {

	testingConfig, err := config.LoadConfig("../SpellBotTest.json") 
	if err != nil {
		t.Errorf("Test not completed; config didn't load due to: %v", err)
	}

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
				loginString: CreateDatabaseString(*testingConfig) + "/SpellBot",
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
				loginString: "remote:testing@tcp(0.0.0.0:3306)/SpellBot",
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

func TestCreateDBString(t *testing.T) {
	type args struct {
		config config.Config
	}
	tests := []struct {
		name string
		args args
		want string	
	}{
		{
			name: "SpellBot Test",
			args: args{
				config: config.Config{
					DatabaseIP: "127.0.0.1",
					DatabasePass: "testing",
					DatabaseUser: "test",
					DatabasePort: "3306",
				},
			},
			want: "test:testing@tcp(127.0.0.1:3306)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if dbstring := CreateDatabaseString(tt.args.config); dbstring != tt.want {
				t.Errorf("CreateDbString() produced: %v, wanted: %v", dbstring, tt.want)
			}
		})
	}
}
