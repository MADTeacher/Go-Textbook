package hashmap

import (
	"errors"
	"fmt"
	"hash/fnv"
	"log"
)

type Cat struct {
	name string
	age  uint8
}

const defaultCapacity uint64 = 1024

type node[K comparable, V any] struct {
	key     K
	value   V
	nextPtr *node[K, V]
}

type HashTable[K comparable, V any] struct {
	capacity uint64
	size     uint64
	table    []*node[K, V]
}

func (hm *HashTable[K, V]) newNode(key K, value V) *node[K, V] {
	return &node[K, V]{key, value, nil}
}

func (hm *HashTable[K, V]) resolvePutCollision(key K, value V, index uint64) {
	// добавление в конец списка
	// currentNode := hm.table[index]
	// for currentNode.nextPtr != nil {
	// 	currentNode = currentNode.nextPtr
	// }
	// currentNode.nextPtr = hm.newNode(key, value)

	// добавление в начало списка
	// currentNode := hm.newNode(key, value)
	// currentNode.nextPtr = hm.table[index]
	// hm.table[index] = currentNode
	// или
	hm.table[index] = &node[K, V]{
		key:     key,
		value:   value,
		nextPtr: hm.table[index],
	}
	fmt.Printf("Collision: {index: %v, key: %v, value:%+v}\n",
		index, key, value)
}

func NewHashTable[K comparable, V any]() *HashTable[K, V] {
	return &HashTable[K, V]{
		capacity: defaultCapacity,
		size:     0,
		table:    make([]*node[K, V], defaultCapacity),
	}
}

func NewHashTableWithCapacity[K comparable, V any](capacity uint64) (*HashTable[K, V], error) {
	if capacity == 0 {
		return new(HashTable[K, V]), errors.New("capacity cannot be zero")
	}

	return &HashTable[K, V]{
		capacity: capacity,
		size:     0,
		table:    make([]*node[K, V], capacity),
	}, nil
}

func (hm *HashTable[K, V]) Capacity() uint64 {
	return hm.capacity
}

func (hm *HashTable[K, V]) Put(key K, value V) {
	index := hm.hash(key)
	if hm.table[index] == nil {
		hm.table[index] = hm.newNode(key, value)
		hm.size++
	} else {
		for it := hm.table[index]; it != nil; {
			if it.key == key {
				it.value = value
				return
			}
			it = it.nextPtr
		}
		hm.resolvePutCollision(key, value, index)
		hm.size++
	}
}

func (hm *HashTable[K, V]) Get(key K) (V, error) {
	index := hm.hash(key)
	if hm.table[index] == nil {
		return *new(V), errors.New("key not found")
	}

	for currentNode := hm.table[index]; currentNode != nil; {
		if currentNode.key == key {
			return currentNode.value, nil
		}
		currentNode = currentNode.nextPtr

	}
	return *new(V), errors.New("key not found")

}

func (hm *HashTable[K, V]) Contains(key K) bool {
	index := hm.hash(key)
	if hm.table[index] == nil {
		return false
	}
	for it := hm.table[index]; it != nil; {
		if it.key == key {
			return true
		}
		it = it.nextPtr
	}
	return false
}

func (hm *HashTable[K, V]) Remove(key K) error {
	index := hm.hash(key)
	if hm.table[index] == nil {
		return errors.New("key not found")
	}

	currentNode := hm.table[index]
	if currentNode.key == key {
		hm.table[index] = currentNode.nextPtr
		currentNode = nil
		hm.size--
		return nil
	}

	var prevNode, delNode *node[K, V]
	for it := hm.table[index]; it.nextPtr != nil; {
		if it.nextPtr.key == key {
			prevNode = it
			delNode = it.nextPtr
		}
		it = it.nextPtr
	}
	if prevNode == nil || delNode == nil {
		return errors.New("key not found")
	}
	prevNode.nextPtr = delNode.nextPtr
	delNode = nil
	hm.size--
	return nil
}

func (hm *HashTable[K, V]) Keys() []K {
	keys := []K{}
	for _, it := range hm.table {
		if it != nil {
			for currentNode := it; currentNode != nil; {
				keys = append(keys, currentNode.key)
				currentNode = currentNode.nextPtr
			}
		}
	}
	return keys
}

func (hm *HashTable[K, V]) Values() []V {
	values := []V{}
	for _, it := range hm.table {
		if it != nil {
			for currentNode := it; currentNode != nil; {
				values = append(values, currentNode.value)
				currentNode = currentNode.nextPtr
			}
		}
	}
	return values
}

func (hm *HashTable[K, V]) ForEach(val func(key K, value V)) {
	if hm.size == 0 {
		return
	}
	for _, it := range hm.table {
		if it != nil {
			for currentNode := it; currentNode != nil; {
				val(currentNode.key, currentNode.value)
				currentNode = currentNode.nextPtr
			}
		}
	}
}

func (hm *HashTable[K, V]) Clear() {
	hm.size = 0
	for _, it := range hm.table {
		if it != nil {
			it = nil
		}
	}
}

func (hm *HashTable[K, V]) hash(key K) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))
	hashValue := h.Sum64()
	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

func main() {
	hashTable, _ := NewHashTableWithCapacity[string, Cat](3)
	hashTable.Put("Home", Cat{"Alex", 4})
	hashTable.Put("1", Cat{"Tom", 6})
	hashTable.Put("2", Cat{"Tommy", 6})
	hashTable.Put("2", Cat{"Max", 1}) // перезапись значения по ключу
	hashTable.Put("3", Cat{"Alex", 4})
	hashTable.Put("4", Cat{"Tom", 6})
	hashTable.Put("5", Cat{"Tommy", 6})
	hashTable.Put("Work", Cat{"Max", 1})
	hashTable.ForEach(func(key string, value Cat) {
		fmt.Printf("%v:%+v\n", key, value)
	})
	fmt.Println("-----Get Value------")
	val, err := hashTable.Get("1") // поменяйте на 0
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%+v\n", val)
	}
	fmt.Println("----- Remove ------")
	hashTable.Remove("Work")
	hashTable.Remove("5")
	hashTable.ForEach(func(key string, value Cat) {
		fmt.Printf("%v:%+v\n", key, value)
	})
	fmt.Println("----- Keys ------")
	fmt.Print("[")
	for _, it := range hashTable.Keys() {
		fmt.Printf("%+v  ", it)
	}
	fmt.Println("]")
	fmt.Println("----- Values ------")
	fmt.Print("[")
	for _, it := range hashTable.Values() {
		fmt.Printf("%+v  ", it)
	}
	fmt.Println("]")
	fmt.Println("----- Contains ------")
	key := "Home" // поменяйте на любой существующий ключ
	fmt.Printf("Key '%v' is contains == %v\n",
		key, hashTable.Contains(key))
	fmt.Println("----- Remove all ------")
	hashTable.Clear()
	hashTable.ForEach(func(key string, value Cat) {
		fmt.Printf("%v:%+v\n", key, value)
	})
}
