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

func getMin(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
func FibonaccianSearch(arr []Worker, x uint8) (int, error) {
	if x < uint8(arr[0].GetID()) || x > uint8(arr[len(arr)-1].GetID()) {
		return -1, errors.New("id not found")
	}
	fm2 := 0
	fm1 := 1
	fm := fm1 + fm2
	offset := -1

	for fm < len(arr) {
		fm2 = fm1
		fm1 = fm
		fm = fm1 + fm2
	}

	for fm > 1 {
		i := getMin(offset+fm2, len(arr)-1)
		if uint8(arr[i].GetID()) < x {
			fm = fm1
			fm1 = fm2
			fm2 = fm - fm1
			offset = i
		} else if uint8(arr[i].GetID()) > x {
			fm = fm2
			fm1 = fm1 - fm2
			fm2 = fm - fm1
		} else {
			return i, nil
		}
	}
	if fm1 == 1 {
		if uint8(arr[offset+1].GetID()) == x {
			return offset + 1, nil
		}
	}

	return -1, errors.New("id not found")
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
	index, _ := FibonaccianSearch(workerSlice, 13) // поиск существующего id
	fmt.Printf("Element is located by the index: %v, his value:%v\n",
		index, workerSlice[index])
	id := 32
	index, err := FibonaccianSearch(workerSlice, uint8(id)) // поиск не существующего id
	if err != nil {
		fmt.Printf("%v: %v\n", err, id)
	}
}
