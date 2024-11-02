# Fibonacci Algorithm Implementations in Go

This directory contains various implementations of the Fibonacci sequence calculation in Go, utilizing different approaches and optimization techniques.

## Theoretical Explanation of the Fibonacci Sequence

The Fibonacci sequence is a series of numbers where each number ($F_n$) is the sum of the two preceding ones:

- $F_0 = 0$
- $F_1 = 1$
- $F_n = F_{n-1} + F_{n-2}$ for $n \ge 2$

This generates the sequence:

$0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...$

### Properties and Applications

- **Natural Occurrences:** The Fibonacci sequence appears in various natural phenomena, such as the branching patterns of trees, the arrangement of leaves on a stem, the flowering of artichokes, and the spiral patterns of shells and galaxies.
- **Golden Ratio Connection:** As the sequence progresses, the ratio of consecutive Fibonacci numbers approximates the golden ratio (approximately 1.61803398875). Mathematically:

$$
\lim_{n \to \infty} \frac{F_{n+1}}{F_n} = \phi
$$

where **$\phi$ (phi)** is the golden ratio.


## Implementations

1. **Simple Recursive Implementation**

   - **Function:** `FibonacciRecursive(n int) int64`
   - **Description:** Uses direct recursion without optimizations. It has exponential complexity O(2^n), making it inefficient for large values of `n`.

2. **Dynamic Programming with Memoization (Top-Down)**

   - **Function:** `FibonacciMemoizationTopDown(n int, memo map[int]int64) int64`
   - **Description:** Uses memoization to store intermediate results, avoiding redundant calculations. Complexity is reduced to O(n).

3. **Dynamic Programming with Tabulation (Bottom-Up)**

   - **Function:** `FibonacciBottomUp(n int) int64`
   - **Description:** Calculates the sequence iteratively, building up a table of results. Has O(n) complexity but uses more memory due to storing the entire table.

4. **Optimized Dynamic Programming (Bottom-Up with Space Optimization)**

   - **Function:** `FibonacciBottomUpOptimized(n int) int64`
   - **Description:** An optimized variant that uses only two variables to store previous values, reducing memory usage.

5. **Matrix Exponentiation**

   - **Function:** `FibonacciMatrix(n int) int64`
   - **Description:** Uses matrix exponentiation to calculate the nth Fibonacci number in O(log n) time.
   - **Auxiliary Functions:** Functions for manipulating and multiplying 2x2 matrices.

6. **Closed-Form Expression (Binet's Formula)**

   - **Function:** `FibonacciClosedForm(n int) int64`
   - **Description:** Uses the closed-form mathematical formula to calculate the nth Fibonacci number. Although it has constant complexity O(1), it may suffer from inaccuracies due to floating-point limitations.

## How to Use

### Run the Main Program

```shell
cd ./code/classic-algorithms/fibonacci
go run main.go
```

### Run the Benchmarks

```shell
cd ./code/classic-algorithms/fibonacci
go test -benchmem -bench .
```

## Benchmark Result

For N=50:
```
goos: windows
goarch: amd64
pkg: m/code/classic-algorithms/fibonacci
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkFibonacciRecursive-8                          1        123006930300 ns/op             8 B/op          1 allocs/op
BenchmarkFibonacciMemoizationTopDown-8            105393             10953 ns/op            5029 B/op          9 allocs/op
BenchmarkFibonacciBottomUp-8                     3927403               307.3 ns/op           416 B/op          1 allocs/op
BenchmarkFibonacciBottomUpOptimized-8           20807792                57.01 ns/op            0 B/op          0 allocs/op
BenchmarkFibonacciMatrix-8                      29267293                36.11 ns/op            0 B/op          0 allocs/op
BenchmarkFibonacciClosedForm-8                  20895547                58.80 ns/op            0 B/op          0 allocs/op
PASS
ok      m/code/classic-algorithms/fibonacci     129.640s
```

For N=500000:
```
goos: windows
goarch: amd64
pkg: m/code/classic-algorithms/fibonacci
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkFibonacciMemoizationTopDown-8                 4         250530025 ns/op        42481900 B/op      11860 allocs/op
BenchmarkFibonacciBottomUp-8                         520           2326206 ns/op         4005902 B/op          1 allocs/op
BenchmarkFibonacciBottomUpOptimized-8               4137            303889 ns/op               0 B/op          0 allocs/op
BenchmarkFibonacciMatrix-8                       6185257               199.0 ns/op             0 B/op          0 allocs/op
BenchmarkFibonacciClosedForm-8                  13824724                72.79 ns/op            0 B/op          0 allocs/op
PASS
ok      m/code/classic-algorithms/fibonacci     8.501s
```
In this case we remove the recursive method because for N>50 this is unfeasible.

## Performance Comparison

The different implementations have varying performance characteristics:

- **Simple Recursive:** Inefficient for large values of `n` due to exponential complexity.
- **Memoization (Top-Down):** Improves performance to O(n) by storing intermediate results.
- **Bottom-Up:** Also has O(n) complexity, building the solution iteratively.
- **Optimized Bottom-Up:** Reduces memory usage by using only two variables.
- **Matrix Exponentiation:** Offers O(log n) performance, efficient for very large `n`.
- **Closed-Form Formula:** Has constant complexity but may not be accurate for large `n` due to floating-point precision limitations.

## Notes

- **Accuracy:** For very large values of `n`, it is recommended to avoid the Closed-Form implementation due to potential rounding errors.
- **Choosing an Implementation:** Depending on your use case (size of `n`, memory constraints, performance needs), different implementations may be more suitable.
s