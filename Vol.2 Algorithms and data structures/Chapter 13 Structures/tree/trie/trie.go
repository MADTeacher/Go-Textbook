package trie

import (
	"errors"
	"fmt"
)

type Worker struct {
	Name string
	Id   uint8
}

type node[T any] struct {
	label      byte
	leftNode   *node[T]
	rightNode  *node[T]
	middleNode *node[T]
	value      *T
	isEnd      bool
	key        string
}

type Trie[T any] struct {
	root *node[T]
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{}
}

func (t *Trie[V]) Contains(key string) bool {
	if len(key) == 0 {
		return false
	}
	_, ok := t.Get(key)
	return ok
}

func (t *Trie[T]) Get(key string) (*T, bool) {
	if len(key) == 0 {
		return nil, false
	}
	x := t.findNode(t.root, key, 0)
	if x == nil || !x.isEnd {
		return nil, false
	}
	return x.value, true
}

func (t *Trie[T]) findNode(curNode *node[T],
	key string, index int) *node[T] {
	if curNode == nil || len(key) == 0 {
		return nil
	}
	c := key[index]
	if c < curNode.label {
		return t.findNode(curNode.leftNode, key, index)
	} else if c > curNode.label {
		return t.findNode(curNode.rightNode, key, index)
	} else if index < len(key)-1 {
		return t.findNode(curNode.middleNode, key, index+1)
	} else {
		return curNode
	}
}

func (t *Trie[T]) Put(key string, value T) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	t.root = t.put(t.root, key, &value, 0, true)
	return nil
}

func (t *Trie[T]) put(curNode *node[T], key string,
	value *T, index int, isEnd bool) *node[T] {
	label := key[index]
	if curNode == nil {
		curNode = &node[T]{
			label: label,
		}
	}
	if label < curNode.label {
		curNode.leftNode = t.put(curNode.leftNode, key,
			value, index, isEnd)
	} else if label > curNode.label {
		curNode.rightNode = t.put(curNode.rightNode, key,
			value, index, isEnd)
	} else if index < len(key)-1 {
		curNode.middleNode = t.put(curNode.middleNode, key,
			value, index+1, isEnd)
	} else {
		curNode.value = value
		curNode.isEnd = isEnd
		curNode.key = key
	}
	return curNode
}

func (t *Trie[T]) Remove(key string) {
	if len(key) == 0 || !t.Contains(key) {
		return
	}
	t.root = t.put(t.root, key, nil, 0, false)
}

// LongestPrefix returns the key that is the longest prefix of 'query'.
func (t *Trie[T]) LongestPrefix(query string) string {
	if len(query) == 0 {
		return ""
	}
	length := 0
	currentNode := t.root
	i := 0
	for currentNode != nil && i < len(query) {
		label := query[i]
		if label < currentNode.label {
			currentNode = currentNode.leftNode
		} else if label > currentNode.label {
			currentNode = currentNode.rightNode
		} else {
			i++
			if currentNode.isEnd {
				length = i
			}
			currentNode = currentNode.middleNode
		}
	}
	return query[:length]
}

func (t *Trie[T]) AllKeys() (queue []string) {
	return t.assemble(t.root, "", queue)
}

func (t *Trie[T]) AllKeysWithPrefix(prefix string) (queue []string, err error) {
	if len(prefix) == 0 {
		return nil, errors.New("query string is empty")
	}
	x := t.findNode(t.root, prefix, 0)
	if x == nil {
		return nil, errors.New("keys not founded")
	}
	if x.isEnd {
		queue = []string{prefix}
	}
	return t.assemble(x.middleNode, prefix, queue), nil
}

func (t *Trie[T]) assemble(curNode *node[T], prefix string,
	queue []string) []string {
	if curNode == nil {
		return queue
	}
	queue = t.assemble(curNode.leftNode, prefix, queue)
	if curNode.isEnd {
		queue = append(queue, prefix+string(curNode.label))
	}
	queue = t.assemble(curNode.middleNode,
		prefix+string(curNode.label), queue)
	return t.assemble(curNode.rightNode, prefix, queue)
}

func (t *Trie[T]) ForEach(myFunc func(key string, value T)) {
	t.forEach(t.root, myFunc)
}

func (t *Trie[T]) forEach(localRoot *node[T],
	myFunc func(key string, value T)) {
	if localRoot != nil {
		if localRoot.isEnd {
			myFunc(localRoot.key, *localRoot.value)
		}
		t.forEach(localRoot.leftNode, myFunc)
		t.forEach(localRoot.middleNode, myFunc)
		t.forEach(localRoot.rightNode, myFunc)
	}
}

func main() {
	trie := NewTrie[[]Worker]()
	trie.Put("manager", []Worker{{"Julie", 1}, {"Alex", 2}, {"Tom", 4}})
	trie.Put("policeman", []Worker{{"George", 3}, {"Max", 60}})
	trie.Put("postman", []Worker{{"Tommy", 94}, {"William", 12}})
	trie.Put("mathematician", []Worker{{"Sophia", 14}, {"Oliver", 13}})
	trie.Put("postwoman", []Worker{{"Sandra", 91}, {"Ann", 6}})
	trie.Put("policewoman", []Worker{{"Elizabeth", 9}, {"Kate", 20}})

	trie.ForEach(func(key string, value []Worker) {
		fmt.Printf("Key: %s ; Value: %+v\n", key, value)
	})

	fmt.Printf("'man' is Contains? %v\n", trie.Contains("man"))
	fmt.Printf("'manager' is Contains? %v\n", trie.Contains("manager"))
	keys, _ := trie.AllKeysWithPrefix("po")
	fmt.Printf("Keys with prefix 'po' %v\n", keys)
	keys, _ = trie.AllKeysWithPrefix("police")
	fmt.Printf("Keys with prefix 'police' %v\n", keys)
	keys, _ = trie.AllKeysWithPrefix("m")
	fmt.Printf("Keys with prefix 'm' %v\n", keys)
	fmt.Printf("All keys: %v\n", trie.AllKeys())

	fmt.Println("----- Remove key 'manager'------")
	trie.Remove("manager")
	trie.ForEach(func(key string, value []Worker) {
		fmt.Printf("Key: %s ; Value: %+v\n", key, value)
	})
}
