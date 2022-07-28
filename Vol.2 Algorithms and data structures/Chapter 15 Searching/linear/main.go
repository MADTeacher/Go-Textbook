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

////////////////////////
func LinearSearchByID(arr []Worker, id uint8) (int, error) {
	for i, val := range arr {
		if id == uint8(val.GetID()) {
			return i, nil
		}
		if id < uint8(val.GetID()) {
			break
		}
	}

	return -1, errors.New("id not found")
}

func LinearSearchByName(arr []Worker, name string) (int, error) {
	for i, val := range arr {
		if name == val.GetName() {
			return i, nil
		}
		if name < val.GetName() {
			break
		}
	}

	return -1, errors.New("name not found")
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
	index, _ := LinearSearchByID(workerSlice, 4) // поиск существующего id
	fmt.Printf("Element is located by the index: %v, his value:%v\n",
		index, workerSlice[index])
	id := 32
	index, err := LinearSearchByID(workerSlice, uint8(id)) // поиск не существующего id
	if err != nil {
		fmt.Printf("%v: %v\n", err, id)
	}
	// сортировка по возрастанию имени
	sort.Slice(workerSlice, func(i, j int) bool {
		return workerSlice[i].GetName() < workerSlice[j].GetName()
	})
	fmt.Printf("Array after sorting by Name: %v\n", workerSlice)
	index, _ = LinearSearchByName(workerSlice, "Kate") // поиск существующего имени
	fmt.Printf("Element is located by the index: %v, his value:%v\n",
		index, workerSlice[index])
}
