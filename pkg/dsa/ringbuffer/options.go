package ringbuffer

type Config struct {
	bufferSize   int
	resizeFactor float64
}

func defaultConfig() Config {
	return Config{
		bufferSize:   8,
		resizeFactor: 2,
	}
}

type RBOption func(*Config)

func WithBufferSize(n int) RBOption {
	return func(c *Config) {
		c.bufferSize = n
	}
}

func WithResizeFactor(factor float64) RBOption {
	return func(c *Config) {
		c.resizeFactor = max(1.5, factor)
	}
}
