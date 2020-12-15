package main

import (
	"log"
	"os"
)

var (
	//WarnLog  *log.Logger
	//InfoLog for logging information
	InfoLog *log.Logger
	//ErrorLog for error information
	ErrorLog *log.Logger
	//DebugLog *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLog = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	//WarnLog = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	//DebugLog = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
