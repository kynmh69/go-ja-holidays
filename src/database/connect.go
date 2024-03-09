package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	// import the dialect
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

var goquDb *goqu.Database

const DATABASE_NAME = "holidays"

func ConnectDatabase() {
	// connect to database
	var err error

	hostname, port, username, password := getConnectionInfo()

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", username, password, hostname, port, DATABASE_NAME)

	db, err := sql.Open("postgres", dataSourceName)
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

	defer db.Close()

	goquDb = goqu.New("postgres", db)
	goquDb.Logger(initLogger())
}

func getConnectionInfo() (string, string, string, string) {
	hostname, ok := os.LookupEnv("MYSQL_HOSTNAME")
	if !ok {
		hostname = "database"
	}
	port, ok := os.LookupEnv("MYSQL_PORT")
	if !ok {
		port = "5432"
	}
	username, ok := os.LookupEnv("MYSQL_USERNAME")
	if !ok {
		username = "app"
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
