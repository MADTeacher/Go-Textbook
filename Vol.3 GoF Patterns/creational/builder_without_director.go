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

func newPizza(pb *PizzaBuilder) Pizza {
	return Pizza{
		Name:        pb.name,
		Dough:       pb.dough,
		Sauce:       pb.sauce,
		CookingTime: pb.cookingTime,
		Topping:     pb.topping,
	}
}

// ///////Builder interface//////////////////
type PizzaBuilder struct {
	name        string
	dough       PizzaBase
	sauce       PizzaSauceType
	cookingTime int
	topping     []PizzaTopLevelType
}

func (pb *PizzaBuilder) SetName(name string) *PizzaBuilder {
	pb.name = name
	return pb
}

func (pb *PizzaBuilder) SetSauce(sauce PizzaSauceType) *PizzaBuilder {
	pb.sauce = sauce
	return pb
}

func (pb *PizzaBuilder) SetDough(dough PizzaBase) *PizzaBuilder {
	pb.dough = dough
	return pb
}

func (pb *PizzaBuilder) SetCookingTime(time int) *PizzaBuilder {
	pb.cookingTime = time
	return pb
}

func (pb *PizzaBuilder) SetTopping(topping []PizzaTopLevelType) *PizzaBuilder {
	pb.topping = topping
	return pb
}

func (pb *PizzaBuilder) Build() Pizza {
	return newPizza(pb)
}

func GetPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}

func main() {
	builder := GetPizzaBuilder()
	pizza := builder.SetName("Margarita").
		SetDough(NewPizzaBase(PizzaDoughType(Thick), PizzaDoughDepth(Wheat))).
		SetCookingTime(15).
		SetSauce(Tomato).
		SetTopping([]PizzaTopLevelType{Bacon, Mozzarella, Mozzarella}).
		Build()

	fmt.Printf("%s\n", pizza.String())
	fmt.Println("-----------------------------------")

	builder.SetName("Salami")
	builder.SetDough(NewPizzaBase(PizzaDoughType(Thick),
		PizzaDoughDepth(Wheat)))
	builder.SetCookingTime(10)
	builder.SetSauce(Barbeque)
	builder.SetTopping([]PizzaTopLevelType{Salami, Mozzarella})

	pizza = builder.Build()
	fmt.Printf("%s\n", pizza.String())
}
