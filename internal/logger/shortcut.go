package logger

var DefaultLogger, _ = New(DEBUG, "", false)

func Debug(msg string, args ...interface{})   { DefaultLogger.Debug(msg, args...) }
func Info(msg string, args ...interface{})    { DefaultLogger.Info(msg, args...) }
func Success(msg string, args ...interface{}) { DefaultLogger.Success(msg, args...) }
func Warn(msg string, args ...interface{})    { DefaultLogger.Warn(msg, args...) }
func Error(msg string, args ...interface{})   { DefaultLogger.Error(msg, args...) }
