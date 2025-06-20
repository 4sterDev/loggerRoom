package loggerRoom

type LogLevel int

const (
	INFO LogLevel = iota
	ERROR
	DEBUG
)

var levelName = map[LogLevel]string{
	INFO:  "INFO",
	ERROR: "ERROR",
	DEBUG: "DEBUG",
}

func (ll LogLevel) String() string {
	return levelName[ll]
}

func (ll LogLevel) ColorString() string {
	switch ll {
	case INFO:
		return colorize(ll.String(), GREEN)
	case ERROR:
		return colorize(ll.String(), RED)
	case DEBUG:
		return colorize(ll.String(), YELLOW)
	default:
		return ""
	}
}
