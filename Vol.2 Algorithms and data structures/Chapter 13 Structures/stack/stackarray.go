package stack

import (
	"errors"
	"fmt"
)

type Cat struct {
	name string
	age  uint8
}

type StackArray[T any] struct {
	fixedSize uint
	data      []T
}

func NewStackArray[T any](fixedSize uint) *StackArray[T] {
	stackArray := &StackArray[T]{
		fixedSize: fixedSize,
	}
	return stackArray
}

func (s *StackArray[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *StackArray[T]) Size() int {
	return len(s.data)
}

func (s *StackArray[T]) Push(value T) error {
	if uint(s.Size()) >= s.fixedSize {
		return errors.New("stack overflow")
	}
	s.data = append([]T{value}, s.data...)
	return nil
}

func (s *StackArray[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	pop := s.data[0]
	s.data = s.data[1:]
	return pop, nil
}

func (s *StackArray[T]) Peak() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	return s.data[0], nil
}

func (s *StackArray[T]) PrintStack() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}
	fmt.Print("Stack: [")
	for _, it := range s.data {
		fmt.Printf("%+v, ", it)
	}
	fmt.Println("]")
}

func main() {
	var fSize uint = 3
	stack := NewStackArray[Cat](fSize)
	stack.Push(Cat{"Max", 4})
	stack.Push(Cat{"Alex", 5})
	stack.Push(Cat{"Tom", 7})
	stack.PrintStack()
	err := stack.Push(Cat{"Tommy", 1})
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("----- Peak ------")
	peakValue, _ := stack.Peak()
	fmt.Printf("Stack head value: %+v\n", peakValue)
	stack.PrintStack()
	fmt.Println("----- Pop ------")
	for it := 0; it < int(fSize)+1; it++ {
		popValue, err := stack.Pop()
		fmt.Printf("Stack head value: %+v, err: %v\n", popValue, err)
		stack.PrintStack()
	}
}
