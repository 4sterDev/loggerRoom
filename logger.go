package loggerRoom

import (
	"fmt"
	"time"
)

type Logger struct {
	Colorize        bool
	LogLevel        LogLevel
	ShowLogLevel    bool
	ShowLogTimeDate bool
	ShowLogTimeTime bool
	ShowTag         bool
}

type LogMessage struct {
	Tag     string
	Message []any
}

func buildLogString(l *Logger, ll LogLevel, tag string, msg ...any) string {
	message := ""
	now := time.Now()
	if l.ShowLogTimeDate {
		message += fmt.Sprintf("%d-%02d-%02d ", now.Year(), now.Month(), now.Day())
	}
	if l.ShowLogTimeTime {
		message += fmt.Sprintf("%02d:%02d:%02d ", now.Hour(), now.Minute(), now.Second())
	}
	if l.ShowLogLevel {
		if l.Colorize {
			message += ll.ColorString() + " "
		} else {
			message += ll.String() + " "
		}
	}
	if l.ShowTag {
		if l.Colorize {
			message += colorize(tag, MAGENTA) + " "
		} else {
			message += tag
		}
	}

	for i := range msg {
		message += formatMessage(msg[i])
	}

	return message
}

func formatMessage(msg any) string {
	switch v := msg.(type) {
	case string:
		return fmt.Sprintf(" %s", v)
	case int:
		return fmt.Sprintf(" %d", v)
	case time.Duration:
		return fmt.Sprintf(" %.3fµs", float32(v)/1000)
	default:
		return ""
	}
}

func (l *Logger) Info(tag string, msg ...any) {
	fmt.Println(buildLogString(l, INFO, tag, msg...))
}

func (l *Logger) Debug(tag string, msg ...any) {
	if l.LogLevel == DEBUG {
		fmt.Println(buildLogString(l, DEBUG, tag, msg...))
	}
}

func (l *Logger) Error(msg LogMessage) {
	fmt.Println(buildLogString(l, ERROR, msg.Tag, msg.Message...))
}
