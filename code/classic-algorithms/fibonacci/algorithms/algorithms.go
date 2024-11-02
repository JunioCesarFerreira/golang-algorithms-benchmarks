package algorithms

import "math"

// Implementação Recursiva Simples (complexidade exponêncial)
func FibonacciRecursive(n int) int64 {
	if n <= 1 {
		return int64(n)
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// Programação Dinâmica com Memoization (Top-Down)
func FibonacciMemoizationTopDown(n int, memo map[int]int64) int64 {
	if n <= 1 {
		return int64(n)
	}
	if val, exists := memo[n]; exists {
		return val
	}
	memo[n] = FibonacciMemoizationTopDown(n-1, memo) + FibonacciMemoizationTopDown(n-2, memo)
	return memo[n]
}

// Programação Dinâmica com Tabulação (Bottom-Up)
func FibonacciBottomUp(n int) int64 {
	if n <= 1 {
		return int64(n)
	}
	fib := make([]int64, n+1)
	fib[0], fib[1] = 0, 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func FibonacciBottomUpOptimized(n int) int64 {
	if n <= 1 {
		return int64(n)
	}
	a, b := int64(0), int64(1)
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Estrutura para representar matrizes 2x2
type matrix2x2 struct {
	a, b, c, d int64
}

// Multiplicação de matrizes 2x2
func multiplyMatrix(m1, m2 matrix2x2) matrix2x2 {
	return matrix2x2{
		a: m1.a*m2.a + m1.b*m2.c,
		b: m1.a*m2.b + m1.b*m2.d,
		c: m1.c*m2.a + m1.d*m2.c,
		d: m1.c*m2.b + m1.d*m2.d,
	}
}

// Exponenciação de matrizes usando exponentiation by squaring
func matrixPower(m matrix2x2, n int) matrix2x2 {
	if n == 1 {
		return m
	}
	if n%2 == 0 {
		halfPower := matrixPower(m, n/2)
		return multiplyMatrix(halfPower, halfPower)
	} else {
		halfPower := matrixPower(m, n/2)
		return multiplyMatrix(multiplyMatrix(halfPower, halfPower), m)
	}
}

// Função principal para calcular Fibonacci usando matriz
func FibonacciMatrix(n int) int64 {
	if n == 0 {
		return 0
	}
	baseMatrix := matrix2x2{a: 1, b: 1, c: 1, d: 0}
	resultMatrix := matrixPower(baseMatrix, n-1)
	return resultMatrix.a
}

// Fórmula Fechada (Fórmula de Binet)
func FibonacciClosedForm(n int) int64 {
	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	result := math.Pow(phi, float64(n)) / sqrt5
	return int64(math.Round(result))
}
