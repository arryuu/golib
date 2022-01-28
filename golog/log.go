package golog

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	//F *os.File

	defPrefix      = ""
	defCallerDepth = 2

	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

	logger *log.Logger
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
	NULL = -1
)

type (
	Level int

	LoggerSt struct {
		logger *log.Logger
		writer []io.Writer
	}
)

func New(writeToFile bool) (re *LoggerSt) {
	var writers []io.Writer
	writers = append(writers, os.Stdout)
	if writeToFile == true {
		file, err := openFile(getFilePath())
		if err == nil {
			writers = append(writers, file)
		}
	}

	return &LoggerSt{
		writer: writers,
		logger: log.New(io.MultiWriter(writers...), defPrefix, log.LstdFlags),
	}
}

func (l *LoggerSt) Debug(v ...interface{}) {
	l.output(DEBUG, fmt.Sprintln(v...))
}

func (l *LoggerSt) Info(v ...interface{}) {
	l.output(INFO, fmt.Sprintln(v...))
}

func (l *LoggerSt) Warn(v ...interface{}) {
	l.output(WARNING, fmt.Sprintln(v...))
}

func (l *LoggerSt) Error(v ...interface{}) {
	l.output(ERROR, fmt.Sprintln(v...))
}

//func (l *LoggerSt) Fatal(v ...interface{}) {
//	l.output(FATAL, fmt.Sprintln(v...))
//	//logger.Fatalln(v)
//}

func (l *LoggerSt) Printf(format string, a ...interface{}) {
	l.output(NULL, fmt.Sprintf(format, a...))
}

func (l *LoggerSt) output(level Level, txt string) {
	l.logger.SetPrefix(setPrefix(level))
	l.logger.Println(txt)
}

func setPrefix(level Level) string {
	_, file, line, ok := runtime.Caller(defCallerDepth)
	if ok {
		if level == NULL {
			logPrefix = ""
		} else {
			logPrefix = fmt.Sprintf("[%s] [%s:%d] ", levelFlags[level], filepath.Base(file), line)
		}
	} else {
		logPrefix = fmt.Sprintf("[%s] ", levelFlags[level])
	}

	return logPrefix
}
