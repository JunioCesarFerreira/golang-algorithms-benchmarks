package main

import (
	"fmt"
	"m/code/classic-algorithms/fibonacci/algorithms"
	"runtime"
	"time"
)

// MeasureFunction é uma função genérica que mede o tempo e a memória utilizados por uma função passada como parâmetro.
func MeasureFunction(function func()) {
	// Força a coleta de lixo para uma medição mais precisa de memória.
	runtime.GC()

	// Lê o estado da memória antes da execução.
	var memStatsBefore runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	// Inicia a medição do tempo.
	startTime := time.Now()

	// Executa a função alvo.
	function()

	// Calcula o tempo decorrido.
	elapsedTime := time.Since(startTime)

	// Lê o estado da memória após a execução.
	var memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsAfter)

	// Calcula a memória utilizada.
	memUsed := memStatsAfter.Alloc - memStatsBefore.Alloc

	// Exibe os resultados.
	fmt.Printf("\tTempo de execução: %s\n", elapsedTime)
	fmt.Printf("\tMemória utilizada: %d bytes\n", memUsed)
}

func main() {

	n := 45 // Altere o valor de n conforme necessário

	fmt.Printf("Calculando o %dº número de Fibonacci:\n\n", n)

	// 1. Implementação Recursiva Simples
	MeasureFunction(func() {
		fmt.Println("1. Implementação Recursiva Simples:")
		fmt.Printf("\tResultado: %d\n\n", algorithms.FibonacciRecursive(n))
	})
	fmt.Print("---\n\n")

	// 2. Programação Dinâmica com Memoization (Top-Down)
	MeasureFunction(func() {
		fmt.Println("2. Memoization (Top-Down):")
		memo := make(map[int]int64)
		fmt.Printf("\tResultado: %d\n\n", algorithms.FibonacciMemoizationTopDown(n, memo))
	})
	fmt.Print("---\n\n")

	// 3. Programação Dinâmica com Tabulação (Bottom-Up)
	MeasureFunction(func() {
		fmt.Println("3. Tabulação (Bottom-Up):")
		fmt.Printf("\tResultado: %d\n\n", algorithms.FibonacciBottomUp(n))
	})
	fmt.Print("---\n\n")

	// 4. Algoritmo de Exponenciação de Matrizes
	MeasureFunction(func() {
		fmt.Println("4. Exponenciação de Matrizes:")
		fmt.Printf("\tResultado: %d\n\n", algorithms.FibonacciMatrix(n))
	})
	fmt.Print("---\n\n")

	// 5. Fórmula Fechada (Fórmula de Binet)
	MeasureFunction(func() {
		fmt.Println("5. Fórmula Fechada (Fórmula de Binet):")
		fmt.Printf("\tResultado: %d\n\n", algorithms.FibonacciClosedForm(n))
	})
	fmt.Print("---\n\n")
}
