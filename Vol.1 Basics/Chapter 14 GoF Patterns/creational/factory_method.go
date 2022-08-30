package main

import (
	"errors"
	"fmt"
)

const (
	Margarita = iota
	Stella
	Mexico
)

// ////////Base interface/////////////
type IPizza interface {
	GetName() string
	GetCost() int
	String() string // для вывода в терминал состояния структуры
}

// ///////Base struct/////////////
type Pizza struct {
	cost int
	name string
}

func (p *Pizza) GetName() string {
	return p.name
}

func (p *Pizza) GetCost() int {
	return p.cost
}

func (p *Pizza) String() string {
	return fmt.Sprintf("Pizza(name: %s, cost: %d)", p.name, p.cost)
}

// /////////Child Structs/////////////
// ////////Margarita///////////
type PizzaMargarita struct {
	Pizza
}

func NewMargarita() *PizzaMargarita {
	return &PizzaMargarita{
		Pizza: Pizza{
			name: "Margarita",
			cost: 150,
		},
	}
}

// ////////Stella///////////
type PizzaStella struct {
	Pizza
}

func NewStella() *PizzaStella {
	return &PizzaStella{
		Pizza: Pizza{
			name: "Stella",
			cost: 180,
		},
	}
}

// ////////Mexico///////////
type PizzaMexico struct {
	Pizza
}

func NewMexico() *PizzaMexico {
	return &PizzaMexico{
		Pizza: Pizza{
			name: "Mexico",
			cost: 231,
		},
	}
}

// /////Factory Method////////////
func PizzaFactory(pizzaType int) (IPizza, error) {
	switch pizzaType {
	case Margarita:
		return NewMargarita(), nil
	case Stella:
		return NewStella(), nil
	case Mexico:
		return NewMexico(), nil
	}
	return nil, errors.New("wrong pizza type")
}

func main() {
	for i := 0; i < 4; i++ {
		pizza, err := PizzaFactory(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(pizza)
		}
	}
}
