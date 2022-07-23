package stack

import (
	"errors"
	"fmt"
)

// SingleLinkedList
type node[T any] struct {
	value   T
	nextPtr *node[T]
}

type singleLinkedList[T any] struct {
	length int
	head   *node[T]
}

func newNode[T any](data T) *node[T] {
	return &node[T]{data, nil}
}

func newSingleLinkedList[T any]() *singleLinkedList[T] {
	return &singleLinkedList[T]{}
}

func (sl *singleLinkedList[T]) size() int {
	return sl.length
}

func (sl *singleLinkedList[T]) isEmpty() bool {
	return sl.length == 0
}

func (sl *singleLinkedList[T]) pushHead(value T) {
	// Добавление в начало списка
	pushNode := newNode[T](value)
	if sl.isEmpty() {
		// если список пуст
		sl.head = pushNode
		sl.length = 1
	} else {
		pushNode.nextPtr = sl.head
		sl.head = pushNode
		sl.length++
	}
}

func (sl *singleLinkedList[T]) popHead() (T, error) {
	// Добавление в начало списка
	if sl.isEmpty() {
		return *new(T), errors.New("list is empty")
	}
	popNode := sl.head
	sl.head = sl.head.nextPtr
	sl.length--
	return popNode.value, nil
}

func (sl *singleLinkedList[T]) printList() {
	for printNode := sl.head; printNode != nil; {
		fmt.Printf("%+v, ", printNode.value)
		printNode = printNode.nextPtr
	}
}

// Stack

type StackLinkedList[T any] struct {
	fixedSize uint
	list      singleLinkedList[T]
}

func NewStackLinkedList[T any](fixedSize uint) *StackLinkedList[T] {
	stackArray := &StackLinkedList[T]{
		fixedSize: fixedSize,
		list:      *newSingleLinkedList[T](),
	}
	return stackArray
}

func (s *StackLinkedList[T]) IsEmpty() bool {
	return s.list.isEmpty()
}

func (s *StackLinkedList[T]) Size() int {
	return s.list.size()
}

func (s *StackLinkedList[T]) Push(value T) error {
	if uint(s.Size()) >= s.fixedSize {
		return errors.New("stack overflow")
	}
	s.list.pushHead(value)
	return nil
}

func (s *StackLinkedList[T]) Pop() (T, error) {
	popNode, err := s.list.popHead()
	if err != nil {
		return popNode, errors.New("stack is empty")
	}
	return popNode, nil
}

func (s *StackLinkedList[T]) Peak() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	return s.list.head.value, nil
}

func (s *StackLinkedList[T]) PrintStack() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}
	fmt.Print("Stack: [")
	s.list.printList()
	fmt.Println("]")
}

func main1() {
	var fSize uint = 3
	stack := NewStackLinkedList[Cat](fSize)
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
