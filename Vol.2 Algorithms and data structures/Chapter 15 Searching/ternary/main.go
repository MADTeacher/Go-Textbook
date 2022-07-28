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
func ternarySearchByIDImpl(arr []Worker, targetId uint8, l, r int) (int, error) {
	if r > l {
		if r-l == 1 {
			if targetId == uint8(arr[r].GetID()) {
				return r, nil
			}
			if targetId == uint8(arr[l].GetID()) {
				return l, nil
			}
			return -1, errors.New("id not found")
		}

		step := int(math.Round(float64(r-l) / 3.0))
		m1 := l + step
		m2 := m1 + step
		if targetId == uint8(arr[m1].GetID()) {
			return m1, nil
		}
		if targetId == uint8(arr[m2].GetID()) {
			return m2, nil
		}

		if targetId < uint8(arr[m1].GetID()) {
			return ternarySearchByIDImpl(arr, targetId, l, m1)
		} else if targetId > uint8(arr[m1].GetID()) &&
			targetId < uint8(arr[m2].GetID()) {
			return ternarySearchByIDImpl(arr, targetId, m1, m2)
		} else {
			return ternarySearchByIDImpl(arr, targetId, m2, r)
		}
	}
	return -1, errors.New("id not found")
}

func TernarySearchByID(arr []Worker, id uint8) (int, error) {
	if id < uint8(arr[0].GetID()) || id > uint8(arr[len(arr)-1].GetID()) {
		return -1, errors.New("id not found")
	}
	return ternarySearchByIDImpl(arr, id, 0, len(arr)-1)
}

////////////поиск по имени/////////////
func ternarySearchByNameImpl(arr []Worker, targetName string, l, r int) (int, error) {
	if r > l {
		if r-l == 1 {
			if targetName == arr[r].GetName() {
				return r, nil
			}
			if targetName == arr[l].GetName() {
				return l, nil
			}
			return -1, errors.New("id not found")
		}

		step := int(math.Round(float64(r-l) / 3.0))
		m1 := l + step
		m2 := m1 + step
		if targetName == arr[m1].GetName() {
			return m1, nil
		}
		if targetName == arr[m2].GetName() {
			return m2, nil
		}

		if targetName < arr[m1].GetName() {
			return ternarySearchByNameImpl(arr, targetName, l, m1)
		} else if targetName > arr[m1].GetName() &&
			targetName < arr[m2].GetName() {
			return ternarySearchByNameImpl(arr, targetName, m1, m2)
		} else {
			return ternarySearchByNameImpl(arr, targetName, m2, r)
		}
	}
	return -1, errors.New("id not found")
}

func TernarySearchByName(arr []Worker, name string) (int, error) {
	if name < arr[0].GetName() || name > arr[len(arr)-1].GetName() {
		return -1, errors.New("id not found")
	}
	return ternarySearchByNameImpl(arr, name, 0, len(arr)-1)
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
	index, _ := TernarySearchByID(workerSlice, 4) // поиск существующего id
	fmt.Printf("Element is located by the index: %v, his value:%v\n",
		index, workerSlice[index])
	id := 32
	index, err := TernarySearchByID(workerSlice, uint8(id)) // поиск не существующего id
	if err != nil {
		fmt.Printf("%v: %v\n", err, id)
	}
	// сортировка по возрастанию имени
	sort.Slice(workerSlice, func(i, j int) bool {
		return workerSlice[i].GetName() < workerSlice[j].GetName()
	})
	fmt.Printf("Array after sorting by Name: %v\n", workerSlice)
	index, _ = TernarySearchByName(workerSlice, "Kate") // поиск существующего имени
	fmt.Printf("Element is located by the index: %v, his value:%v\n",
		index, workerSlice[index])
}
