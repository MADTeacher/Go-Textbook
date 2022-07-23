package dynamicarray

import (
	"errors"
	"fmt"
)

type Cat struct {
	name string
	age  uint8
}

type DynamicArray[T any] struct {
	length   int
	capacity int
	arr      []T
}

func NewDynamicArray[T any](capacity int) *DynamicArray[T] {
	if capacity <= 0 {
		panic("Array capacity cannot be <= 0")
	}

	return &DynamicArray[T]{
		capacity: capacity,
		arr:      make([]T, capacity),
	}
}

func (da *DynamicArray[T]) GetLenght() int {
	return da.length
}

func (da *DynamicArray[T]) GetCapacity() int {
	return da.capacity
}

// обновление значения элемента массива
func (da *DynamicArray[T]) Put(index int, element T) error {
	err := da.checkRangeFromIndex(index)
	if err != nil {
		return err
	}

	da.arr[index] = element
	fmt.Printf("Current state: %+v\n", *da)
	return nil
}

// добавление нового элемента в массив
func (da *DynamicArray[T]) Add(element T) {
	if da.length == da.capacity {
		da.newCapacity()
	}

	da.arr[da.length] = element
	da.length++
	fmt.Printf("Current state: %+v\n", *da)
}

// Удаление элемента по индексу
func (da *DynamicArray[T]) Remove(index int) error {
	err := da.checkRangeFromIndex(index)
	if err != nil {
		return err
	}

	copy(da.arr[index:], da.arr[index+1:da.length])
	da.arr[da.length-1] = *new(T)
	da.length--
	fmt.Printf("Current state: %+v\n", *da)
	return nil
}

// Получение значения элемента по индексу
func (da *DynamicArray[T]) Get(index int) (T, error) {
	err := da.checkRangeFromIndex(index)
	if err != nil {
		return *new(T), err
	}
	return da.arr[index], nil
}

func (da *DynamicArray[T]) IsEmpty() bool {
	return da.length == 0
}

// Получение среза значений массива
func (da *DynamicArray[T]) GetData() []T {
	return da.arr[:da.length]
}

// Провека на правильность обращения к элементу по индексу
func (da *DynamicArray[T]) checkRangeFromIndex(index int) error {
	if index >= da.length || index < 0 {
		return errors.New(fmt.Sprintf("Index %d out of range %d",
			index, da.length))
	}
	return nil
}

// Выделение увеличение размера памяти, выделяемого под массив
func (da *DynamicArray[T]) newCapacity() {
	da.capacity = da.capacity * 2 // или << 1

	newArr := make([]T, da.capacity)
	copy(newArr, da.arr)
	da.arr = newArr
	fmt.Printf("New capacity = %d\n", da.capacity)
}

func main() {
	dynamicArray := NewDynamicArray[Cat](1)
	fmt.Printf("Start state: %+v\n", *dynamicArray)
	dynamicArray.Add(Cat{"Max", 4})
	dynamicArray.Add(Cat{"Alex", 5})
	dynamicArray.Add(Cat{"Tom", 7})
	dynamicArray.Remove(0)
	dynamicArray.Put(1, Cat{"Tommy", 1})
}
