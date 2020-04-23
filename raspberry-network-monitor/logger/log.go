package logger

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

const (
	logFolder = "~/.raspberry-cli"
	logFile   = "~/.raspberry-cli/network-monitor.log"
)

func init() {
	if logDir, err := homedir.Expand(logFolder); err == nil {
		os.MkdirAll(logDir, os.ModePerm)
		mw := io.MultiWriter(os.Stdout, getLogFile())
		log.SetOutput(mw)
	} else {
		panic(err)
	}
}

func getLogFile() *os.File {
	//logFile := config.GetLogFile()
	f, openFileErr := os.OpenFile(getLogFilePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if openFileErr != nil {
		f, createFileErr := os.Create(logFile)
		if createFileErr != nil {
			panic(createFileErr.Error())
		}
		return f
	}

	return f
}

func getLogFilePath() string {
	if logFilePath, err := homedir.Expand(logFile); err != nil {
		panic(err)
	} else {
		return logFilePath
	}
}

//func init() {
//	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		log.Fatalf("error opening file: %v", err)
//	}
//	defer f.Close()
//
//	log.SetOutput(f)
//}

/*
Debug prints an object to log as a JSON
*/
func Debug(obj interface{}) {
	jsonBytes, _ := json.MarshalIndent(&obj, "", "  ")
	log.Println(string(jsonBytes))
}

/*
Println prints message to log
*/
func Println(v ...interface{}) {
	log.Println(v)
}

/*
Printf prints message to log
*/
func Printf(msg string, v ...interface{}) {
	log.Printf(msg, v)
}
