package util

import (
	"io"
	"log"
	"os"
)

func SetLogger() (err error) {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	return
}
