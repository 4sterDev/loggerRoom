package loggerRoom

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	INFO LogLevel = iota
	ERROR
	DEBUG
)

type Logger struct {
	Colorize        bool
	LogLevel        LogLevel
	ShowLogLevel    bool
	ShowLogTimeDate bool
	ShowLogTimeTime bool
}

func (l *Logger) Info(msg ...any) {
	level := "INFO"
	if l.Colorize {
		level = colorize(level, YELLOW)
	}

	printTime(l.ShowLogTimeDate, l.ShowLogTimeTime)
	fmt.Printf("%s", level)
	for i := range msg {
		printMessage(msg[i])
	}
	fmt.Printf("\n")
}

func (l *Logger) Debug(msg ...interface{}) {
	if l.LogLevel == DEBUG {
		level := "DEBUG"
		if l.Colorize {
			level = colorize(level, GREEN)
		}

		printTime(l.ShowLogTimeDate, l.ShowLogTimeTime)
		fmt.Printf("%s", level)
		for i := range msg {
			printMessage(msg[i])
		}
		fmt.Printf("\n")
	}
}

func (l *Logger) Error(msg ...interface{}) {
	level := "ERROR"
	if l.Colorize {
		level = colorize(level, RED)
	}

	printTime(l.ShowLogTimeDate, l.ShowLogTimeTime)
	fmt.Printf("%s", level)
	for i := range msg {
		printMessage(msg[i])
	}
	fmt.Printf("\n")
}

func printTime(d bool, t bool) {
	now := time.Now()

	if d {
		fmt.Printf("%d-%02d-%02d ", now.Year(), now.Month(), now.Day())
	}

	if t {
		fmt.Printf("%02d:%02d:%02d ", now.Hour(), now.Minute(), now.Second())
	}
}

func printMessage(msg any) {
	switch v := msg.(type) {
	case string:
		fmt.Printf(" %s", v)
	case int:
		fmt.Printf(" %d", v)
	case time.Duration:
		fmt.Printf(" %.3fµs", float32(v)/1000)
	}
}
