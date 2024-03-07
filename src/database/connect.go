package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //mysql driver
)

var Db *sql.DB

func ConnectDatabase() {
	// connect to database
	var err error

	log.Println("Connecting to database...")

	hostname, ok := os.LookupEnv("MYSQL_HOSTNAME")
	if !ok {
		hostname = "dastabase"
	}
	port, ok := os.LookupEnv("MYSQL_PORT")
	if !ok {
		port = "3306"
	}
	username, ok := os.LookupEnv("MYSQL_USERNAME")
	if !ok {
		username = "api"
	}
	password, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		password = "password"
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/holidays", username, password, hostname, port)

	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln("can not open database.", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalln("can not ping.", err)
	}
}
