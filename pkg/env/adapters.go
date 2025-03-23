package env

import (
	"strconv"
	"strings"
)

func IntAdapter(str string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(str))
}

func F64Adapter(str string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(str), 64)
}

func StrAdapter(str string) (string, error) {
	return str, nil
}

func BoolAdapter(str string) (bool, error) {
	return strconv.ParseBool(str)
}
