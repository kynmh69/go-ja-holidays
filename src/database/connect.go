package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	// import the dialect
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

var goquDb *goqu.Database

func ConnectDatabase() {
	// connect to database
	var err error

	hostname, port, username, password := getConnectionInfo()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/holidays", username, password, hostname, port)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln("can not open database.", err)
	}

	log.Println("Connecting to database...", hostname, port)

	err = db.Ping()
	if err == nil {
		log.Println("Connected to database.")
	} else {
		log.Fatalln("can not ping.", err)
	}

	goquDb = goqu.New("mysql", db)
	goquDb.Logger(initLogger())
}

func getConnectionInfo() (string, string, string, string) {
	hostname, ok := os.LookupEnv("MYSQL_HOSTNAME")
	if !ok {
		hostname = "database"
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
	return hostname, port, username, password
}

func GetDbConnection() *goqu.Database {
	return goquDb
}

func initLogger() *log.Logger {
	return log.New(os.Stdout, "[SQL] ", log.LstdFlags|log.Lshortfile)
}
