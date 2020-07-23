package main

import (
	"fmt"
	"os"
	"time"

	logger "logger"

	log "github.com/sirupsen/logrus"
)

var logFile *os.File

func init() {

	logFilePath := fmt.Sprintf("logs/%s", time.Now().Format("20060102"))
	logFileName := time.Now().Format("15-04-05")
	logFile = logger.InitFile(logFilePath, logFileName)

	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		DisableQuote:           true,
	})

}
func main() {

	defer func() {
		_ = logFile.Close()
	}()

	log.Fatal("asdasd")

}
