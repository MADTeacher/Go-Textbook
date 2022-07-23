package binarysearchtree

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type IKey interface {
	GetKey() int
}

// Node
type node[T IKey] struct {
	data      T
	leftNode  *node[T]
	rightNode *node[T]
}

func (n *node[T]) GetKey() int {
	return n.data.GetKey()
}

///////////// Worker/////////////
type Worker struct {
	Name string
	Id   uint8
}

func (c *Worker) GetKey() int {
	return int(c.Id)
}

///////////// BinarySearchTree ////////////////
type BinarySearchTree[T IKey] struct {
	root *node[T]
}

func NewBinarySearchTree[T IKey]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (t *BinarySearchTree[T]) IsEmpty() bool {
	return t.root == nil
}

func (t *BinarySearchTree[T]) Insert(newValue T) {
	newNode := node[T]{data: newValue}
	if t.root == nil {
		t.root = &newNode
	} else {
		currentNode := t.root
		var parent *node[T]
		for {
			parent = currentNode
			if newNode.GetKey() < currentNode.GetKey() {
				// двигаемся влево
				currentNode = currentNode.leftNode
				if currentNode == nil {
					parent.leftNode = &newNode
					return
				}
			} else {
				// двигаемся вправо
				currentNode = currentNode.rightNode
				if currentNode == nil {
					parent.rightNode = &newNode
					return
				}
			}
		}
	}
}

func (t *BinarySearchTree[T]) Remove(key int) error {
	currentNode := t.root
	parentNode := t.root
	var isLeftNode bool = true

	for currentNode.GetKey() != key {
		parentNode = currentNode
		if key < currentNode.GetKey() {
			isLeftNode = true
			currentNode = currentNode.leftNode
		} else {
			isLeftNode = false
			currentNode = currentNode.rightNode
		}
		if currentNode == nil {
			// Узел не найден
			return errors.New("key not found")
		}
	}

	// Если узел не имеет потомков, он просто удаляетс
	if currentNode.leftNode == nil &&
		currentNode.rightNode == nil {
		if currentNode == t.root {
			t.root = nil
		} else if isLeftNode {
			parentNode.leftNode = nil
		} else {
			parentNode.rightNode = nil
		}
	} else if currentNode.rightNode == nil {
		// Если нет правого потомка, узел заменяется левым поддеревом
		if currentNode == t.root {
			t.root = currentNode.leftNode
		} else if isLeftNode {
			parentNode.leftNode = currentNode.leftNode
		} else {
			parentNode.rightNode = currentNode.leftNode
		}
	} else if currentNode.leftNode == nil {
		// Если нет левого потомка, узел заменяется правым поддеревом
		if currentNode == t.root {
			t.root = currentNode.rightNode
		} else if isLeftNode {
			parentNode.leftNode = currentNode.rightNode
		} else {
			parentNode.rightNode = currentNode.rightNode
		}
	} else {
		// Два потомка, узел заменяется преемником
		successor := t.getSuccessor(currentNode)

		// Родитель currentNode связывается с посредником
		if currentNode == t.root {
			t.root = successor
		} else if isLeftNode {
			parentNode.leftNode = successor
		} else {
			parentNode.rightNode = successor
		}
		successor.leftNode = currentNode.leftNode
	}
	return nil
}

// Метод возвращает узел со следующим значением после удаляемого
// сначала осуществляется переход к правому потомку, а затем
// отслеживается цепочка левых потомков данного узла
func (t *BinarySearchTree[T]) getSuccessor(delNode *node[T]) *node[T] {
	successorParent := delNode
	successor := delNode
	currentNode := delNode.rightNode
	// Переход к правому потомку
	for currentNode != nil { // Пока остаются левые потомки
		successorParent = successor
		successor = currentNode
		currentNode = currentNode.leftNode
	}

	if successor != delNode.rightNode {
		// Если преемник не является правым потомком создаются связи между узлами
		successorParent.leftNode = successor.rightNode
		successor.rightNode = delNode.rightNode
	}
	return successor
}

func (t *BinarySearchTree[T]) Find(key int) (T, error) {
	currentNode := t.root
	for currentNode.GetKey() != key {
		if key < currentNode.GetKey() {
			currentNode = currentNode.leftNode
		} else {
			currentNode = currentNode.rightNode
		}
		if currentNode == nil {
			return *new(T), errors.New("key not found")
		}
	}
	return currentNode.data, nil
}

// симетричный обход дерева//////////////

func (t *BinarySearchTree[T]) SymmetricTraversal(myFunc func(data T)) {
	fmt.Print("Symmetric traversal: ")
	t.symmetricTraversal(t.root, myFunc)
	fmt.Println()
}

func (t *BinarySearchTree[T]) symmetricTraversal(localRoot *node[T], myFunc func(data T)) {
	if localRoot != nil {
		t.symmetricTraversal(localRoot.leftNode, myFunc)
		myFunc(localRoot.data)
		t.symmetricTraversal(localRoot.rightNode, myFunc)
	}
}

///////////////////////////////////////

// Неупорядоченный обход
func (t *BinarySearchTree[T]) TraversalAftertProcessing(myFunc func(data T)) {
	fmt.Print("Traversal aftert processing: ")
	t.traversalAftertProcessing(t.root, myFunc)
	fmt.Println()
}

func (t *BinarySearchTree[T]) traversalAftertProcessing(localRoot *node[T], myFunc func(data T)) {
	if localRoot != nil {
		myFunc(localRoot.data)
		t.traversalAftertProcessing(localRoot.leftNode, myFunc)
		t.traversalAftertProcessing(localRoot.rightNode, myFunc)
	}
}

func (t *BinarySearchTree[T]) TraversalBeforeProcessing(myFunc func(data T)) {
	fmt.Print("Traversal before processing: ")
	t.traversalBeforeProcessing(t.root, myFunc)
	fmt.Println()
}

func (t *BinarySearchTree[T]) traversalBeforeProcessing(localRoot *node[T], myFunc func(data T)) {
	if localRoot != nil {
		t.traversalBeforeProcessing(localRoot.leftNode, myFunc)
		t.traversalBeforeProcessing(localRoot.rightNode, myFunc)
		myFunc(localRoot.data)
	}
}

///////////////////////////////////////

func (t *BinarySearchTree[T]) Minimum() (T, error) {
	if t.root == nil {
		return *new(T), errors.New("tree is empty")
	}
	var current, last *node[T]
	current = t.root
	for current != nil {
		last = current
		current = current.leftNode
	}
	return last.data, nil
}

func (t *BinarySearchTree[T]) Maximum() (T, error) {
	if t.root == nil {
		return *new(T), errors.New("tree is empty")
	}
	var current, last *node[T]
	current = t.root
	for current != nil {
		last = current
		current = current.rightNode
	}
	return last.data, nil
}

func (t *BinarySearchTree[T]) PrintTree() {
	str := "BinarySearchTree\n"
	if !t.IsEmpty() {
		t.createStringTree(&str, "", t.root, true)
	}
	fmt.Println(str)
}

func (t *BinarySearchTree[T]) createStringTree(str *string, strPrefix string, node *node[T], isTail bool) {
	if node.rightNode != nil {
		newStrPrefix := strPrefix
		if isTail {
			newStrPrefix += "│   "
		} else {
			newStrPrefix += "    "
		}
		t.createStringTree(str, newStrPrefix, node.rightNode, false)
	}
	*str += strPrefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += fmt.Sprintf("%v\n", node.GetKey())
	if node.leftNode != nil {
		newStrPrefix := strPrefix
		if isTail {
			newStrPrefix += "    "
		} else {
			newStrPrefix += "│   "
		}
		t.createStringTree(str, newStrPrefix, node.leftNode, true)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	workers := []*Worker{
		{"Max", 83}, {"Alex", 58}, {"Tom", 98}, {"Tommy", 62},
		{"Max", 70}, {"Alex", 34}, {"Tom", 22},
		{"Max", 60}, {"Alex", 99}, {"Tom", 91}, {"Tommy", 94},
		{"Tommy", 85},
	}
	tree := NewBinarySearchTree[*Worker]()
	for _, it := range workers {
		tree.Insert(it)
		fmt.Println("----------------------------")
		fmt.Printf("Added %+v\n", *it)
		tree.PrintTree()
	}
	fmt.Println("----- Find ------")
	key, err := tree.Find(62)
	fmt.Printf("Founded key value: %+v , err: %v\n", key, err)
	fmt.Println("----- Max key ------")
	key, err = tree.Maximum()
	fmt.Printf("Max key value: %+v , err: %v\n", key, err)
	fmt.Println("----- Min key ------")
	key, err = tree.Minimum()
	fmt.Printf("Min key value: %+v , err: %v\n", key, err)
	fmt.Println("----- Remove ------")
	removingKey := workers[rand.Intn(len(workers))]
	fmt.Printf("Remove node with key: %+v\n", *removingKey)
	tree.Remove(removingKey.GetKey())
	tree.PrintTree()
	// Варианты обхода дерева
	tree.SymmetricTraversal(func(data *Worker) {
		fmt.Printf("%v ", data.GetKey())
	})
	tree.TraversalAftertProcessing(func(data *Worker) {
		fmt.Printf("%v ", data.GetKey())
	})
	tree.TraversalBeforeProcessing(func(data *Worker) {
		fmt.Printf("%v ", data.GetKey())
	})
}
