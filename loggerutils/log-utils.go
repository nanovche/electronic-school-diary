package logger

import (
	"log"
	"os"
)

type logType int

const (
	Error logType = 1
	Event
)

func (d logType) String() string {
	return [...]string{"Error: ", "Event: "}[d]
}

func InitLogger(fileName string, typeOfLogger string) *log.Logger {
	var file *os.File
	var err error
	if file, err = os.Create("fileName"); err != nil {

	}
	return log.New(file, typeOfLogger + ": ", 1)
}
