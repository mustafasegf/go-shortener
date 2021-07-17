package util

import (
	"io"
	"log"
	"os"
)

func SetLogger() (err error) {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", os.ModeDir)
	}
	logFile, err := os.OpenFile("log/log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	return
}
