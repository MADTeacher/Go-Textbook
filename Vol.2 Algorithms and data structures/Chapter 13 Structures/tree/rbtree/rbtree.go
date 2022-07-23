package rbtree

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type color bool

const (
	red   color = false
	black color = true
)

type IKey interface {
	GetKey() int
}

// Node
type node[T IKey] struct {
	data      T
	parent    *node[T]
	leftNode  *node[T]
	rightNode *node[T]
	color     color
}

func (n *node[T]) GetKey() int {
	return n.data.GetKey()
}

func (n *node[T]) grandfather() *node[T] {
	if n != nil && n.parent != nil {
		return n.parent.parent
	}
	return nil
}

func (n *node[T]) uncle() *node[T] {
	if n == nil || n.parent == nil || n.parent.parent == nil {
		return nil
	}
	return n.parent.brother()
}

func (n *node[T]) brother() *node[T] {
	if n == nil || n.parent == nil {
		return nil
	}
	if n == n.parent.leftNode {
		return n.parent.rightNode
	}
	return n.parent.leftNode
}

///////////// Worker/////////////
type Worker struct {
	Name string
	Id   uint8
}

func (c *Worker) GetKey() int {
	return int(c.Id)
}

///////////// RBTree ////////////////
type RBTree[T IKey] struct {
	root *node[T]
}

func NewRBTree[T IKey]() *RBTree[T] {
	return &RBTree[T]{}
}

func (t *RBTree[T]) IsEmpty() bool {
	return t.root == nil
}

func getNodeColor[T IKey](node *node[T]) color {
	if node == nil {
		return black
	}
	return node.color
}

func (tree *RBTree[T]) replaceNode(a *node[T], b *node[T]) {
	if a.parent == nil {
		tree.root = b
	} else {
		if a == a.parent.leftNode {
			a.parent.leftNode = b
		} else {
			a.parent.rightNode = b
		}
	}
	if b != nil {
		b.parent = a.parent
	}
}

func (tree *RBTree[T]) rotateLeft(node *node[T]) {
	right := node.rightNode
	tree.replaceNode(node, right)
	node.rightNode = right.leftNode
	if right.leftNode != nil {
		right.leftNode.parent = node
	}
	right.leftNode = node
	node.parent = right
}

func (tree *RBTree[T]) rotateRight(node *node[T]) {
	left := node.leftNode
	tree.replaceNode(node, left)
	node.leftNode = left.rightNode
	if left.rightNode != nil {
		left.rightNode.parent = node
	}
	left.rightNode = node
	node.parent = left
}

func (t *RBTree[T]) Insert(newValue T) {
	var newNode *node[T]
	if t.root == nil {
		t.root = &node[T]{data: newValue, color: red}
		newNode = t.root
	} else {
		currentNode := t.root
		insertNode := &node[T]{data: newValue, color: red}
		for {
			if insertNode.GetKey() == currentNode.GetKey() {
				currentNode.data = newValue
				return
			}
			if insertNode.GetKey() < currentNode.GetKey() {
				if currentNode.leftNode == nil {
					currentNode.leftNode = insertNode
					newNode = currentNode.leftNode
					break
				} else {
					currentNode = currentNode.leftNode
				}
			}
			if insertNode.GetKey() > currentNode.GetKey() {
				if currentNode.rightNode == nil {
					currentNode.rightNode = insertNode
					newNode = currentNode.rightNode
					break
				} else {
					currentNode = currentNode.rightNode
				}
			}
		}
		newNode.parent = currentNode
	}
	t.insertRBTreeCase1(newNode)
}

func (t *RBTree[T]) insertRBTreeCase1(node *node[T]) {
	if node.parent == nil {
		node.color = black
	} else {
		t.insertRBTreeCase2(node)
	}
}

func (t *RBTree[T]) insertRBTreeCase2(node *node[T]) {
	if getNodeColor(node.parent) == black {
		return
	}
	t.insertRBTreeCase3(node)
}

func (t *RBTree[T]) insertRBTreeCase3(node *node[T]) {
	uncle := node.uncle()
	if getNodeColor(uncle) == red {
		node.parent.color = black
		uncle.color = black
		node.grandfather().color = red
		t.insertRBTreeCase1(node.grandfather())
	} else {
		t.insertRBTreeCase4(node)
	}
}

func (t *RBTree[T]) insertRBTreeCase4(node *node[T]) {
	grandfather := node.grandfather()
	if node == node.parent.rightNode &&
		node.parent == grandfather.leftNode {
		t.rotateLeft(node.parent)
		node = node.leftNode
	} else if node == node.parent.leftNode &&
		node.parent == grandfather.rightNode {
		t.rotateRight(node.parent)
		node = node.rightNode
	}
	t.insertRBTreeCase5(node)
}

func (t *RBTree[T]) insertRBTreeCase5(node *node[T]) {
	node.parent.color = black
	grandfather := node.grandfather()
	grandfather.color = red
	if node == node.parent.leftNode &&
		node.parent == grandfather.leftNode {
		t.rotateRight(grandfather)
	} else if node == node.parent.rightNode &&
		node.parent == grandfather.rightNode {
		t.rotateLeft(grandfather)
	}
}

func (t *RBTree[T]) findNode(key int) (*node[T], error) {
	currentNode := t.root
	for currentNode.GetKey() != key {
		if key < currentNode.GetKey() {
			currentNode = currentNode.leftNode
		} else {
			currentNode = currentNode.rightNode
		}
		if currentNode == nil {
			return nil, errors.New("key not found")
		}
	}
	return currentNode, nil
}

func (t *RBTree[T]) findLeftMaximumNode(node *node[T]) (*node[T], error) {
	currentNode := node.leftNode
	if currentNode == nil {
		return nil, errors.New("node is empty")
	}
	for currentNode.rightNode != nil {
		currentNode = currentNode.rightNode
	}
	return currentNode, nil
}

func (t *RBTree[T]) Remove(key int) error {
	var childNode *node[T]
	removingNode, err := t.findNode(key)
	if err != nil {
		return err
	}
	if removingNode.leftNode != nil &&
		removingNode.rightNode != nil {
		successor, err := t.findLeftMaximumNode(removingNode)
		if err != nil {
			return err
		}
		removingNode.data = successor.data
		removingNode = successor
	}
	if removingNode.leftNode == nil ||
		removingNode.rightNode == nil {
		if removingNode.rightNode == nil {
			childNode = removingNode.leftNode
		} else {
			childNode = removingNode.rightNode
		}
		if removingNode.color == black {
			removingNode.color = getNodeColor(childNode)
			t.deleteRBTreeCase1(removingNode)
		}
		t.replaceNode(removingNode, childNode)
		if removingNode.parent == nil && childNode != nil {
			childNode.color = black
		}
	}
	return nil
}

func (t *RBTree[T]) deleteRBTreeCase1(node *node[T]) {
	if node.parent == nil {
		return
	}
	t.deleteRBTreeCase2(node)
}

func (t *RBTree[T]) deleteRBTreeCase2(node *node[T]) {
	brother := node.brother()
	if getNodeColor(brother) == red {
		node.parent.color = red
		brother.color = black
		if node == node.parent.leftNode {
			t.rotateLeft(node.parent)
		} else {
			t.rotateRight(node.parent)
		}
	}
	t.deleteRBTreeCase3(node)
}

func (t *RBTree[T]) deleteRBTreeCase3(node *node[T]) {
	brother := node.brother()
	if getNodeColor(node.parent) == black &&
		getNodeColor(brother) == black &&
		getNodeColor(brother.leftNode) == black &&
		getNodeColor(brother.rightNode) == black {
		brother.color = red
		t.deleteRBTreeCase1(node.parent)
	} else {
		t.deleteRBTreeCase4(node)
	}
}

func (t *RBTree[T]) deleteRBTreeCase4(node *node[T]) {
	brother := node.brother()
	if getNodeColor(node.parent) == red &&
		getNodeColor(brother) == black &&
		getNodeColor(brother.leftNode) == black &&
		getNodeColor(brother.rightNode) == black {
		brother.color = red
		node.parent.color = black
	} else {
		t.deleteRBTreeCase5(node)
	}
}

func (t *RBTree[T]) deleteRBTreeCase5(node *node[T]) {
	brother := node.brother()
	if node == node.parent.leftNode &&
		getNodeColor(brother) == black &&
		getNodeColor(brother.leftNode) == red &&
		getNodeColor(brother.rightNode) == black {
		brother.color = red
		brother.leftNode.color = black
		t.rotateRight(brother)
	} else if node == node.parent.rightNode &&
		getNodeColor(brother) == black &&
		getNodeColor(brother.rightNode) == red &&
		getNodeColor(brother.leftNode) == black {
		brother.color = red
		brother.rightNode.color = black
		t.rotateLeft(brother)
	}
	t.deleteRBTreeCase6(node)
}

func (t *RBTree[T]) deleteRBTreeCase6(node *node[T]) {
	brother := node.brother()
	brother.color = getNodeColor(node.parent)
	node.parent.color = black
	if node == node.parent.leftNode &&
		getNodeColor(brother.rightNode) == red {
		brother.rightNode.color = black
		t.rotateLeft(node.parent)
	} else if getNodeColor(brother.leftNode) == red {
		brother.leftNode.color = black
		t.rotateRight(node.parent)
	}
}

func (t *RBTree[T]) Find(key int) (T, error) {
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

// симметричный обход дерева//////////////

func (t *RBTree[T]) SymmetricTraversal(myFunc func(data T)) {
	fmt.Print("Symmetric traversal: ")
	t.symmetricTraversal(t.root, myFunc)
	fmt.Println()
}

func (t *RBTree[T]) symmetricTraversal(localRoot *node[T], myFunc func(data T)) {
	if localRoot != nil {
		t.symmetricTraversal(localRoot.leftNode, myFunc)
		myFunc(localRoot.data)
		t.symmetricTraversal(localRoot.rightNode, myFunc)
	}
}

/////////////////////////////////////

func (t *RBTree[T]) Minimum() (T, error) {
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

func (t *RBTree[T]) Maximum() (T, error) {
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

func (t *RBTree[T]) PrintRBTree() {
	str := "RBTree\n"
	if !t.IsEmpty() {
		t.createStringTree(&str, "", t.root, true)
	}
	fmt.Println(str)
}

func (t *RBTree[T]) createStringTree(str *string, strPrefix string, node *node[T], isTail bool) {
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
	tree := NewRBTree[*Worker]()
	for _, it := range workers {
		tree.Insert(it)
		fmt.Println("----------------------------")
		fmt.Printf("Added %+v\n", *it)
		tree.PrintRBTree()
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
	tree.PrintRBTree()
	// Варианты обхода дерева
	tree.SymmetricTraversal(func(data *Worker) {
		fmt.Printf("%v ", data.GetKey())
	})
}
