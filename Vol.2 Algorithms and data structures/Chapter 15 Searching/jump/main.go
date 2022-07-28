package main

import (
	"errors"
	"fmt"
	"math"
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
func JumpSearchByID(arr []Worker, x uint8) (int, error) {
	if x < uint8(arr[0].GetID()) || x > uint8(arr[len(arr)-1].GetID()) {
		return -1, errors.New("id not found")
	}

	step, pos := 0, 0
	step = int(math.Floor(math.Sqrt(float64(len(arr)))))
	for uint8(arr[pos].GetID()) < x {
		if uint8(arr[step].GetID()) > x || step >= len(arr) {
			break
		} else {
			pos = step
			step += int(math.Floor(math.Sqrt(float64(len(arr)))))
		}
	}
	for uint8(arr[pos].GetID()) < x {
		pos++
	}
	if uint8(arr[pos].GetID()) == x {
		return pos, nil
	}
	return -1, errors.New("id not found")
}

////////////поиск по имени/////////////
func JumpSearchByName(arr []Worker, x string) (int, error) {
	if x < arr[0].GetName() || x > arr[len(arr)-1].GetName() {
		return -1, errors.New("id not found")
	}
	step, pos := 0, 0
	step = int(math.Floor(math.Sqrt(float64(len(arr)))))
	for arr[pos].GetName() < x {
		if arr[step].GetName() > x || step >= len(arr) {
			break
		} else {
			pos = step
			step += int(math.Floor(math.Sqrt(float64(len(arr)))))
		}
	}
	for arr[pos].GetName() < x {
		pos++
	}
	if arr[pos].GetName() == x {
		return pos, nil
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
	index, _ := JumpSearchByID(workerSlice, 94) // поиск существующего id
	fmt.Printf("Element is located by the index: %v, his value:%v\n",
		index, workerSlice[index])
	id := 32
	index, err := JumpSearchByID(workerSlice, uint8(id)) // поиск не существующего id
	if err != nil {
		fmt.Printf("%v: %v\n", err, id)
	}
	// сортировка по возрастанию имени
	sort.Slice(workerSlice, func(i, j int) bool {
		return workerSlice[i].GetName() < workerSlice[j].GetName()
	})
	fmt.Printf("Array after sorting by Name: %v\n", workerSlice)
	index, _ = JumpSearchByName(workerSlice, "Kate") // поиск существующего имени
	fmt.Printf("Element is located by the index: %v, his value:%v\n",
		index, workerSlice[index])
}
