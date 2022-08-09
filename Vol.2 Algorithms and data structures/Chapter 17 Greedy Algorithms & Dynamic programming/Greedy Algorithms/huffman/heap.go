package main

import "errors"

type IKey interface {
	GetKey() int
}

///////////// Heap////////////////
type Heap[T IKey] struct {
	currentSize int
	heapArray   []*T
	maxSize     int
}

func NewHeap[T IKey](maxSize int) *Heap[T] {
	return &Heap[T]{
		maxSize:     maxSize,
		currentSize: 0,
		heapArray:   make([]*T, maxSize),
	}
}

func (h *Heap[T]) IsEmpty() bool {
	return h.currentSize == 0
}

func (h *Heap[T]) Size() int {
	return h.currentSize
}

func (h *Heap[T]) trickleUp(index int) {
	var parent int = (index - 1) / 2
	bottom := h.heapArray[index]
	for index > 0 &&
		(*h.heapArray[parent]).GetKey() >
			(*bottom).GetKey() {
		h.heapArray[index] = h.heapArray[parent]
		index = parent
		parent = (parent - 1) / 2
	}
	h.heapArray[index] = bottom
}

func (h *Heap[T]) trickleDown(index int) {
	var largerChild int
	top := h.heapArray[index]     // Сохранение корня
	for index < h.currentSize/2 { // Пока у узла имеется хотя бы один потомок
		var leftChild int = 2*index + 1
		var rightChild int = leftChild + 1
		// Определение большего потомка
		if rightChild < h.currentSize && // (Правый потомок существует?)
			(*h.heapArray[leftChild]).GetKey() >
				(*h.heapArray[rightChild]).GetKey() {
			largerChild = rightChild
		} else {
			largerChild = leftChild
		}

		if (*top).GetKey() <= (*h.heapArray[largerChild]).GetKey() {
			break
		}

		// Потомок сдвигается вверх
		h.heapArray[index] = h.heapArray[largerChild]
		index = largerChild // Переход вниз
	}
	h.heapArray[index] = top // index <- корень
}

func (h *Heap[T]) Insert(newValue T) error {
	if h.currentSize == h.maxSize {
		return errors.New("heap is full")
	}
	h.heapArray[h.currentSize] = &newValue
	h.trickleUp(h.currentSize)
	h.currentSize++
	return nil
}

func (h *Heap[T]) Remove() (T, error) {
	if h.currentSize <= 0 {
		return *new(T), errors.New("heap is empty")
	}
	root := h.heapArray[0]
	h.currentSize--
	h.heapArray[0] = h.heapArray[h.currentSize]
	h.trickleDown(0)
	return *root, nil
}
