package config

import (
	"encoding/json"
)

//Config hold general configuration
type Config struct {
	BotToken     string `json:"bot_token"`
	BotPrefix    string `json:"bot_prefix"`
	DatabaseIP   string `json:"database_ip"`
	DatabaseUser string `json:"database_user"`
	DatabasePass string `json:"database_password"`
	DatabasePort string `json:"database_port"`
	DatabaseName string `json:"database_name"`
	SSLMode      string `json:"ssl_mode"`
}

//LoadConfig loads the main config file for the application
func LoadConfig(fileBytes []byte) (*Config, error) {
	var confData Config
	err := json.Unmarshal(fileBytes, &confData)
	if err != nil {
		return nil, err
	}

	return &confData, nil
}
