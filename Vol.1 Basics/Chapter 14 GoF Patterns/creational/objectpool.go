package main

import (
	"errors"
	"fmt"
	"sync"
)

type Platter struct {
	isEmpty bool
	id      int
}

func (p *Platter) GetID() int {
	return p.id
}

func (p *Platter) IsEmpty() bool {
	return p.isEmpty
}

func (p *Platter) SetPizza() {
	p.isEmpty = false
}

func (p *Platter) GetPizza() {
	p.Reset()
}

func (p *Platter) Reset() {
	p.isEmpty = true
}

func (p *Platter) String() string {
	return fmt.Sprintf("Platter id: %d, empty: %v", p.id, p.isEmpty)
}

func newPlatter(id int) *Platter {
	return &Platter{
		id:      id,
		isEmpty: true,
	}
}

// ////////////ObjectPool///////////////
type ObjectPool struct {
	resources []*Platter
	capacity  int
	mulock    *sync.Mutex
}

func NewObjectPool(capacity int) ObjectPool {
	resources := []*Platter{}
	for it := 0; it < capacity; it++ {
		resources = append(resources, newPlatter(it))
	}
	return ObjectPool{
		resources: resources,
		capacity:  capacity,
		mulock:    &sync.Mutex{},
	}
}

func (o *ObjectPool) GetObject() (*Platter, error) {
	o.mulock.Lock()
	defer o.mulock.Unlock()
	if o.capacity > 0 {
		platter := o.resources[0]
		o.resources = o.resources[1:]
		return platter, nil
	} else {
		return nil, errors.New("object pool is emty")
	}
}

func (o *ObjectPool) ReleaseObject(platter *Platter) {
	o.mulock.Lock()
	defer o.mulock.Unlock()
	platter.Reset()
	o.resources = append(o.resources, platter)
}

func (o *ObjectPool) FreeAmountObjects() int {
	o.mulock.Lock()
	defer o.mulock.Unlock()
	return len(o.resources)
}

func (o *ObjectPool) String() string {
	o.mulock.Lock()
	defer o.mulock.Unlock()
	return fmt.Sprintf("Amount objects into pool: %v", len(o.resources))
}

func main() {
	objPool := NewObjectPool(3)
	fmt.Println(objPool.String())
	myPlatter, _ := objPool.GetObject()
	fmt.Println(myPlatter)
	fmt.Println(objPool.String())
	myPlatter.SetPizza()
	fmt.Println(myPlatter)
	objPool.ReleaseObject(myPlatter)
	fmt.Println(objPool.String())
}
