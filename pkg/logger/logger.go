package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Logger struct {
	TimeFormat string
	FileName   string
}

func log(level LogLevel, message *string, filename string, lineNumber int) {

	fmt.Printf("[%v][%v][%v:%d][%v] %v \n",
		time.Now().Format(time.RFC3339),
		os.Getpid(),
		filename,
		lineNumber,
		level.String(),
		*message,
	)

}

func getCallerInfo() (string, int) {
	_, file, no, ok := runtime.Caller(2)

	if !ok {
		return "", -1
	}

	return filepath.Base(file), no
}

func (l *Logger) Debug(message string) {
	fileName, lineNumber := getCallerInfo()
	log(DEBUG, &message, fileName, lineNumber)
}

func (l *Logger) Info(message string) {
	fileName, lineNumber := getCallerInfo()
	log(INFO, &message, fileName, lineNumber)
}

func (l *Logger) Warn(message string) {
	fileName, lineNumber := getCallerInfo()
	log(WARN, &message, fileName, lineNumber)
}

func (l *Logger) Error(message string) {
	fileName, lineNumber := getCallerInfo()
	log(ERROR, &message, fileName, lineNumber)
}
