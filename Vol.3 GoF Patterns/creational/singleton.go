package main

import (
	"fmt"
	"sync"
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

var Margarita *Pizza
var once sync.Once

func GetMargaritaInstance() *Pizza {
	if Margarita == nil {
		once.Do(func() { // код выпонится только 1 раз
			Margarita = &Pizza{
				Name:        "Margarita",
				Dough:       NewPizzaBase(PizzaDoughType(Thick), PizzaDoughDepth(Wheat)),
				Sauce:       Tomato,
				Topping:     []PizzaTopLevelType{Bacon, Mozzarella, Mozzarella},
				CookingTime: 15,
			}
		})
	}
	return Margarita
}

func main() {
	pizza := GetMargaritaInstance()
	fmt.Println(pizza)
	fmt.Println("-----------------------------------")
	pizza2 := GetMargaritaInstance()
	pizza2.Name = "New Mega Margarita"
	fmt.Println(pizza)
}
