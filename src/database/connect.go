package database

import (
	"fmt"
	"github.com/kynmh69/go-ja-holidays/logging"
	"gorm.io/gorm"
	"os"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
)

var db *gorm.DB

const NAME = "holidays"

func ConnectDatabase() {
	// connect to database
	var err error
	logger := logging.GetLogger()
	hostname, port, dataSourceName := CreateConnectInfo()
	db, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		logger.Fatalln("can not open database.", err)
	}

	logger.Info("Connecting to database...", hostname, port)
}

func CreateConnectInfo() (string, string, string) {
	hostname, port, username, password, databaseName := getConnectionInfo()

	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo",
		hostname, port, username, password, databaseName,
	)
	return hostname, port, dataSourceName
}

func getConnectionInfo() (hostname, port, username, password, databaseName string) {
	hostname, ok := os.LookupEnv("PSQL_HOSTNAME")
	if !ok {
		hostname = "database"
	}
	port, ok = os.LookupEnv("PSQL_PORT")
	if !ok {
		port = "5432"
	}
	username, ok = os.LookupEnv("PSQL_USERNAME")
	if !ok {
		username = "app"
	}
	password, ok = os.LookupEnv("PSQL_PASSWORD")
	if !ok {
		password = "password"
	}
	databaseName, ok = os.LookupEnv("DATABASE")

	if !ok {
		databaseName = NAME
	}
	return
}

func GetDbConnection() *gorm.DB {
	return db
}
