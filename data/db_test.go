package data

import (
	"io/ioutil"
	"testing"

	"github.com/soliel/SpellBot/config"
)

func TestInitDB(t *testing.T) {
	configJSON, err := ioutil.ReadFile("../Configuration Files/SpellBotTest.json")
	if err != nil {
		t.Errorf("Test was unable to load configuration file.")
	}

	testingConfig, err := config.LoadConfig(configJSON)
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
				loginString: CreateDatabaseString(*testingConfig),
			},
			wantErr: false,
		},
		{
			name: "Invalid Credential Test",
			args: args{
				loginString: "postgres://blah:testing@192.168.56.4:3306/SpellBot",
			},
			wantErr: true,
		},
		{
			name: "Unable to connect test",
			args: args{
				loginString: "postgres://remote:testing@0.0.0.0:3306/SpellBot",
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
					DatabaseIP:   "127.0.0.1",
					DatabasePass: "testing",
					DatabaseUser: "test",
					DatabasePort: "3306",
					DatabaseName: "testdb",
					SSLMode:      "disable",
				},
			},
			want: "postgres://test:testing@127.0.0.1:3306/testdb?sslmode=disable",
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
