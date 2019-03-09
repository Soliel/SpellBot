package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver connection
	"log"
)

var db *sql.DB

//InitDB should be used with connection string in the format of user:password@tcp(127.0.0.1:3306)/hello
func InitDB(loginString string) {
	db, err := sql.Open("mysql", loginString)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to connect to database, ", err)
	}
}
