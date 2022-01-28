package golog

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logPath    = "runtime/logs/"
	logName    = "log"
	logExt     = "log"
	timeFormat = "20060102"
)

func getFilePath() string {
	return fmt.Sprintf("%s%s%s.%s", logPath, logName, time.Now().Format(timeFormat), logExt)
}

func openFile(filePath string) (*os.File, error) {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle, err
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+logPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
