package linkedlist

import (
	"errors"
	"fmt"
)

type Cat struct {
	name string
	age  uint8
}

type cyclicNode[T any] struct {
	data    T
	nextPtr *cyclicNode[T]
	prevPtr *cyclicNode[T]
}

type CyclicLinkedList[T any] struct {
	length int
	head   *cyclicNode[T]
}

func newCyclicNode[T any](data T) *cyclicNode[T] {
	return &cyclicNode[T]{data, nil, nil}
}

func NewCyclicLinkedList[T any]() *CyclicLinkedList[T] {
	return &CyclicLinkedList[T]{0, nil}
}

func (cl *CyclicLinkedList[T]) Size() int {
	return cl.length
}

func (cl *CyclicLinkedList[T]) IsEmpty() bool {
	return cl.length == 0
}

func (cl *CyclicLinkedList[T]) Add(data T) {
	node := newCyclicNode(data)
	if cl.IsEmpty() {
		node.prevPtr = node
		node.nextPtr = node
		cl.head = node
	} else {
		node.prevPtr = cl.head.prevPtr
		node.nextPtr = cl.head
		cl.head.prevPtr.nextPtr = node
		cl.head.prevPtr = node
		cl.head = node
	}
	cl.length++
}

func (cl *CyclicLinkedList[T]) ForEach(val func(data T)) {
	node := cl.head
	val(node.data)
	for i := 0; i < cl.length-1; i++ {
		node = node.nextPtr
		val(node.data)
	}
}

func (cl *CyclicLinkedList[T]) ReverseForEach(val func(data T)) {
	node := cl.head
	val(node.data)
	for i := 0; i < cl.length-1; i++ {
		node = node.prevPtr
		val(node.data)
	}
}

func (cl *CyclicLinkedList[T]) Rotate(delta int) {
	// изменение элемента на которую указывает вершина
	// зависит от знака аргумента delta
	// не может выходить за длину списка
	if cl.length > 0 {
		if delta < 0 {
			scalingCoef := cl.length - 1 - delta/cl.length
			delta += scalingCoef * cl.length
		}
		delta %= cl.length

		if delta > cl.length/2 {
			delta = cl.length - delta
			for i := 0; i < delta; i++ {
				cl.head = cl.head.prevPtr
			}
		} else if delta == 0 {
			return
		} else {
			for i := 0; i < delta; i++ {
				cl.head = cl.head.nextPtr
			}

		}
	}
}

func (cl *CyclicLinkedList[T]) Remove() bool {
	if cl.IsEmpty() {
		return false
	}

	currentNode := cl.head
	nextNode := currentNode.nextPtr
	prevNode := currentNode.prevPtr

	if cl.length == 1 {
		cl.head = nil
	} else {
		cl.head = nextNode
		nextNode.prevPtr = prevNode
		prevNode.nextPtr = nextNode
	}
	cl.length--

	return true
}

func (cl *CyclicLinkedList[T]) RemoveAll() {
	for cl.Remove() {
	}
}

func (cl *CyclicLinkedList[T]) Value() (T, error) {
	if cl.IsEmpty() {
		return *new(T), errors.New("list is empty")
	}
	return cl.head.data, nil
}

func main3() {
	cLinkedList := NewCyclicLinkedList[Cat]()
	cLinkedList.Add(Cat{"Max", 4})
	cLinkedList.Add(Cat{"Alex", 5})
	cLinkedList.Add(Cat{"Tom", 7})
	cLinkedList.Add(Cat{"Tommy", 1})
	fmt.Printf("List size: %d\n", cLinkedList.Size())
	cLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
	fmt.Println("--Remove data--")
	cLinkedList.Remove()
	cLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
	fmt.Println("--Rotate--")
	data, _ := cLinkedList.Value()
	fmt.Printf("Head data before rotate: %+v\n", data)
	cLinkedList.Rotate(-1)
	data, _ = cLinkedList.Value()
	fmt.Printf("Head data after rotate: %+v\n", data)
	fmt.Println("--ReverseForEach--")
	cLinkedList.ReverseForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
}
