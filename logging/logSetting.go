package logging

import (
	"io"
	"log"
	"os"
)

func LogSetting() io.Writer {
	logFile, err := os.OpenFile("./logging/guruchan.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("File open failure,file=logFile err = %s", err.Error())
	}
	multiLogFile := io.MultiWriter(os.Stdout, logFile)
	return multiLogFile
}
