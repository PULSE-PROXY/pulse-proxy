package logger

import (
	"os"
	"sync"
)

type LogLevel string

type logEntry struct {
	Timestamp string   `json:"timestamp"`
	Level     LogLevel `json:"level"`
	Message   string   `json:"message"`
}

type ILogger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Success(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type Logger struct {
	minLevel LogLevel
	file     *os.File
	mutex    sync.Mutex
	toJSON   bool
}
