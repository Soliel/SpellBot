package data

import (
	"github.com/soliel/SpellBot/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver connection
)

var db *sql.DB

//InitDB should be used with connection string in the format of user:password@tcp(127.0.0.1:3306)/hello
func InitDB(loginString string) error {
	database, err := sql.Open("mysql", loginString)
	if err != nil {
		return err
	}

	err = database.Ping()
	if err != nil {
		return err
	}

	db = database
	return nil
}

func CreateDatabaseString(conf config.Config) string {
	return conf.DatabaseUser + ":" + conf.DatabasePass + "@tcp(" + conf.DatabaseIP + ":" + conf.DatabasePort + ")"
}
