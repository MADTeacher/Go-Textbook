package main

import "fmt"

// ////////////interface////////////////////
type ICommand interface {
	Execute()
}

// ////////////ChiefAssistant  //////////////////
type ChiefAssistant struct{}

func (c *ChiefAssistant) PreparePizzaDough() {
	fmt.Println("ChiefAssistant prepares the pizza dough")
}

func (c *ChiefAssistant) PrepareTopping() {
	fmt.Println("ChiefAssistant cuts pizza toppings")
}

func (c *ChiefAssistant) PrepareSauce() {
	fmt.Println("ChiefAssistant prepares sauce")
}

// ///////Oven /////////////////
type Oven struct{}

func (o *Oven) PrepareOven() {
	fmt.Println("Oven is heating up")
}

func (o *Oven) CookingPizza() {
	fmt.Println("Pizza is cooked in oven")
}

// /////////KitchenHandler ////////////
type Chief struct{}

func (c *Chief) MakePizzaBase() {
	fmt.Println("Chief rolls out a pizza base")
}

func (c *Chief) AppliedSauce() {
	fmt.Println("Chief applies sauce to the base of pizza")
}

func (c *Chief) AddToppingToPizza() {
	fmt.Println("Chief adds toppings to pizza")
}

func (c *Chief) BonAppetit() {
	fmt.Println("Chief wishes the client a bon appetit!")
}

//////////Oven Commands///////////////
type PrepareOvenCommand struct {
	executor *Oven
}

func (c *PrepareOvenCommand) Execute() {
	c.executor.PrepareOven()
}

type CookingPizzaCommand struct {
	executor *Oven
}

func (c *CookingPizzaCommand) Execute() {
	c.executor.CookingPizza()
}

//////////Chief Commands///////////////
type MakePizzaBaseCommand struct {
	executor *Chief
}

func (m *MakePizzaBaseCommand) Execute() {
	m.executor.MakePizzaBase()
}

type AppliedSauceCommand struct {
	executor *Chief
}

func (m *AppliedSauceCommand) Execute() {
	m.executor.AppliedSauce()
}

type AddToppingCommand struct {
	executor *Chief
}

func (m *AddToppingCommand) Execute() {
	m.executor.AddToppingToPizza()
}

type BonAppetitCommand struct {
	executor *Chief
}

func (m *BonAppetitCommand) Execute() {
	m.executor.BonAppetit()
}

//////////ChiefAssistant Commands///////////////
type PrepareDoughCommand struct {
	executor *ChiefAssistant
}

func (m *PrepareDoughCommand) Execute() {
	m.executor.PreparePizzaDough()
}

type PrepareToppingCommand struct {
	executor *ChiefAssistant
}

func (m *PrepareToppingCommand) Execute() {
	m.executor.PrepareTopping()
}

type PrepareSauceCommand struct {
	executor *ChiefAssistant
}

func (m *PrepareSauceCommand) Execute() {
	m.executor.PrepareSauce()
}

/////////////Pizzeria////////////////////////

type Pizzeria struct {
	history []ICommand
}

func NewPizzeria() *Pizzeria {
	return &Pizzeria{
		history: []ICommand{},
	}
}

func (p *Pizzeria) AddCommand(command ICommand) {
	p.history = append(p.history, command)
}

func (p *Pizzeria) Cooking() {
	if len(p.history) > 0 {
		for _, it := range p.history {
			it.Execute()
		}
	} else {
		fmt.Println("Hystory buffer is clear")
	}
	p.history = []ICommand{} // clear history
}

func main() {
	chief := &Chief{}
	assistant := &ChiefAssistant{}
	oven := &Oven{}

	pizzeria := Pizzeria{}

	pizzeria.AddCommand(&PrepareDoughCommand{assistant})
	pizzeria.AddCommand(&MakePizzaBaseCommand{chief})
	pizzeria.AddCommand(&PrepareSauceCommand{assistant})
	pizzeria.AddCommand(&AppliedSauceCommand{chief})
	pizzeria.AddCommand(&PrepareOvenCommand{oven})
	pizzeria.AddCommand(&PrepareToppingCommand{assistant})
	pizzeria.AddCommand(&AddToppingCommand{chief})
	pizzeria.AddCommand(&CookingPizzaCommand{oven})
	pizzeria.AddCommand(&BonAppetitCommand{chief})
	pizzeria.Cooking()
}
