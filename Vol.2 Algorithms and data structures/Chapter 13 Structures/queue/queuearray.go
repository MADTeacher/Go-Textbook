package queue

import (
	"errors"
	"fmt"
)

type Cat struct {
	name string
	age  uint8
}

type QueueArray[T any] struct {
	fixedSize uint
	data      []T
}

func NewQueueArray[T any](fixedSize uint) *QueueArray[T] {
	return &QueueArray[T]{
		fixedSize: fixedSize,
		data:      []T{},
	}
}

func (q *QueueArray[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *QueueArray[T]) Size() int {
	return len(q.data)
}

func (q *QueueArray[T]) Enqueue(value T) error {
	if uint(q.Size()) >= q.fixedSize {
		return errors.New("queue is full")
	}
	q.data = append(q.data, value)
	return nil
}

func (q *QueueArray[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("queue is empty")
	}
	data := q.data[0]
	q.data = q.data[1:]
	return data, nil
}

func (q *QueueArray[T]) Peak() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("queue is empty")
	}
	return q.data[0], nil
}

func (q *QueueArray[T]) PrintQueue() {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
		return
	}
	fmt.Print("Queue: [")
	for _, it := range q.data {
		fmt.Printf("%+v, ", it)
	}
	fmt.Println("]")
}

func main1() {
	var fSize uint = 3
	queue := NewQueueArray[Cat](fSize)
	queue.Enqueue(Cat{"Max", 4})
	queue.Enqueue(Cat{"Alex", 5})
	queue.Enqueue(Cat{"Tom", 7})
	queue.PrintQueue()
	err := queue.Enqueue(Cat{"Tommy", 1})
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("----- Peak ------")
	peakValue, _ := queue.Peak()
	fmt.Printf("Queue head value: %+v\n", peakValue)
	queue.PrintQueue()
	fmt.Println("----- Dequeue ------")
	for it := 0; it < int(fSize)+1; it++ {
		popValue, err := queue.Dequeue()
		fmt.Printf("Queue head value: %+v, err: %v\n", popValue, err)
		queue.PrintQueue()
	}
}
