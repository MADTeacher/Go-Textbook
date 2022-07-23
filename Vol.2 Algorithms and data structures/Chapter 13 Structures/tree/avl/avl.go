package avl

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
	height    uint8
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

///////////// AVLTree ////////////////
type AVLTree[T IKey] struct {
	root *node[T]
}

func NewAVLTree[T IKey]() *AVLTree[T] {
	return &AVLTree[T]{}
}

func (t *AVLTree[T]) IsEmpty() bool {
	return t.root == nil
}

func (t *AVLTree[T]) height(p *node[T]) uint8 {
	if p == nil {
		return 0
	}
	return p.height
}

func (t *AVLTree[T]) bfactor(p *node[T]) int {
	return int(t.height(p.rightNode)) - int(t.height(p.leftNode))
}

func (t *AVLTree[T]) updateHeight(p *node[T]) {
	hl := t.height(p.leftNode)
	hr := t.height(p.rightNode)
	if hl > hr {
		p.height = hl + 1
	} else {
		p.height = hr + 1
	}
}

func (t *AVLTree[T]) rigthRotate(p *node[T]) *node[T] {
	q := p.leftNode
	p.leftNode = q.rightNode
	q.rightNode = p
	t.updateHeight(p)
	t.updateHeight(q)
	return q
}

func (t *AVLTree[T]) leftRotate(q *node[T]) *node[T] {
	p := q.rightNode
	q.rightNode = p.leftNode
	p.leftNode = q
	t.updateHeight(q)
	t.updateHeight(p)
	return p
}

func (t *AVLTree[T]) balance(p *node[T]) *node[T] {
	t.updateHeight(p)
	if t.bfactor(p) >= 2 {
		if t.bfactor(p.rightNode) < 0 {
			p.rightNode = t.rigthRotate(p.rightNode)
		}
		return t.leftRotate(p)
	}
	if t.bfactor(p) <= -2 {
		if t.bfactor(p.leftNode) > 0 {
			p.leftNode = t.leftRotate(p.leftNode)
		}
		return t.rigthRotate(p)
	}
	return p
}

func (t *AVLTree[T]) insert(p *node[T], newValue *T) *node[T] {
	if p == nil {
		return &node[T]{data: *newValue, height: 1}
	}
	if (*newValue).GetKey() < p.GetKey() {
		p.leftNode = t.insert(p.leftNode, newValue)
	} else {
		p.rightNode = t.insert(p.rightNode, newValue)
	}
	return t.balance(p)
}

func (t *AVLTree[T]) Insert(newValue T) {
	t.root = t.insert(t.root, &newValue)
}

func (t *AVLTree[T]) findMinimumRootNode(node *node[T]) *node[T] {
	currentNode := node.leftNode
	if currentNode == nil {
		return node
	}
	for currentNode.leftNode != nil {
		currentNode = currentNode.leftNode
	}
	return currentNode
}

func (t *AVLTree[T]) removeNodeWithMinKey(node *node[T]) *node[T] {
	if node.leftNode == nil {
		return node.rightNode
	}
	node.leftNode = t.removeNodeWithMinKey(node.leftNode)
	return t.balance(node)
}

func (t *AVLTree[T]) remove(node *node[T], key int) *node[T] {
	if node == nil {
		return nil
	}
	if key < node.GetKey() {
		node.leftNode = t.remove(node.leftNode, key)
	} else if key > node.GetKey() {
		node.rightNode = t.remove(node.rightNode, key)
	} else { // найден узел с необходимым значением ключа
		q := node.leftNode
		r := node.rightNode
		node = nil
		if r == nil {
			return q
		}
		minNode := t.findMinimumRootNode(r)
		minNode.rightNode = t.removeNodeWithMinKey(r)
		minNode.leftNode = q
		return t.balance(minNode)
	}
	return t.balance(node)
}

func (t *AVLTree[T]) Find(key int) (T, error) {
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

func (t *AVLTree[T]) Remove(key int) error {
	_, err := t.Find(key)
	if err != nil {
		return err
	}
	t.root = t.remove(t.root, key)
	return nil
}

// симетричный обход дерева//////////////
func (t *AVLTree[T]) SymmetricTraversal(myFunc func(data T)) {
	fmt.Print("Symmetric traversal: ")
	t.symmetricTraversal(t.root, myFunc)
	fmt.Println()
}

func (t *AVLTree[T]) symmetricTraversal(localRoot *node[T], myFunc func(data T)) {
	if localRoot != nil {
		t.symmetricTraversal(localRoot.leftNode, myFunc)
		myFunc(localRoot.data)
		t.symmetricTraversal(localRoot.rightNode, myFunc)
	}
}

///////////////////////////////////////
func (t *AVLTree[T]) Minimum() (T, error) {
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

func (t *AVLTree[T]) Maximum() (T, error) {
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

func (t *AVLTree[T]) PrintTree() {
	str := "AVLTree\n"
	if !t.IsEmpty() {
		t.createStringTree(&str, "", t.root, true)
	}
	fmt.Println(str)
}

func (t *AVLTree[T]) createStringTree(str *string, strPrefix string,
	node *node[T], isTail bool) {
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
	tree := NewAVLTree[*Worker]()
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
}
