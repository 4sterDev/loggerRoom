package loggerRoom

type Color int

const (
	RED Color = iota
	GREEN
	YELLOW
	RESET
)

func (c Color) String() string {
	switch c {
	case RED:
		return "\033[31m"
	case GREEN:
		return "\033[32m"
	case YELLOW:
		return "\033[33m"
	case RESET:
		return "\033[0m"
	}
	return ""
}

func colorize(msg string, color Color) string {
	return color.String() + msg + RESET.String()
}
