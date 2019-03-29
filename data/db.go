package data

import (
	"database/sql"

	_ "github.com/lib/pq" //Driver connection
	"github.com/soliel/SpellBot/config"
)

var db *sql.DB

//InitDB should be used with connection string in the format of user:password@tcp(127.0.0.1:3306)/hello
func InitDB(loginString string) error {
	database, err := sql.Open("postgres", loginString)
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

//CreateDatabaseString converts the config file into an appropriate string.
func CreateDatabaseString(conf config.Config) string {
	return "postgres://" +
		conf.DatabaseUser + ":" +
		conf.DatabasePass + "@" +
		conf.DatabaseIP + ":" +
		conf.DatabasePort + "/" +
		conf.DatabaseName + "?sslmode=" +
		conf.SSLMode
}
