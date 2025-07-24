package logger

// Public methods
func (l *Logger) Debug(msg string, args ...interface{})   { l.log(DEBUG, msg, args...) }
func (l *Logger) Info(msg string, args ...interface{})    { l.log(INFO, msg, args...) }
func (l *Logger) Success(msg string, args ...interface{}) { l.log(SUCCESS, msg, args...) }
func (l *Logger) Warn(msg string, args ...interface{})    { l.log(WARN, msg, args...) }
func (l *Logger) Error(msg string, args ...interface{})   { l.log(ERROR, msg, args...) }
