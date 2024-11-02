package main

import (
	"m/code/classic-algorithms/fibonacci/algorithms"
	"testing"
)

const n = 50

func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.FibonacciRecursive(n)
	}
}

func BenchmarkFibonacciMemoizationTopDown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memo := make(map[int]int64)
		algorithms.FibonacciMemoizationTopDown(n, memo)
	}
}

func BenchmarkFibonacciBottomUp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.FibonacciBottomUp(n)
	}
}

func BenchmarkFibonacciBottomUpOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.FibonacciBottomUpOptimized(n)
	}
}

func BenchmarkFibonacciMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.FibonacciMatrix(n)
	}
}

func BenchmarkFibonacciClosedForm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.FibonacciClosedForm(n)
	}
}
