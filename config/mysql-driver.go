package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:aminrais19@tcp(localhost:3306)/account_service_app_project")

	if err != nil {
		log.Fatal("error sql open", err.Error())

	}
	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db", errPing.Error())
	} else {
		fmt.Println("succes connect to DB")
	}
	return db
}
