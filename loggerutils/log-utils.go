package loggerutils

import (
	"log"
	"os"
)

type logType int

const (
	Error logType = 1
	Event
)

const(
	ErrorLoggerFileName = "errorLogger.txt"
	EventLoggerFileName = "eventLogger.txt"
)

func (d logType) String() string {
	return [...]string{"Error: ", "Event: "}[d]
}

func InitLogger(fileName string, typeOfLogger string) *log.Logger {
	var file *os.File
	var err error
	if file, err = os.Create("fileName"); err != nil {
		log.Println("Failed to initialize log files. This logger will print to standard error.")
		return log.New(os.Stderr, typeOfLogger, 1)
	}
	return log.New(file, typeOfLogger + ": ", 1)
}
