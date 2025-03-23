package log

import "fmt"

type Level uint8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l Level) String() string {
	switch l {

	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"

	default:
		panic(fmt.Sprintf("invalid log level, v: %d", l))
	}
}

func (l Level) Tag() string {
	switch l {

	case LevelDebug:
		return "D"
	case LevelInfo:
		return "I"
	case LevelWarn:
		return "W"
	case LevelError:
		return "E"
	case LevelFatal:
		return "F"
	default:
		panic(fmt.Sprintf("invalid log level, v: %d", l))
	}
}
