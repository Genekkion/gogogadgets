package test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Prints the caller of the function for debugging purposes.
func printCaller() {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		fn := runtime.FuncForPC(pc)
		fmt.Printf("Called from %s:%d (%s)\n",
			filepath.Base(file), line, fn.Name())
	}
}

// Assert Checks if a condition is true and fails if not.
func Assert(t *testing.T, message string, got bool) {
	if got {
		return
	}

	t.Errorf("%s", message)
	fmt.Printf("Expression is false")

	printCaller()
	t.FailNow()
}

// AssertEqual Checks if two values are equal using reflect.DeepEqual and fails if not.
func AssertEqual[T any](t *testing.T, message string, expected T, got T) {
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("%s", message)
		fmt.Printf("Expected: %v\n", expected)
		fmt.Printf("Got     : %v\n", got)

		printCaller()
		t.FailNow()
	}
}

// NilErr Checks if err is nil, and if not, fails the test.
func NilErr(t *testing.T, err error) {
	if err == nil {
		return
	}

	t.Errorf("Expected nil error")
	fmt.Printf("Error: %v\n", err)

	printCaller()
	t.FailNow()
}

// NonNil Checks if the value is non-nil, and if not, fails the test.
func NonNil[T any](t *testing.T, value *T) {
	if value != nil {
		return
	}

	t.Errorf("Expected non-nil value")
	fmt.Printf("Value: %v\n", value)

	printCaller()
	t.FailNow()
}
