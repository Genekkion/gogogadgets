package env

import (
	"os"

	log "github.com/Genekkion/gogogadgets/pkg/log/charm"
)

type Adapter[T any] func(str string) (T, error)

func GetEnvOrDef[T any](key string, defV T, adapters ...Adapter[T]) T {
	str, ok := os.LookupEnv(key)
	if !ok {
		return defV
	}

	for _, adapter := range adapters {
		v, err := adapter(str)
		if err == nil {
			return v
		}
	}

	log.Warn("Unable to adapt environment variable, using default value",
		"key", key, "default", defV)
	return defV
}

func GetEnvOrFatal[T any](key string, adapters ...Adapter[T]) T {
	str, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("Required environment variable not set, exiting", "key", key)
	}
	for _, adapter := range adapters {
		v, err := adapter(str)
		if err == nil {
			return v
		}
	}

	log.Fatal("Required environment variable not set, exiting", "key", key)
	panic("should not reach here")
}

func GetStringEnvOrFatal(key string) string {
	return GetEnvOrFatal(key, StrAdapter)
}

func GetStringEnv(key string, defV string) string {
	return GetEnvOrDef(key, defV, StrAdapter)
}

func GetIntEnvOrFatal(key string) int {
	return GetEnvOrFatal(key, IntAdapter)
}

func GetIntEnvDef(key string, defV int) int {
	return GetEnvOrDef(key, defV, IntAdapter)
}

func GetF64EnvOrFatal(key string) float64 {
	return GetEnvOrFatal(key, F64Adapter)
}

func GetF64EnvDef(key string, defV float64) float64 {
	return GetEnvOrDef(key, defV, F64Adapter)
}

func GetBoolEnvOrFatal(key string) bool {
	return GetEnvOrFatal(key, BoolAdapter)
}

func GetBoolEnvDef(key string, defV bool) bool {
	return GetEnvOrDef(key, defV, BoolAdapter)
}
