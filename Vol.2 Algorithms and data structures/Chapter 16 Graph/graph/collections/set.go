package collections

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

func (s *Set[T]) ForEach(val func(value *T)) {
	if s.Size() == 0 {
		return
	}
	for key, _ := range *s {
		val(&key)
	}
}
