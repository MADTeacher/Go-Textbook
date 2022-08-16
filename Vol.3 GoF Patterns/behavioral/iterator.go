package main

import "fmt"

// ////////////interface////////////////////
type Iterator interface {
	HasNext() bool
	Next() PizzaItem
}

// ////////////PizzaItem //////////////////
type PizzaItem struct {
	number int
}

func (p *PizzaItem) String() string {
	return fmt.Sprintln("Slice of pizza with number:", p.number)
}

// ///////PizzaSliceIterator /////////////////
type PizzaSliceIterator struct {
	pizza []PizzaItem
	index int
}

func newPizzaSliceIterator(pizza []PizzaItem) *PizzaSliceIterator {
	return &PizzaSliceIterator{
		pizza: pizza,
		index: 0,
	}
}

func (p *PizzaSliceIterator) HasNext() bool {
	return p.index < len(p.pizza)
}

func (p *PizzaSliceIterator) Next() PizzaItem {
	pizzaItem := p.pizza[p.index]
	p.index++
	return pizzaItem
}

// ///////PizzaAggregator  /////////////////
type PizzaAggregator struct {
	pizza []PizzaItem
}

func NewPizzaAggregator(amountSlices int) *PizzaAggregator {
	if amountSlices <= 0 {
		panic("pizza must have at least one slice")
	}
	pizzaSlices := make([]PizzaItem, amountSlices)
	for i := 1; i <= amountSlices; i++ {
		pizzaSlices[i-1] = PizzaItem{i}
	}

	return &PizzaAggregator{pizzaSlices}
}

func (p *PizzaAggregator) AmountSlices() int {
	return len(p.pizza)
}

func (p *PizzaAggregator) Iterator() Iterator {
	return newPizzaSliceIterator(p.pizza)
}

//////////////////////////////////
func main() {
	pizza := NewPizzaAggregator(15)
	iterator := pizza.Iterator()
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Print(item.String())
	}
}
