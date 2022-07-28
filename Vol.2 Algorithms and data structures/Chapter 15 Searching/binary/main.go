package main

import (
	"errors"
	"fmt"
	"sort"
)

////// Worker /////////
type Worker struct {
	Name string
	Id   uint8
}

func (c *Worker) GetID() int {
	return int(c.Id)
}

func (c *Worker) GetName() string {
	return c.Name
}

////////////поиск по id/////////////
func binarySearchByIDImpl(arr []Worker, targetId uint8, l, r int) (int, error) {
	if r < l || len(arr) == 0 {
		return -1, errors.New("id not found")
	}
	mid := l + (r-l)/2
	if uint8(arr[mid].GetID()) > targetId {
		return binarySearchByIDImpl(arr, targetId, l, mid-1)
	} else if uint8(arr[mid].GetID()) < targetId {
		return binarySearchByIDImpl(arr, targetId, mid+1, r)
	} else {
		return mid, nil
	}
}

func BinarySearchByID(arr []Worker, id uint8) (int, error) {
	if id < uint8(arr[0].GetID()) || id > uint8(arr[len(arr)-1].GetID()) {
		return -1, errors.New("id not found")
	}
	return binarySearchByIDImpl(arr, id, 0, len(arr)-1)
}

////////////поиск по имени/////////////
func binarySearchByNameImpl(arr []Worker, targetName string, l, r int) (int, error) {
	if r < l || len(arr) == 0 {
		return -1, errors.New("id not found")
	}
	mid := l + (r-l)/2
	if arr[mid].GetName() > targetName {
		return binarySearchByNameImpl(arr, targetName, l, mid-1)
	} else if arr[mid].GetName() < targetName {
		return binarySearchByNameImpl(arr, targetName, mid+1, r)
	} else {
		return mid, nil
	}
}

func BinarySearchByName(arr []Worker, name string) (int, error) {
	if name < arr[0].GetName() || name > arr[len(arr)-1].GetName() {
		return -1, errors.New("id not found")
	}
	return binarySearchByNameImpl(arr, name, 0, len(arr)-1)
}

func main() {
	workerSlice := []Worker{{"Julie", 1}, {"Alex", 2}, {"Tom", 4},
		{"George", 3}, {"Max", 60}, {"Tommy", 94}, {"William", 12},
		{"Sophia", 14}, {"Oliver", 13}, {"Sandra", 91},
		{"Ann", 6}, {"Elizabeth", 9}, {"Kate", 20}}
	// сортировка по возрастанию id
	sort.Slice(workerSlice, func(i, j int) bool {
		return workerSlice[i].GetID() < workerSlice[j].GetID()
	})
	fmt.Printf("Array after sorting by id: %v\n", workerSlice)
	index, _ := BinarySearchByID(workerSlice, 4) // поиск существующего id
	fmt.Printf("Element is located by the index: %v, his value:%v\n",
		index, workerSlice[index])
	id := 32
	index, err := BinarySearchByID(workerSlice, uint8(id)) // поиск не существующего id
	if err != nil {
		fmt.Printf("%v: %v\n", err, id)
	}
	// сортировка по возрастанию имени
	sort.Slice(workerSlice, func(i, j int) bool {
		return workerSlice[i].GetName() < workerSlice[j].GetName()
	})
	fmt.Printf("Array after sorting by Name: %v\n", workerSlice)
	index, _ = BinarySearchByName(workerSlice, "Kate") // поиск существующего имени
	fmt.Printf("Element is located by the index: %v, his value:%v\n",
		index, workerSlice[index])
}
