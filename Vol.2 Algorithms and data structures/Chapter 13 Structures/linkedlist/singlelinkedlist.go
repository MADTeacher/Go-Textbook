package linkedlist

import (
	"errors"
	"fmt"
)

type singleNode[T any] struct {
	data    T
	nextPtr *singleNode[T]
}

type SingleLinkedList[T any] struct {
	length int
	head   *singleNode[T]
	tail   *singleNode[T]
}

func newSingleNode[T any](data T) *singleNode[T] {
	return &singleNode[T]{data, nil}
}

func NewSingleLinkedList[T any]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}

func (sl *SingleLinkedList[T]) Size() int {
	return sl.length
}

func (sl *SingleLinkedList[T]) IsEmpty() bool {
	return sl.length == 0
}

func (sl *SingleLinkedList[T]) PushTail(data T) error {
	// Добавление в конец списка
	node := newSingleNode[T](data)
	if sl.length <= 0 {
		// если список пуст
		sl.head = node
		sl.tail = node
		sl.length = 1
		return nil
	}

	if sl.tail.nextPtr != nil {
		return errors.New("it is not tail")
	}
	if sl.tail == nil {
		return errors.New("tail cannot be nil")
	}
	sl.tail.nextPtr = node
	sl.tail = node
	sl.length++
	return nil
}

func (sl *SingleLinkedList[T]) PushHead(data T) error {
	// Добавление в начало списка
	node := newSingleNode[T](data)
	if sl.length <= 0 {
		// если список пуст
		sl.head = node
		sl.tail = node
		sl.length = 1
		return nil
	}
	if sl.head == nil {
		return errors.New("head cannot be nil")
	}
	node.nextPtr = sl.head
	sl.head = node
	sl.length++
	return nil
}

func (sl *SingleLinkedList[T]) Insert(index int, data T) error {
	// вставка по индексу
	if index < 0 || index > sl.length-1 {
		return errors.New("index out of range, starting with zero")
	}
	if index == 0 {
		return sl.PushHead(data)
	}
	if index == sl.length-1 {
		return sl.PushTail(data)
	}

	node := sl.head
	for it := 0; it < index-1; it++ {
		node = node.nextPtr
	}

	insertNode := newSingleNode[T](data)
	insertNode.nextPtr = node.nextPtr
	node.nextPtr = insertNode
	sl.length++
	return nil
}

func (sl *SingleLinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index > sl.length-1 {
		return *new(T), errors.New("index out of range, starting with zero")
	}
	if index == 0 {
		return sl.head.data, nil
	}
	if index == sl.length-1 {
		return sl.tail.data, nil
	}
	node := sl.head
	for it := 0; it < index; it++ {
		node = node.nextPtr
	}
	return node.data, nil
}

func (sl *SingleLinkedList[T]) Remove(index int) error {
	if index < 0 || index > sl.length-1 {
		return errors.New("index out of range, starting with zero")
	}
	if index == 0 {
		node := sl.head
		sl.head = node.nextPtr
		node = nil
		sl.length--
		return nil
	}

	node := sl.head
	for it := 0; it < index-1; it++ {
		node = node.nextPtr
	}
	if index == sl.length-1 {
		sl.tail = node
		node.nextPtr = nil
		sl.length--
		return nil
	}
	deleteNode := node.nextPtr
	node.nextPtr = deleteNode.nextPtr
	sl.length--
	return nil
}

func (sl *SingleLinkedList[T]) ForEach(val func(data T)) {
	node := sl.head
	val(node.data)
	for node.nextPtr != nil {
		node = node.nextPtr
		val(node.data)
	}
}

func main1() {
	sLinkedList := NewSingleLinkedList[Cat]()
	sLinkedList.PushHead(Cat{"Max", 4})
	sLinkedList.PushHead(Cat{"Alex", 5})
	sLinkedList.PushTail(Cat{"Tom", 7})
	sLinkedList.Insert(2, Cat{"Tommy", 1})
	fmt.Printf("List size: %d\n", sLinkedList.Size())
	sLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
	data, _ := sLinkedList.Get(1)
	fmt.Printf("Get data with index: %d, %+v\n", 1, data)

	fmt.Println("--Remove data--")
	sLinkedList.Remove(1)
	sLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
}
