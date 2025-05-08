package log

import (
	"os"

	"github.com/Genekkion/gogogadgets/pkg/log"
)

var defaultLogger = New(os.Stdout)

func SetLevel(level log.Level) {
	defaultLogger.SetLevel(level)
}

func GetLevel() log.Level {
	return defaultLogger.GetLevel()
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

func DebugCaller(msg string, kv ...any) {
	defaultLogger.DebugCaller(msg, kv...)
}

func ErrorWrapper(err error) error {
	return defaultLogger.ErrorWrapper(err)
}

func FatalWrapper(err error) {
	defaultLogger.FatalWrapper(err)
}
