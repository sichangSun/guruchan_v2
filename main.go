package main

import (
	"database/sql"
	"fmt"
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
		log.Fatal("file=logFile err = %s", err.Error())
	}
	// init get.env file
	err = godotenv.Load("docker-compose")
	if err != nil {
		log.Fatal("file=docker-compose.env err = %s", err.Error())
	}

	os.Getenv("docker-compose")
	db, err := sql.Open("mysql", "root:root1234@tcp(localhost:3306)/pro-Test")
	if err != nil {
		fmt.Println("error occurred. err=", err)
	}
	defer db.Close()
}
