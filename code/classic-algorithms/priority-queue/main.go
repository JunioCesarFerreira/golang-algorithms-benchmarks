package main

import (
	"fmt"
)

// Item representa um elemento da fila de prioridade.
type Item struct {
	Value    string // O valor do item
	Priority int    // Prioridade (quanto menor, mais alta)
}

// MinHeap representa uma heap mínima baseada em array.
type MinHeap struct {
	arr []*Item
}

// Retorna o índice do pai de um nó
func parent(i int) int {
	return (i - 1) / 2
}

// Retorna os índices dos filhos esquerdo e direito
func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

// Insert adiciona um novo item ao heap
func (h *MinHeap) Insert(item *Item) {
	h.arr = append(h.arr, item)
	h.heapifyUp(len(h.arr) - 1)
}

// heapifyUp reorganiza a heap após inserção
func (h *MinHeap) heapifyUp(i int) {
	for i > 0 && h.arr[parent(i)].Priority > h.arr[i].Priority {
		h.arr[parent(i)], h.arr[i] = h.arr[i], h.arr[parent(i)]
		i = parent(i)
	}
}

// ExtractMin remove e retorna o item de maior prioridade
func (h *MinHeap) ExtractMin() (*Item, error) {
	if len(h.arr) == 0 {
		return nil, fmt.Errorf("fila vazia")
	}

	min := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1] // Substitui a raiz pelo último elemento
	h.arr = h.arr[:len(h.arr)-1]   // Remove o último elemento
	h.heapifyDown(0)               // Reorganiza o heap

	return min, nil
}

// heapifyDown reorganiza a heap após remoção
func (h *MinHeap) heapifyDown(i int) {
	smallest := i
	left := leftChild(i)
	right := rightChild(i)

	if left < len(h.arr) && h.arr[left].Priority < h.arr[smallest].Priority {
		smallest = left
	}
	if right < len(h.arr) && h.arr[right].Priority < h.arr[smallest].Priority {
		smallest = right
	}

	if smallest != i {
		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		h.heapifyDown(smallest)
	}
}

// PrintHeap exibe a heap no formato de array
func (h *MinHeap) PrintHeap() {
	fmt.Print("Heap atual: ")
	for _, item := range h.arr {
		fmt.Printf("(%s, %d) ", item.Value, item.Priority)
	}
	fmt.Println()
}

// PriorityQueue baseada na MinHeap
type PriorityQueue struct {
	heap MinHeap
}

// Enqueue adiciona um elemento à fila com uma prioridade específica
func (pq *PriorityQueue) Enqueue(value string, priority int) {
	item := &Item{Value: value, Priority: priority}
	pq.heap.Insert(item)
}

// Dequeue remove o item com maior prioridade (menor valor de Priority)
func (pq *PriorityQueue) Dequeue() (*Item, error) {
	return pq.heap.ExtractMin()
}

// PrintQueue exibe o estado atual da fila de prioridade
func (pq *PriorityQueue) PrintQueue() {
	pq.heap.PrintHeap()
}

func main() {
	pq := &PriorityQueue{}

	// Inserindo elementos na fila
	pq.Enqueue("Tarefa 1", 3)
	pq.Enqueue("Tarefa 2", 1)
	pq.Enqueue("Tarefa 3", 4)
	pq.Enqueue("Tarefa 4", 2)

	// Mostrando a fila antes das remoções
	pq.PrintQueue()

	// Extraindo elementos com maior prioridade (menor valor de Priority)
	fmt.Println("\nRemovendo elementos da fila:")
	for len(pq.heap.arr) > 0 {
		item, _ := pq.Dequeue()
		fmt.Printf("Removido: (%s, %d)\n", item.Value, item.Priority)
		pq.PrintQueue()
	}
}
