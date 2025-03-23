package log

import (
	"github.com/Genekkion/gogogadgets/pkg/log"
	cl "github.com/charmbracelet/log"
)

type LoggerOption func(*Logger)

func WithLevel(level log.Level) LoggerOption {
	return func(l *Logger) {
		l.SetLevel(level)
	}
}

type Destructor func() error

func WithDestructor(destructors ...Destructor) LoggerOption {
	return func(l *Logger) {
		l.destructors = append(l.destructors, destructors...)
	}
}

func WithStyle(styles *cl.Styles) LoggerOption {
	return func(l *Logger) {
		l.logger.SetStyles(styles)
	}
}

func WithReportTimestamp(flag bool) LoggerOption {
	return func(l *Logger) {
		l.logger.SetReportTimestamp(flag)
	}
}

func WithReportCaller(flag bool) LoggerOption {
	return func(l *Logger) {
		l.logger.SetReportCaller(flag)
	}
}
