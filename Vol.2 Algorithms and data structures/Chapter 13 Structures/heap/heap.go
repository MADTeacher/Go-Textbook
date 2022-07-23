package heap

import (
	"errors"
	"fmt"
)

type IKey interface {
	GetKey() int
}

///////////// Cat/////////////
type Cat struct {
	name string
	age  uint8
}

func (c *Cat) GetKey() int {
	return int(c.age)
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

func NewHeapFromSlice[T IKey](values []T) *Heap[T] {
	maxSize := len(values)
	heap := &Heap[T]{
		maxSize:     maxSize,
		currentSize: 0,
		heapArray:   make([]*T, maxSize),
	}
	for _, it := range values {
		heap.Insert(it)
	}
	return heap
}

func (h *Heap[T]) IsEmpty() bool {
	return h.currentSize == 0
}

func (h *Heap[T]) trickleUp(index int) {
	var parent int = (index - 1) / 2
	bottom := h.heapArray[index]
	for index > 0 &&
		(*h.heapArray[parent]).GetKey() <
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
			(*h.heapArray[leftChild]).GetKey() <
				(*h.heapArray[rightChild]).GetKey() {
			largerChild = rightChild
		} else {
			largerChild = leftChild
		}

		if (*top).GetKey() >= (*h.heapArray[largerChild]).GetKey() {
			break
		}

		// Потомок сдвигается вверх
		h.heapArray[index] = h.heapArray[largerChild]
		index = largerChild // Переход вниз
	}
	h.heapArray[index] = top // index <- корень
}

func (h *Heap[T]) change(index int, newValue T) error {
	if index < 0 || index >= h.currentSize {
		return errors.New("heap index out of bounds")
	}
	oldValue := h.heapArray[index]
	h.heapArray[index] = &newValue

	if (*oldValue).GetKey() < newValue.GetKey() { // Если узел повышается,
		h.trickleUp(index) // выполняется смещение вверх.
	} else { // Если узел понижается,
		h.trickleDown(index) // выполняется смещение вниз.
	}
	return nil
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

func (h *Heap[T]) PrintHeap() {
	fmt.Print("heapArray: ")
	for it := 0; it < h.currentSize; it++ {
		if h.heapArray[it] != nil {
			fmt.Printf("%v  ", (*h.heapArray[it]).GetKey())
		} else {
			fmt.Print("-- ")
		}
	}
	fmt.Println()

	var nBlanks, itemsPerRow, column, j int = 32, 1, 0, 0
	dots := "..............................."
	fmt.Println(dots + dots)
	for h.currentSize > 0 {
		if column == 0 {
			for it := 0; it < nBlanks; it++ {
				fmt.Print(" ")
			}
		}
		// fmt.Printf("%v", *h.heapArray[j])
		fmt.Printf("%v", (*h.heapArray[j]).GetKey())
		j++
		if j == h.currentSize {
			break
		}

		column++
		if column == itemsPerRow {
			// Конец строки
			nBlanks /= 2     // Половина пробелов
			itemsPerRow *= 2 // Вдвое больше элементов
			column = 0       // Начать заново
			fmt.Println()    // Переход на новую строку
		} else {
			for it := 0; it < nBlanks*2-2; it++ {
				fmt.Print(" ")
			}
		}

	}
	fmt.Println("\n" + dots + dots + "\n")
}

func main() {
	cats := []*Cat{
		{"Max", 4}, {"Alex", 5}, {"Tom", 7}, {"Tommy", 1},
		{"Max", 14}, {"Alex", 53}, {"Tom", 79}, {"Tommy", 11},
		{"Max", 43}, {"Alex", 5}, {"Tom", 17}, {"Tommy", 31},
	}
	heap := NewHeapFromSlice[*Cat](cats)
	///////////////////
	//или
	///////////////////
	// heap := NewHeap[*Cat](31)
	// heap.Insert(&Cat{"Max", 4})
	// heap.Insert(&Cat{"Alex", 5})
	// heap.Insert(&Cat{"Tom", 7})
	// heap.Insert(&Cat{"Tommy", 1})
	// heap.Insert(&Cat{"Max", 14})
	// heap.Insert(&Cat{"Alex", 53})
	// heap.Insert(&Cat{"Tom", 79})
	// heap.Insert(&Cat{"Tommy", 11})
	// heap.Insert(&Cat{"Max", 43})
	// heap.Insert(&Cat{"Alex", 5})
	// heap.Insert(&Cat{"Tom", 17})
	// heap.Insert(&Cat{"Tommy", 31})
	heap.PrintHeap()
	fmt.Println("----- Remove ------")
	root, err := heap.Remove()
	fmt.Printf("Remove root: %+v, err: %v\n", root, err)
	heap.PrintHeap()
	fmt.Println("----- Change ------")
	heap.change(3, &Cat{"Tommy", 99})
	heap.PrintHeap()
}
