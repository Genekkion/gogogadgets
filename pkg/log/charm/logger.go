package log

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"runtime"

	"github.com/Genekkion/gogogadgets/pkg/config"
	"github.com/Genekkion/gogogadgets/pkg/log"

	cl "github.com/charmbracelet/log"
)

const (
	defaultLevel = cl.DebugLevel
)

var (
	defaultLogger = New(os.Stdout)
)

type Logger struct {
	logger      *cl.Logger
	destructors []Destructor
}

func New(w io.Writer, opts ...LoggerOption) *Logger {
	l := &Logger{
		logger: cl.NewWithOptions(w, cl.Options{
			ReportTimestamp: true,
			Level:           defaultLevel,
		}),
	}

	l.logger.SetStyles(defaultStyle)

	for _, opt := range opts {
		opt(l)
	}

	return l
}

func NewFromFile(filePath string, opts ...config.Option) (*Logger, error) {
	stat, err := os.Stat(filePath)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	} else if stat.IsDir() {
		return nil, fmt.Errorf("directory found at file path specified, path: %s", filePath)
	}

	file, err := os.OpenFile(filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return New(file, WithDestructor(file.Close)), nil
}

func (l *Logger) Close() error {
	for i := len(l.destructors) - 1; i >= 0; i-- {
		err := l.destructors[i]()
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Logger) SetLevel(level log.Level) {
	l.logger.SetLevel(Level(level).ToCharm())
}

func SetLevel(level log.Level) {
	defaultLogger.SetLevel(level)
}

func (l Logger) GetLevel() log.Level {
	return log.Level(FromCharm(l.logger.GetLevel()))
}

func GetLevel() log.Level {
    return defaultLogger.GetLevel()
}

func (l Logger) Debug(msg string, kv ...any) {
	l.logger.Debug(msg, kv...)
}

func (l Logger) Info(msg string, kv ...any) {
	l.logger.Info(msg, kv...)
}

func (l Logger) Warn(msg string, kv ...any) {
	l.logger.Warn(msg, kv...)
}

func (l Logger) Error(msg string, kv ...any) {
	l.logger.Error(msg, kv...)
}

func (l Logger) Fatal(msg string, kv ...any) {
	l.logger.Fatal(msg, kv...)
}

func Debug(msg string, kv ...any) {
	defaultLogger.Debug(msg, kv...)
}

func Info(msg string, kv ...any) {
	defaultLogger.Info(msg, kv...)
}

func Warn(msg string, kv ...any) {
	defaultLogger.Warn(msg, kv...)
}

func Error(msg string, kv ...any) {
	defaultLogger.Error(msg, kv...)
}

func Fatal(msg string, kv ...any) {
	defaultLogger.Fatal(msg, kv...)
}

func DebugKV(kv ...any) {
	defaultLogger.Debug("", kv...)
}

func InfoKV(kv ...any) {
	defaultLogger.Info("", kv...)
}

func WarnKV(kv ...any) {
	defaultLogger.Warn("", kv...)
}

func ErrorKV(kv ...any) {
	defaultLogger.Error("", kv...)
}

func FatalKV(kv ...any) {
	defaultLogger.Fatal("", kv...)
}

func getSkipCount(skip ...int) int {
	var v int
	if len(skip) == 0 {
		v = 1
	} else {
		v = skip[0]
	}
	return v
}

func (l Logger) printStackTrace(skip ...int) {
	pc, file, line, ok := runtime.Caller(getSkipCount(skip...))
	if !ok {
		l.Error("Error retrieving stack trace.")
	} else {
		l.Error("Stack trace:",
			"function", runtime.FuncForPC(pc).Name(),
			"file", file,
			"line", line,
		)
	}
}

// Wrapper function to log errors if they exist
func (l Logger) ErrorWrapper(err error, skip ...int) error {
	if err != nil {
		l.printStackTrace(skip...)
		l.Error("Program received error", "err", err)
	}
	return err
}

func ErrorWrapper(err error) error {
	return defaultLogger.ErrorWrapper(err)
}

func (l Logger) FatalWrapper(err error, skip ...int) {
	if err != nil {
		l.printStackTrace(skip...)
		l.Fatal("Program received error", "err", err)
	}
}

func FatalWrapper(err error) {
	defaultLogger.FatalWrapper(err, 2)
}

func (l Logger) DebugCaller(msg string, kv ...any) {
	l.Debug(msg, kv...)
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		l.Debug("Error retrieving stack trace.")
	} else {
		l.Debug("Stack trace:",
			"function", runtime.FuncForPC(pc).Name(),
			"file", file,
			"line", line,
		)
	}
}

func DebugCaller(msg string, kv ...any) {
	defaultLogger.DebugCaller(msg, kv...)
}
