# Go Map vs Slice Search Benchmark

This directory contains a Go program that benchmarks three methods of searching for an element:

1. Sequential search in a slice
2. Lookup in a map
3. Binary search in a sorted slice

## How to Run the Benchmarks

Execute the following command in your terminal:

```bash
go test -benchmem -bench .
```

where:
- `-bench .` runs all benchmark functions.
- `-benchmem` includes memory allocation statistics.

## Result

For size=10^5
```
goos: windows
goarch: amd64
pkg: m/code/golang/maps
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSequentialSearch-8         58882        60649 ns/op           0 B/op          0 allocs/op
BenchmarkMapLookup-8            175301660            6.671 ns/op       0 B/op          0 allocs/op
BenchmarkBinarySearch-8          29268220           42.44 ns/op        0 B/op          0 allocs/op
PASS
ok      m/code/golang/maps      8.331s
```

For size=10^6
```
goos: windows
goarch: amd64
pkg: m/code/golang/maps
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSequentialSearch-8         1230       908568 ns/op            0 B/op          0 allocs/op
BenchmarkMapLookup-8           161427666            7.014 ns/op        0 B/op          0 allocs/op
BenchmarkBinarySearch-8         25511613           76.27 ns/op         0 B/op          0 allocs/op
PASS
ok      m/code/golang/maps      6.901s
```

Using a map for lookups is significantly more efficient in terms of speed compared to sequential search in a slice. Binary search in a sorted slice is faster than sequential search but still slower than map lookups.

## Explanation

**Why are map lookups faster than sequential search in slices in the benchmark results?**

### 1. Time Complexity Differences

- **Sequential Search in Slice**
  - **Time Complexity:** O(n)
  - The time it takes to find an element increases linearly with the number of elements.
  - In the benchmark (`BenchmarkSequentialSearch-8`), the average time per operation is **85,392 ns/op**, which is relatively high due to the linear scan through the slice.

- **Map Lookup**
  - **Time Complexity:** O(1) on average
  - Maps in Go are implemented as hash tables, providing constant-time average complexity for lookups.
  - In the benchmark (`BenchmarkMapLookup-8`), the average time per operation is **6.689 ns/op**, which is significantly faster.

- **Binary Search in Sorted Slice**
  - **Time Complexity:** O(log n)
  - Requires the slice to be sorted, but provides faster search times than sequential search.
  - In the benchmark (`BenchmarkBinarySearch-8`), the average time per operation is **54.08 ns/op**, faster than sequential search but slower than map lookups.

### 2. Implementation Overheads

- **Sequential Search**
  - Involves iterating over each element and comparing it with the target value.
  - Higher CPU time due to multiple comparisons.
  - No additional memory allocations (`0 B/op`, `0 allocs/op`).

- **Map Lookup**
  - Uses hashing to directly access the value associated with a key.
  - Slight overhead due to hash function computation, but still much faster.
  - No additional memory allocations during lookup.

- **Binary Search**
  - Requires data to be sorted, which may add overhead if data changes frequently.
  - Each lookup involves fewer comparisons than sequential search.
  - No additional memory allocations during search.

### 3. Memory Allocations

All benchmarks show **`0 B/op`** and **`0 allocs/op`**, indicating no additional memory is allocated during search operations.

## Conclusion

- **Maps Provide Fastest Lookups**
  - Map lookups are the fastest due to constant-time complexity.
  - Ideal for large datasets with frequent searches.

- **Binary Search as a Middle Ground**
  - Faster than sequential search.
  - Useful when you can maintain a sorted slice and want to avoid the overhead of a map.

- **Sequential Search for Simplicity**
  - Suitable for small datasets where the overhead of a map or sorting isn't justified.
  - Simpler implementation but not efficient for large datasets.

### Key Takeaways

- **Choose Maps for Performance-Critical Searches**
  - When quick lookup times are essential, and you can afford the memory to maintain a map.

- **Use Binary Search with Sorted Data**
  - Offers a good balance when data can be kept sorted and memory usage is a concern.

- **Avoid Sequential Search in Large Datasets**
  - Performance degrades linearly with the number of elements, making it inefficient for large collections.

### Additional Notes

- **Consider Data Mutability**
  - If the dataset changes frequently, maintaining a sorted slice for binary search may introduce overhead.
  - Maps handle dynamic data well but come with memory overhead.

- **Memory vs. Speed Trade-off**
  - Maps consume more memory but offer the fastest lookup speed.
  - Slices use less memory but may result in slower searches depending on the method.

Additional test results with construction of structures:
```
goos: windows
goarch: amd64
pkg: m/code/golang/maps
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSequentialSearch-8                38222             30663 ns/op               0 B/op          0 allocs/op
BenchmarkMapLookup-8                    148374537               16.37 ns/op            0 B/op          0 allocs/op
BenchmarkBinarySearch-8                 10214425                99.66 ns/op            0 B/op          0 allocs/op
BenchmarkMapLookupConstruction-8             152           9270245 ns/op         3065902 B/op         11 allocs/op
BenchmarkBinarySearchConstruction-8         1952            650823 ns/op              24 B/op          1 allocs/op
PASS
ok      m/code/golang/maps      10.464s
```
