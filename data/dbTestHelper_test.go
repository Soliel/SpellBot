package data

import (
	"errors"
	"io/ioutil"

	"github.com/soliel/SpellBot/config"
)

func OpenTestDB() error {
	testOpenConn := InitDB("teststring") //This string would normally return an error. if it returns nil we're in buisness.
	if testOpenConn != nil {
		return nil
	}

	loadBytes, err := ioutil.ReadFile("../Configuration Files/SpellBotTest.json")
	if err != nil {
		return errors.New("Unable to read file: " + err.Error())
	}

	conf, err := config.LoadConfig(loadBytes)
	if err != nil {
		return errors.New("Unable to parse JSON: " + err.Error())
	}

	err = InitDB(CreateDatabaseString(*conf))
	if err != nil {
		return errors.New("Error initalizing database connection: " + err.Error())
	}

	return nil
}
