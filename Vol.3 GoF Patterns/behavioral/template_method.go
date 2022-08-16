package main

import "fmt"

/////////////interfaces///////////
type IPizzaMaker interface {
	PrepareSauce(pizza *Pizza)
	PrepareTopping(pizza *Pizza)
	Cooking(pizza *Pizza)
}

// ////////////Pizza   //////////////////
type Pizza struct {
	ingredients []string
}

func NewPizza() *Pizza {
	return &Pizza{[]string{"base"}}
}

func (p *Pizza) AddIngredient(ingredient string) {
	fmt.Printf(" %v added to pizza\n", ingredient)
	p.ingredients = append(p.ingredients, ingredient)
}

func (p *Pizza) String() string {
	return fmt.Sprintf("Ingredients of pizza:  %v", p.ingredients)
}

// ////////////PizzaMaker //////////////////
type PizzaMaker struct {
	IPizzaMaker
}

func NewPizzaMaker(maker IPizzaMaker) *PizzaMaker {
	return &PizzaMaker{
		IPizzaMaker: maker,
	}
}

func (p *PizzaMaker) MakePizza(pizza *Pizza) {
	p.PrepareSauce(pizza)
	p.PrepareTopping(pizza)
	p.Cooking(pizza)
}

// ////////////MargaritaMaker //////////////////
type MargaritaMaker struct {
	PizzaMaker
}

func (p *MargaritaMaker) PrepareSauce(pizza *Pizza) {
	pizza.AddIngredient("Tomato")
}

func (p *MargaritaMaker) PrepareTopping(pizza *Pizza) {
	pizza.AddIngredient("Bacon")
	pizza.AddIngredient("Mozzarella")
	pizza.AddIngredient("Mozzarella")
}

func (p *MargaritaMaker) Cooking(pizza *Pizza) {
	fmt.Println("Margarita will be ready in 10 minutes")
}

// ////////////SalamiMaker //////////////////
type SalamiMaker struct {
	PizzaMaker
}

func (p *SalamiMaker) PrepareSauce(pizza *Pizza) {
	pizza.AddIngredient("Pesto")
}

func (p *SalamiMaker) PrepareTopping(pizza *Pizza) {
	pizza.AddIngredient("Salami")
	pizza.AddIngredient("Salami")
	pizza.AddIngredient("Mozzarella")
}

func (p *SalamiMaker) Cooking(pizza *Pizza) {
	fmt.Println("Salami will be ready in 15 minutes")
}

// ////////////Chief  //////////////////
type Chief struct {
	templatePizza *PizzaMaker
}

func NewChief(template *PizzaMaker) *Chief {
	return &Chief{template}
}

func (p *Chief) SetCookTemplate(template *PizzaMaker) {
	p.templatePizza = template
}

func (p *Chief) MakePizza() *Pizza {
	pizza := NewPizza()
	p.templatePizza.MakePizza(pizza)
	return pizza
}

// ////////////////////////////////////
func main() {
	chief := NewChief(NewPizzaMaker(&MargaritaMaker{}))
	fmt.Println("********Cooking pizza 'Margarita'********")
	fmt.Println(chief.MakePizza())
	fmt.Println("********Cooking pizza 'Salami'********")
	chief.SetCookTemplate(NewPizzaMaker(&SalamiMaker{}))
	fmt.Println(chief.MakePizza())
}
