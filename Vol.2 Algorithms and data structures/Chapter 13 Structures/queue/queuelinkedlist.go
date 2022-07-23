package queue

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
	tail   *node[T]
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

func (sl *singleLinkedList[T]) push(value T) {
	// Добавление в начало списка
	pushNode := newNode[T](value)
	if sl.isEmpty() {
		// если список пуст
		sl.head = pushNode
		sl.tail = pushNode
		sl.length = 1
	} else {
		sl.tail.nextPtr = pushNode
		sl.tail = pushNode
		sl.length++
	}
}

func (sl *singleLinkedList[T]) pop() (T, error) {
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

// Queue

type QueueLinkedList[T any] struct {
	list singleLinkedList[T]
}

func NewQueueLinkedList[T any]() *QueueLinkedList[T] {
	queue := &QueueLinkedList[T]{
		list: *newSingleLinkedList[T](),
	}
	return queue
}

func (q *QueueLinkedList[T]) IsEmpty() bool {
	return q.list.isEmpty()
}

func (q *QueueLinkedList[T]) Size() int {
	return q.list.size()
}

func (q *QueueLinkedList[T]) Enqueue(value T) {
	q.list.push(value)
}

func (q *QueueLinkedList[T]) Dequeue() (T, error) {
	popNode, err := q.list.pop()
	if err != nil {
		return popNode, errors.New("queue is empty")
	}
	return popNode, nil
}

func (q *QueueLinkedList[T]) Peak() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("queue is empty")
	}
	return q.list.head.value, nil
}

func (q *QueueLinkedList[T]) PrintQueue() {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
		return
	}
	fmt.Print("Queue: [")
	q.list.printList()
	fmt.Println("]")
}

func main2() {
	queue := NewQueueLinkedList[Cat]()
	queue.Enqueue(Cat{"Max", 4})
	queue.Enqueue(Cat{"Alex", 5})
	queue.Enqueue(Cat{"Tom", 7})
	queue.Enqueue(Cat{"Tommy", 1})
	queue.PrintQueue()
	fmt.Println("----- Peak ------")
	peakValue, _ := queue.Peak()
	fmt.Printf("Queue head value: %+v\n", peakValue)
	queue.PrintQueue()
	fmt.Println("----- Dequeue ------")
	for !queue.IsEmpty() {
		popValue, err := queue.Dequeue()
		fmt.Printf("Queue head value: %+v, err: %v\n", popValue, err)
		queue.PrintQueue()
	}
}
