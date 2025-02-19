package main

import (
	"fmt"
)

// MinHeap estrutura a árvore binária em forma de array
type MinHeap struct {
	arr []int
}

// Retorna o índice do pai de um nó
func parent(i int) int {
	return (i - 1) / 2
}

// Retorna os índices dos filhos esquerdo e direito de um nó
func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

// Insert adiciona um novo elemento ao heap
func (h *MinHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

// heapifyUp reorganiza a heap após inserção
func (h *MinHeap) heapifyUp(i int) {
	for i > 0 && h.arr[parent(i)] > h.arr[i] {
		h.arr[parent(i)], h.arr[i] = h.arr[i], h.arr[parent(i)]
		i = parent(i)
	}
}

// ExtractMin remove e retorna o menor elemento do heap
func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.arr) == 0 {
		return 0, fmt.Errorf("heap vazia")
	}

	min := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1] // Substitui a raiz pelo último elemento
	h.arr = h.arr[:len(h.arr)-1]   // Remove o último elemento
	h.heapifyDown(0)               // Reorganiza o heap

	return min, nil
}

// heapifyDown reorganiza a heap após a remoção
func (h *MinHeap) heapifyDown(i int) {
	smallest := i
	left := leftChild(i)
	right := rightChild(i)

	if left < len(h.arr) && h.arr[left] < h.arr[smallest] {
		smallest = left
	}
	if right < len(h.arr) && h.arr[right] < h.arr[smallest] {
		smallest = right
	}

	if smallest != i {
		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		h.heapifyDown(smallest)
	}
}

// PrintHeap exibe a heap no formato de array
func (h *MinHeap) PrintHeap() {
	fmt.Println("Heap:", h.arr)
}

func main() {
	heap := &MinHeap{}

	// Inserindo valores na heap
	values := []int{10, 4, 15, 20, 1, 7}
	for _, v := range values {
		heap.Insert(v)
	}

	heap.PrintHeap() // Exibe a heap

	// Extraindo elementos na ordem de prioridade (menor primeiro)
	for len(heap.arr) > 0 {
		min, _ := heap.ExtractMin()
		fmt.Println("Removido:", min)
		heap.PrintHeap()
	}
}
