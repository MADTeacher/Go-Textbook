package main

import "fmt"

// /////////interfaces/////////////////
type IPizzaBase interface {
	Cost() int
}

type IDecorator interface {
	IPizzaBase
	Name() string
}

// ///////PizzaBase /////////////////
type PizzaBase struct {
	cost int
}

func (p *PizzaBase) Cost() int {
	return p.cost
}

// ///////PizzaMargarita /////////////////
type PizzaMargarita struct {
	wrapper IPizzaBase
	cost    int
}

func (p *PizzaMargarita) Cost() int {
	return p.cost + p.wrapper.Cost()
}

func (p *PizzaMargarita) Name() string {
	return "Margarita"
}

// ///////PizzaSalami /////////////////
type PizzaSalami struct {
	wrapper IPizzaBase
	cost    int
}

func (p *PizzaSalami) Cost() int {
	return (p.cost + p.wrapper.Cost()) * 2
}

func (p *PizzaSalami) Name() string {
	return "Salami"
}

func main() {
	pizzaBase := &PizzaBase{3}
	fmt.Printf("Pizza base cost = %v\n", pizzaBase.Cost())
	margarita := PizzaMargarita{pizzaBase, 10}
	fmt.Printf("Pizza cost '%v' = %v\n", margarita.Name(), margarita.Cost())
	salami := PizzaSalami{pizzaBase, 10}
	fmt.Printf("Pizza cost '%v' = %v\n", salami.Name(), salami.Cost())
}
