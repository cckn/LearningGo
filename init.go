package main

import (
	"fmt"
	"logger"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var logFile *os.File

func init() {

	logFilePath := fmt.Sprintf("logs/%s", time.Now().Format("20060102"))
	logFileName := time.Now().Format("15-04")
	logFile = logger.InitFile(logFilePath, logFileName)

	log.SetFormatter(&log.JSONFormatter{
		// DisableTimestamp: true,
	})

}

func fileClose() {
	log.Printf("---------------------END---------------------")
	_ = logFile.Close()
}
