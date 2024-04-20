package main

import (
	"database/sql"
	"guruchan-back/logging"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//logsetting
	logger := log.New(logging.LogSetting(), "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	err := logger.Output(2, "open logfile success")
	if err != nil {
		log.Fatal("file=logFile errMsg = %s", err.Error())
	}
	// init get.env file
	err = godotenv.Load("docker-compose")
	if err != nil {
		log.Fatal("file=docker-compose.env errMsg = %s", err.Error())
	}
	//local db connect
	dsn := os.Getenv("LOCALDSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("db connect failed ,errMsg = %s", err.Error())
	}
	defer db.Close()

	//TOdo router
	// r := router.New()

}
