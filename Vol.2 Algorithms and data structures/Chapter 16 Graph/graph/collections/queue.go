package collections

import (
	"errors"
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

func (sl *singleLinkedList[T]) pushQueue(value T) {
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

func (sl *singleLinkedList[T]) popQueue() (T, error) {
	// Изъятие из начала списка
	if sl.isEmpty() {
		return *new(T), errors.New("list is empty")
	}
	popNode := sl.head
	sl.head = sl.head.nextPtr
	sl.length--
	return popNode.value, nil
}

func (sl *singleLinkedList[T]) pushStack(value T) {
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

func (sl *singleLinkedList[T]) popStack() (T, error) {
	if sl.isEmpty() {
		return *new(T), errors.New("list is empty")
	}
	popNode := sl.head
	sl.head = sl.head.nextPtr
	sl.length--
	return popNode.value, nil
}

// Queue

type Queue[T any] struct {
	list singleLinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	queue := &Queue[T]{
		list: *newSingleLinkedList[T](),
	}
	return queue
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.isEmpty()
}

func (q *Queue[T]) Size() int {
	return q.list.size()
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.pushQueue(value)
}

func (q *Queue[T]) Dequeue() (T, error) {
	popNode, err := q.list.popQueue()
	if err != nil {
		return popNode, errors.New("queue is empty")
	}
	return popNode, nil
}
