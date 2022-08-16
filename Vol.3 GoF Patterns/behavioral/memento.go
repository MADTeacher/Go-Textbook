package main

import "fmt"

// ////////////Memento  //////////////////
type Memento struct {
	states []string
}

func NewMemento(states []string) *Memento {
	return &Memento{
		states: states,
	}
}

func (m *Memento) GetState() []string {
	st := make([]string, len(m.states))
	copy(st, m.states)
	return st
}

// ////////////Pizza  //////////////////
type Pizza struct {
	state []string
}

func NewPizza() *Pizza {
	return &Pizza{
		state: []string{"base"},
	}
}

func (p *Pizza) CreateMemento() *Memento {
	st := make([]string, len(p.state))
	copy(st, p.state)
	return NewMemento(st)
}

func (p *Pizza) SetMemento(memento *Memento) {
	p.state = memento.GetState()
}

func (p *Pizza) AddIngredient(product string) {
	fmt.Printf("In pizza added ingredient: %v \n", product)
	p.state = append(p.state, product)
}

func (p *Pizza) String() string {
	return fmt.Sprintf("Current pizza state: %v \n", p.state)
}

// //////Chief //////////////////////
type Chief struct {
	pizza       *Pizza
	pizzaStates []*Memento
}

func NewChief(pizza *Pizza) *Chief {
	return &Chief{
		pizza:       pizza,
		pizzaStates: []*Memento{},
	}
}

func (c *Chief) AddIngredientToPizza(product string) {
	c.pizzaStates = append(c.pizzaStates, c.pizza.CreateMemento())
	c.pizza.AddIngredient(product)
}

func (c *Chief) UndoAddedIngredient() {
	if len(c.pizzaStates) <= 1 {
		c.pizza.SetMemento(c.pizzaStates[0])
		fmt.Println("Original pizza state")
		fmt.Printf("%v \n", c.pizza.String())
	} else {
		fmt.Println("Undoing previous action")
		state := c.pizzaStates[len(c.pizzaStates)-1]
		c.pizzaStates = c.pizzaStates[:len(c.pizzaStates)-1]
		c.pizza.SetMemento(state)
		fmt.Printf("%v \n", c.pizza.String())
	}
}

// ////////////////////////////////////
func main() {
	pizza := NewPizza()

	chief := NewChief(pizza)
	fmt.Print(pizza)
	fmt.Println("********Add ingredients********")
	chief.AddIngredientToPizza("sauce")
	chief.AddIngredientToPizza("olives")
	chief.AddIngredientToPizza("salami")
	chief.AddIngredientToPizza("cheese")
	fmt.Print(pizza)
	fmt.Println("********Undoing action********")
	chief.UndoAddedIngredient()
	chief.UndoAddedIngredient()
	chief.UndoAddedIngredient()
	chief.UndoAddedIngredient()
	fmt.Println("********Add ingredients again********")
	chief.AddIngredientToPizza("sauce")
	chief.AddIngredientToPizza("4 cheese")
	fmt.Print(pizza)
}
