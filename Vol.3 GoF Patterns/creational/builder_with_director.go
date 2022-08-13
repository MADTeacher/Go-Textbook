package main

import (
	"fmt"
)

type PizzaDoughDepth string
type PizzaDoughType string
type PizzaSauceType string
type PizzaTopLevelType string

// Константы для конфигурации приготовляемой пиццы
const (
	// Толщина теста
	Thin  PizzaDoughDepth = "thin"
	Thick PizzaDoughDepth = "thick"
	// тип теста (пшеница, рожь, кукуруза)
	Wheat PizzaDoughType = "wheat"
	Rye   PizzaDoughType = "rye"
	Corn  PizzaDoughType = "corn"
	// Тип соуса
	Pesto       PizzaSauceType = "pesto"
	WhiteGarLic PizzaSauceType = "whiteGarLic"
	Barbeque    PizzaSauceType = "barbeque"
	Tomato      PizzaSauceType = "tomato"
	// Тип топпинга пиццы
	Mozzarella PizzaTopLevelType = "mozzarella"
	Salami     PizzaTopLevelType = "salami"
	Bacon      PizzaTopLevelType = "bacon"
	Mushrooms  PizzaTopLevelType = "mushrooms"
	Shrimps    PizzaTopLevelType = "shrimps"
)

// ////////Основа для пиццы/////////////
type PizzaBase struct {
	DoughDepth PizzaDoughDepth
	DoughType  PizzaDoughType
}

func (pb *PizzaBase) String() string {
	return fmt.Sprintf("dough type: %s:%s\n", pb.DoughType, pb.DoughDepth)
}

func NewPizzaBase(typeDough PizzaDoughType,
	depthDough PizzaDoughDepth) PizzaBase {
	return PizzaBase{
		DoughDepth: depthDough,
		DoughType:  typeDough,
	}
}

// ///////////Описываем пиццу/////////////////
type Pizza struct {
	Name        string
	Dough       PizzaBase
	Sauce       PizzaSauceType
	CookingTime int
	Topping     []PizzaTopLevelType
}

func (p *Pizza) String() string {
	infoStr := fmt.Sprintf("Pizza name: %s\n", p.Name)
	infoStr += p.Dough.String()
	infoStr += "sauce type: " + string(p.Sauce) + "\n"
	infoStr += "topping: {"
	for _, it := range p.Topping {
		infoStr += string(it) + " "
	}
	infoStr += "}\n"
	infoStr += fmt.Sprintf("cooking time: %d minutes", p.CookingTime)
	return infoStr
}

// ///////Builder interface//////////////////
type Builder interface {
	PrepareDough()
	AddSauce()
	AddTopping()
	GetPizza() Pizza
}

// /////////MargaritaPizzaBuilder//////////////////
type MargaritaPizzaBuilder struct {
	pizza Pizza
}

func (m *MargaritaPizzaBuilder) AddSauce() {
	m.pizza.Sauce = Tomato
}

func (m *MargaritaPizzaBuilder) AddTopping() {
	m.pizza.Topping = append(m.pizza.Topping,
		[]PizzaTopLevelType{Bacon, Mozzarella, Mozzarella}...)
}

func (m *MargaritaPizzaBuilder) PrepareDough() {
	m.pizza.Dough = NewPizzaBase(PizzaDoughType(Thick),
		PizzaDoughDepth(Wheat))
}

func (m *MargaritaPizzaBuilder) GetPizza() Pizza {
	return m.pizza
}

func NewMargaritaPizzaBuilder() *MargaritaPizzaBuilder {
	return &MargaritaPizzaBuilder{
		pizza: Pizza{
			Name:        "Margarita",
			CookingTime: 15,
		},
	}
}

// /////////SalamiPizzaBuilder//////////////////
type SalamiPizzaBuilder struct {
	pizza Pizza
}

func (m *SalamiPizzaBuilder) AddSauce() {
	m.pizza.Sauce = Barbeque
}

func (m *SalamiPizzaBuilder) AddTopping() {
	m.pizza.Topping = append(m.pizza.Topping,
		[]PizzaTopLevelType{Salami, Mozzarella}...)
}

func (m *SalamiPizzaBuilder) PrepareDough() {
	m.pizza.Dough = NewPizzaBase(PizzaDoughType(Thin),
		PizzaDoughDepth(Rye))
}

func (m *SalamiPizzaBuilder) GetPizza() Pizza {
	return m.pizza
}

func NewSalamiPizzaBuilder() *SalamiPizzaBuilder {
	return &SalamiPizzaBuilder{
		pizza: Pizza{
			Name:        "Salami",
			CookingTime: 10,
		},
	}
}

// /////////Director//////////////////
type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) MakePizza() {
	if d.builder == nil {
		panic("builder is nil")
	}
	d.builder.PrepareDough()
	d.builder.AddSauce()
	d.builder.AddTopping()
}

func main() {
	director := Director{}
	builders := []Builder{NewMargaritaPizzaBuilder(),
		NewSalamiPizzaBuilder()}
	for _, it := range builders {
		director.SetBuilder(it)
		director.MakePizza()
		pizza := it.GetPizza()
		fmt.Printf("%s\n", pizza.String())
		fmt.Println("-----------------------------------")
	}
}
