package main

import (
	"errors"
	"fmt"
)

type Compare[T any] func(a, b T) bool

////// Worker /////////
type Worker struct {
	Name string
	Id   uint8
}

func (c *Worker) GetID() int {
	return int(c.Id)
}

func (c *Worker) GetName() string {
	return c.Name
}

/////////Heap/////////////
type heap[T any] struct {
	currentSize int
	heapArray   []T
	maxSize     int
	comp        Compare[T]
}

func newHeapFromSlice[T any](values []T, comp Compare[T]) *heap[T] {
	maxSize := len(values)
	heap := &heap[T]{
		maxSize:     maxSize,
		currentSize: len(values),
		heapArray:   values,
		comp:        comp,
	}
	// heap.trickleUp(len(values) - 1)
	heap.trickleDown(0)
	return heap
}

func (h *heap[T]) trickleDown(index int) {
	var largerChild int
	top := h.heapArray[index]     // Сохранение корня
	for index < h.currentSize/2 { // Пока у узла имеется хотя бы один потомок
		var leftChild int = 2*index + 1
		var rightChild int = leftChild + 1
		// Определение большего потомка
		if rightChild < h.currentSize && // (Правый потомок существует?)
			h.comp(h.heapArray[leftChild], h.heapArray[rightChild]) {
			largerChild = rightChild
		} else {
			largerChild = leftChild
		}

		if !h.comp(top, h.heapArray[largerChild]) {
			break
		}

		// Потомок сдвигается вверх
		h.heapArray[index] = h.heapArray[largerChild]
		index = largerChild // Переход вниз
	}
	h.heapArray[index] = top // index <- корень
}

func (h *heap[T]) remove() (T, error) {
	if h.currentSize <= 0 {
		return *new(T), errors.New("heap is empty")
	}
	root := h.heapArray[0]
	h.currentSize--
	h.heapArray[0] = h.heapArray[h.currentSize]
	h.trickleDown(0)
	return root, nil
}

/////////Heap Sort/////////////
func HeapSort[T any](arr []T, comp Compare[T]) ([]T, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}
	heap := newHeapFromSlice(arr, comp)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], _ = heap.remove()
	}
	return arr, nil
}

func main() {
	workerSlice := []Worker{{"Julie", 1}, {"Alex", 2}, {"Tom", 4},
		{"George", 3}, {"Max", 60}, {"Tommy", 94}, {"William", 12},
		{"Sophia", 14}, {"Oliver", 13}, {"Sandra", 91},
		{"Ann", 6}, {"Elizabeth", 9}, {"Kate", 20}}

	fmt.Printf("Array before sort: %v\n", workerSlice)
	fmt.Println("---------Sort by id-----------")
	sortedArray, _ := HeapSort(workerSlice, func(a, b Worker) bool {
		return a.GetID() > b.GetID() // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = HeapSort(workerSlice, func(a, b Worker) bool {
		return a.GetID() < b.GetID() // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
	fmt.Println("---------Sort by name-----------")
	sortedArray, _ = HeapSort(workerSlice, func(a, b Worker) bool {
		return a.GetName() > b.GetName() // по убыванию
	})
	fmt.Printf("Array after descending sorting: %v\n", sortedArray)
	sortedArray, _ = HeapSort(workerSlice, func(a, b Worker) bool {
		return a.GetName() < b.GetName() // по возрастанию
	})
	fmt.Printf("Array after ascending sorting: %v\n", sortedArray)
}
