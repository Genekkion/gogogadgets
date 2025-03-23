package log

import (
	"fmt"

	"github.com/Genekkion/gogogadgets/pkg/log"

	cl "github.com/charmbracelet/log"
)

const (
	defaultLogLevel Level = Level(log.LevelDebug)

	LevelDebug = log.LevelDebug
	LevelInfo  = log.LevelInfo
	LevelWarn  = log.LevelWarn
	LevelError = log.LevelError
	LevelFatal = log.LevelFatal
)

type Level log.Level

func (l Level) ToCharm() cl.Level {
	switch log.Level(l) {
	case log.LevelDebug:
		return cl.DebugLevel
	case log.LevelInfo:
		return cl.InfoLevel
	case log.LevelWarn:
		return cl.WarnLevel
	case log.LevelError:
		return cl.ErrorLevel
	case log.LevelFatal:
		return cl.FatalLevel
	default:
		panic(fmt.Sprintf("invalid log level, v: %d", l))
	}
}

func FromCharm(level cl.Level) log.Level {
	switch level {
	case cl.DebugLevel:
		return log.LevelDebug
	case cl.InfoLevel:
		return log.LevelInfo
	case cl.WarnLevel:
		return log.LevelWarn
	case cl.ErrorLevel:
		return log.LevelError
	case cl.FatalLevel:
		return log.LevelFatal
	default:
		panic(fmt.Sprintf("invalid log level, v: %d", level))
	}
}
