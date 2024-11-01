package main

import (
	"math/rand"
	"sort"
	"testing"
)

const size = 100000

// generateTestData generates a slice and a map with random integers.
func generateTestData(n int) ([]int, int) {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		num := rand.Intn(n * 10)
		slice[i] = num
	}
	return slice, slice[size-1]
}

// BenchmarkSequentialSearch benchmarks searching for a value in a slice.
func BenchmarkSequentialSearch(b *testing.B) {
	slice, target := generateTestData(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found := false
		for _, v := range slice {
			if v == target {
				found = true
				break
			}
		}
		_ = found
	}
}

// BenchmarkMapLookup benchmarks looking up a value in a map.
func BenchmarkMapLookup(b *testing.B) {
	slice, target := generateTestData(size)

	n := len(slice)
	m := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		num := slice[i]
		m[num] = true
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, found := m[target]
		_ = found
	}
}

// BenchmarkBinarySearch benchmarks searching for a value in a sorted slice using binary search.
func BenchmarkBinarySearch(b *testing.B) {
	slice, target := generateTestData(size)

	sort.Ints(slice)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := sort.SearchInts(slice, target)
		found := index < len(slice) && slice[index] == target
		_ = found
	}
}

// BenchmarkMapLookup benchmarks looking up a value in a map.
func BenchmarkMapLookupConstruction(b *testing.B) {
	slice, target := generateTestData(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(slice)
		m := make(map[int]bool, n)
		for i := 0; i < n; i++ {
			num := slice[i]
			m[num] = true
		}
		_, found := m[target]
		_ = found
	}
}

// BenchmarkBinarySearch benchmarks searching for a value in a sorted slice using binary search.
func BenchmarkBinarySearchConstruction(b *testing.B) {
	slice, target := generateTestData(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(slice)
		index := sort.SearchInts(slice, target)
		found := index < len(slice) && slice[index] == target
		_ = found
	}
}
