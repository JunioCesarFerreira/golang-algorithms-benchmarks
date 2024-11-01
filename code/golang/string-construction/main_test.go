package main

import (
	"fmt"
	"strconv"
	"testing"
)

// Benchmark using fmt.Sprintf
func BenchmarkSprintfNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("This is a test string number %d", i)
	}
}

// Benchmark using string concatenation
func BenchmarkConcatNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "This is a test string number " + strconv.Itoa(i)
	}
}

// Benchmark using fmt.Sprintf
func BenchmarkSprintfStringNumber(b *testing.B) {
	const test = "test"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("This is a %s string number %d", test, i)
	}
}

// Benchmark using string concatenation
func BenchmarkConcatStringNumber(b *testing.B) {
	const test = "test"
	for i := 0; i < b.N; i++ {
		_ = "This is a " + test + " string number " + strconv.Itoa(i)
	}
}

// Benchmark using fmt.Sprintf
func BenchmarkSprintfString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := strconv.Itoa(i)
		_ = fmt.Sprintf("This is a %s string", test)
	}
}

// Benchmark using string concatenation
func BenchmarkConcatString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := strconv.Itoa(i)
		_ = "This is a " + test + " string"
	}
}
