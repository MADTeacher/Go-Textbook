package linkedlist

import (
	"errors"
	"fmt"
)

type doublyNode[T any] struct {
	data    T
	nextPtr *doublyNode[T]
	prevPtr *doublyNode[T]
}

type DoublyLinkedList[T any] struct {
	length int
	head   *doublyNode[T]
	tail   *doublyNode[T]
}

func newDoublyNode[T any](data T) *doublyNode[T] {
	return &doublyNode[T]{data, nil, nil}
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (dl *DoublyLinkedList[T]) Size() int {
	return dl.length
}

func (dl *DoublyLinkedList[T]) IsEmpty() bool {
	return dl.length == 0
}

func (dl *DoublyLinkedList[T]) PushTail(data T) error {
	// Добавление в конец списка
	node := newDoublyNode[T](data)
	if dl.length <= 0 {
		// если список пуст
		dl.head = node
		dl.tail = node
		dl.length = 1
		return nil
	}

	if dl.tail.nextPtr != nil {
		return errors.New("it is not tail")
	}
	if dl.tail == nil {
		return errors.New("tail cannot be nil")
	}
	dl.tail.nextPtr = node
	node.prevPtr = dl.tail
	dl.tail = node
	dl.length++
	return nil
}

func (dl *DoublyLinkedList[T]) PushHead(data T) error {
	// Добавление в начало списка
	node := newDoublyNode[T](data)
	if dl.length <= 0 {
		// если список пуст
		dl.head = node
		dl.tail = node
		dl.length = 1
		return nil
	}
	if dl.head.prevPtr != nil {
		return errors.New("it is not head")
	}
	if dl.head == nil {
		return errors.New("head cannot be nil")
	}
	node.nextPtr = dl.head
	dl.head.prevPtr = node
	dl.head = node
	dl.length++
	return nil
}

func (dl *DoublyLinkedList[T]) Insert(index int, data T) error {
	// вставка по индексу
	if index < 0 || index > dl.length-1 {
		return errors.New("index out of range, starting with zero")
	}
	if index == 0 {
		return dl.PushHead(data)
	}
	if index == dl.length-1 {
		return dl.PushTail(data)
	}

	node := dl.head
	for it := 0; it < index; it++ {
		node = node.nextPtr
	}

	insertNode := newDoublyNode[T](data)
	insertNode.nextPtr = node
	node.prevPtr.nextPtr = insertNode
	insertNode.prevPtr = node.prevPtr
	node.prevPtr = insertNode

	dl.length++
	return nil
}

func (dl *DoublyLinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index > dl.length-1 {
		return *new(T), errors.New("index out of range, starting with zero")
	}
	if index == 0 {
		return dl.head.data, nil
	}
	if index == dl.length-1 {
		return dl.tail.data, nil
	}
	node := dl.head
	for it := 0; it < index; it++ {
		node = node.nextPtr
	}
	return node.data, nil
}

func (dl *DoublyLinkedList[T]) Remove(index int) error {
	if index < 0 || index > dl.length-1 {
		return errors.New("index out of range, starting with zero")
	}
	if index == 0 {
		node := dl.head
		dl.head = node.nextPtr
		dl.head.prevPtr = nil
		node = nil
		dl.length--
		return nil
	}

	node := dl.head
	for it := 0; it < index-1; it++ {
		node = node.nextPtr
	}
	if index == dl.length-1 {
		dl.tail.prevPtr = nil
		dl.tail = node
		dl.tail.nextPtr = nil
		dl.length--
		return nil
	}

	deleteNode := node.nextPtr
	node.nextPtr = deleteNode.nextPtr
	node.nextPtr.prevPtr = deleteNode.prevPtr
	dl.length--
	return nil
}

func (dl *DoublyLinkedList[T]) ForEach(val func(data T)) {
	node := dl.head
	val(node.data)
	for node.nextPtr != nil {
		node = node.nextPtr
		val(node.data)
	}
}

func (dl *DoublyLinkedList[T]) ReverseForEach(val func(data T)) {
	node := dl.tail
	val(node.data)
	for node.prevPtr != nil {
		node = node.prevPtr
		val(node.data)
	}
}

func main2() {
	dLinkedList := NewDoublyLinkedList[Cat]()
	dLinkedList.PushHead(Cat{"Max", 4})
	dLinkedList.PushHead(Cat{"Alex", 5})
	dLinkedList.PushTail(Cat{"Tom", 7})
	dLinkedList.Insert(2, Cat{"Tommy", 1})
	fmt.Printf("List size: %d\n", dLinkedList.Size())
	dLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
	data, _ := dLinkedList.Get(1)
	fmt.Printf("Get data with index: %d, %+v\n", 1, data)

	fmt.Println("--Remove data--")
	dLinkedList.Remove(3)
	dLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})

	fmt.Println("--ReverseForEach--")
	dLinkedList.ReverseForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
}
