package logger

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

const (
	DEBUG   LogLevel = "DEBUG"
	INFO    LogLevel = "INFO"
	SUCCESS LogLevel = "SUCCESS"
	WARN    LogLevel = "WARN"
	ERROR   LogLevel = "ERROR"
)

var levelOrder = map[LogLevel]int{
	DEBUG:   0,
	INFO:    1,
	SUCCESS: 2,
	WARN:    3,
	ERROR:   4,
}

func (l *Logger) Close() {
	if l.file != nil {
		_ = l.file.Close()
	}
}

func (l *Logger) shouldLog(level LogLevel) bool {
	return levelOrder[level] >= levelOrder[l.minLevel]
}

func colorForLevel(level LogLevel) string {
	switch level {
	case INFO:
		return ColorCyan
	case WARN:
		return ColorYellow
	case ERROR:
		return ColorRed
	case SUCCESS:
		return ColorGreen
	case DEBUG:
		return ColorPurple
	default:
		return ColorWhite
	}
}

func stripColors(text string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(text, "")
}

func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if !l.shouldLog(level) {
		return
	}

	message := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	if l.toJSON {
		entry := logEntry{
			Timestamp: time.Now().Format(time.RFC3339),
			Level:     level,
			Message:   message,
		}
		jsonBytes, _ := json.Marshal(entry)
		l.output(string(jsonBytes), false)
	} else {
		color := colorForLevel(level)
		space := "-"
		final := fmt.Sprintf("%s[%s] [%s]%s %s %s", color, timestamp, level, ColorReset, space, message)
		l.output(final, true)
	}
}

func (l *Logger) output(line string, stripColor bool) {
	fmt.Println(line)

	if l.file != nil {
		go func() {
			l.mutex.Lock()
			defer l.mutex.Unlock()
			if stripColor {
				line = stripColors(line)
			}
			_, _ = l.file.WriteString(line + "\n")
		}()
	}
}
