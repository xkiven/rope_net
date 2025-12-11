package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

// 使用这个包时自动init
func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

// Info 一般情况的日志
func Info(v ...interface{}) {
	infoLogger.Println(v...)
}

// Error 报错时的日志
func Error(v ...interface{}) {
	errorLogger.Println(v...)
}
