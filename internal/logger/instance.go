package logger

import "os"

func NewLogger() ILogger {
	return &Logger{
		minLevel: INFO,
	}
}

func New(minLevel LogLevel, filePath string, jsonFormat bool) (*Logger, error) {
	var file *os.File
	var err error
	if filePath != "" {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
	}
	return &Logger{
		minLevel: minLevel,
		file:     file,
		toJSON:   jsonFormat,
	}, nil
}
