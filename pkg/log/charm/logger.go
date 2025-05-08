package log

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"runtime"

	"github.com/Genekkion/gogogadgets/pkg/log"

	cl "github.com/charmbracelet/log"
)

const (
	defaultLevel = cl.DebugLevel
)

type Logger struct {
	logger      *cl.Logger
	destructors []Destructor
}

func New(w io.Writer, opts ...LoggerOption) Logger {
	l := Logger{
		logger: cl.NewWithOptions(w, cl.Options{
			ReportTimestamp: true,
			Level:           defaultLevel,
		}),
	}

	l.logger.SetStyles(defaultStyle)

	for _, opt := range opts {
		opt(&l)
	}

	return l
}

func NewFromFile(filePath string, opts ...LoggerOption) (*Logger, error) {
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

	l := New(file,
		append([]LoggerOption{WithDestructor(file.Close)}, opts...)...,
	)
	return &l, nil
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

func (l Logger) GetLevel() log.Level {
	return log.Level(FromCharm(l.logger.GetLevel()))
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

func (l Logger) printStackTrace() {
	buf := make([]byte, 1024*8)
	n := runtime.Stack(buf, false)
	l.Error("Stack trace:", "trace", string(buf[:n]))
}

// Wrapper function to log errors if they exist
func (l Logger) ErrorWrapper(err error) error {
	if err != nil {
		l.printStackTrace()
		l.Error("Program received error", "err", err)
	}
	return err
}

func (l Logger) FatalWrapper(err error) {
	if err != nil {
		l.printStackTrace()
		l.Fatal("Program received error", "err", err)
	}
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
