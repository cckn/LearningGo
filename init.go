package main

import (
	"fmt"
	"os"
	"time"

	logger "github.com/cckn/go.logger"

	log "github.com/sirupsen/logrus"
)

var logFile *os.File

func init() {
	logFilePath := fmt.Sprintf("logs/%s", time.Now().Format("20060102"))
	logFileName := time.Now().Format("15")
	logFile = logger.InitFile(logFilePath, logFileName)

	_, _ = logFile.WriteString("\n")

	log.SetFormatter(&log.JSONFormatter{
		// DisableTimestamp: true,
	})
	log.Println("- - - - - - - - - - - -RUN- - - - - - - - - - - -")

}

func fileClose() {
	log.Printf("- - - - - - - - - - - -END- - - - - - - - - - - -")
	_ = logFile.Close()
}
