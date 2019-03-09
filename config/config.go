package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BotToken     string `json:"bot_token"`
	BotPrefix    string `json:"bot_prefix"`
	DatabaseIP   string `json:"database_ip"`
	DatabaseUser string `json:"database_user"`
	DatabasePass string `json:"database_password"`
	DatabasePort string `json:"database_port"`
}

func LoadConfig(filename string) (*Config, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var confData Config
	err = json.Unmarshal(body, &confData)
	if err != nil {
		return nil, err
	}

	return &confData, nil
}
