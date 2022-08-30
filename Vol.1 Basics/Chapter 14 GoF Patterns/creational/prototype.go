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

func (p *Pizza) Clone() Pizza {
	pizza := Pizza{
		Name:        p.Name,
		Dough:       p.Dough,
		Sauce:       p.Sauce,
		CookingTime: p.CookingTime,
	}
	pizza.Topping = append(pizza.Topping, p.Topping...)
	return pizza
}

func main() {
	pizza := Pizza{
		Name:        "Margarita",
		Dough:       NewPizzaBase(PizzaDoughType(Thick), PizzaDoughDepth(Wheat)),
		Sauce:       Tomato,
		Topping:     []PizzaTopLevelType{Bacon, Mozzarella, Mozzarella},
		CookingTime: 15,
	}
	fmt.Printf("%s\n", pizza.String())
	fmt.Println("-----------------------------------")
	newPizza := pizza.Clone()
	newPizza.Topping = append(newPizza.Topping, Mushrooms)
	fmt.Printf("%s\n", newPizza.String())
	fmt.Println("-----------Old Pizza----------------")
	fmt.Printf("%s\n", pizza.String())
}
