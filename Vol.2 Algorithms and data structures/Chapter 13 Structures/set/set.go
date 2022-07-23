package set

import (
	"fmt"
)

type Cat struct {
	name string
	age  uint8
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vals ...T) *Set[T] {
	set := make(Set[T])
	for _, item := range vals {
		set.Add(item)
	}
	return &set
}

func (s *Set[T]) Add(value T) bool {
	oldLength := len(*s)
	(*s)[value] = struct{}{}
	return oldLength != len(*s)
}

func (s *Set[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Set[T]) Size() int {
	return len(*s)
}

func (s *Set[T]) RemoveAll() {
	*s = *NewSet[T]()
}

func (s *Set[T]) Contains(value T) bool {
	if _, ok := (*s)[value]; !ok {
		return false
	}
	return true
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	diff := NewSet[T]()
	for elem := range *s {
		if !other.Contains(elem) {
			diff.Add(elem)
		}
	}
	return diff
}

func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	symDiff := NewSet[T]()
	for it := range *s {
		if !other.Contains(it) {
			symDiff.Add(it)
		}
	}
	for it := range *other {
		if !s.Contains(it) {
			symDiff.Add(it)
		}
	}
	return symDiff
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	intersectionSet := NewSet[T]()
	if s.Size() < other.Size() {
		for it := range *s {
			if other.Contains(it) {
				intersectionSet.Add(it)
			}
		}
	} else {
		for it := range *other {
			if s.Contains(it) {
				intersectionSet.Add(it)
			}
		}
	}
	return intersectionSet
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if s.Size() > other.Size() {
		return false
	}
	for it := range *s {
		if !other.Contains(it) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	unionSet := NewSet[T]()

	for it := range *s {
		unionSet.Add(it)
	}
	for it := range *other {
		unionSet.Add(it)
	}
	return unionSet
}

func (s *Set[T]) ForEach(val func(value T)) {
	if s.Size() == 0 {
		return
	}
	for key, _ := range *s {
		val(key)
	}
}

func (s *Set[T]) PrintSet() {
	fmt.Print("Set: [")
	s.ForEach(func(value T) {
		fmt.Printf("%+v, ", value)
	})
	fmt.Println("]")
}

func main() {
	setA := NewSet(1, 2, 3, 4, 5, 6, 3, 3, 3, 2, 1, 4, 7, 8, 5)
	fmt.Println("----- setA ------")
	setA.PrintSet()

	setB := NewSet(10, 22, 1, 4, 21, 4, 5, 21, 11, 10)
	fmt.Println("----- setB ------")
	setB.PrintSet()

	fmt.Println("----- Contains ------")
	value := 10 // поменяйте на любой существующий ключ
	fmt.Printf("SetA is contains'%v' == %v\n",
		value, setA.Contains(value))
	fmt.Printf("SetB is contains '%v'== %v\n",
		value, setB.Contains(value))

	fmt.Println("----- A - B ------")
	setA.Difference(setB).PrintSet()
	fmt.Println("----- B - A ------")
	setB.Difference(setA).PrintSet()
	fmt.Println("----- Symmetric Difference ------")
	setA.SymmetricDifference(setB).PrintSet()
	fmt.Println("----- Union ------")
	setA.Union(setB).PrintSet()
	fmt.Println("----- Intersect ------")
	setA.Intersect(setB).PrintSet()
}
