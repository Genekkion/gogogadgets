package log

type Logger interface {
	Close() error

	Debug(msg string, kv ...any)
	Info(msg string, kv ...any)
	Warn(msg string, kv ...any)
	Error(msg string, kv ...any)
	Fatal(msg string, kv ...any)
}
