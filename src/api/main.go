package main

import (
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/util"
)

func init() {
	util.LoggerInitialize()
	database.ConnectDatabase()
}
func main() {
	
}