package main

import (
	"errors"
	"fmt"
	"time"
)

// /////////interfaces/////////////////
type IOvenImplementor interface {
	WarmUp(temperature float64)
	CoolDown(temperature float64)
	CookPizza(pizza *Pizza)
	GetTemperature() float64
	GetOvenType() string
}

// ///////OriginalOven/////////////////
type Pizza struct {
	CookingTemperature float64
	Name               string
	CookingTime        int
	isCook             bool
}

func (p *Pizza) Cook() {
	p.isCook = true
}

func (p *Pizza) IsCooked() bool {
	return p.isCook
}

func NewPizza(name string, cookingTime int,
	cookingTemperature float64) *Pizza {
	return &Pizza{
		isCook:             false,
		CookingTemperature: cookingTemperature,
		Name:               name,
		CookingTime:        cookingTime,
	}
}

// ///////ClassicOvenImplementor /////////////////
type ClassicOvenImplementor struct {
	temperature float64
	ovenType    string
}

func NewClassicOvenImplementor(t float64) *ClassicOvenImplementor {
	if t < 0 {
		panic("This isn't a refrigerator")
	}
	return &ClassicOvenImplementor{
		temperature: t,
		ovenType:    "ClassicOven",
	}
}

func (ci *ClassicOvenImplementor) WarmUp(temperature float64) {
	if temperature < 0 {
		panic("This isn't a refrigerator")
	}
	time.Sleep(time.Duration(temperature-ci.temperature) * 30 * time.Millisecond)
	ci.temperature = temperature
}

func (ci *ClassicOvenImplementor) CoolDown(temperature float64) {
	if temperature < 0 {
		panic("This isn't a refrigerator")
	}
	time.Sleep(time.Duration(ci.temperature-temperature) * 25 * time.Millisecond)
	ci.temperature = temperature
}

func (ci *ClassicOvenImplementor) CookPizza(pizza *Pizza) {
	pizza.Cook()
}

func (ci *ClassicOvenImplementor) GetTemperature() float64 {
	return ci.temperature
}

func (ci *ClassicOvenImplementor) GetOvenType() string {
	return ci.ovenType
}

// ///////ClassicOvenImplementor /////////////////
type ElectricalOvenImplementor struct {
	temperature float64
	ovenType    string
}

func NewElectricalOvenImplementor(t float64) *ElectricalOvenImplementor {
	if t < 0 {
		panic("This isn't a refrigerator")
	}
	return &ElectricalOvenImplementor{
		temperature: t,
		ovenType:    "ElectricalOven",
	}
}

func (ei *ElectricalOvenImplementor) WarmUp(temperature float64) {
	if temperature < 0 {
		panic("This isn't a refrigerator")
	}
	time.Sleep(time.Duration(temperature-ei.temperature) * 30 * time.Millisecond)
	ei.temperature = temperature
}

func (ei *ElectricalOvenImplementor) CoolDown(temperature float64) {
	if temperature < 0 {
		panic("This isn't a refrigerator")
	}
	time.Sleep(time.Duration(ei.temperature-temperature) * 25 * time.Millisecond)
	ei.temperature = temperature
}

func (ei *ElectricalOvenImplementor) CookPizza(pizza *Pizza) {
	pizza.Cook()
}

func (ei *ElectricalOvenImplementor) GetTemperature() float64 {
	return ei.temperature
}

func (ei *ElectricalOvenImplementor) GetOvenType() string {
	return ei.ovenType
}

// ///////Oven /////////////////
type Oven struct {
	ovenImpl IOvenImplementor
}

func NewOven(ovenImpl IOvenImplementor) *Oven {
	if ovenImpl == nil {
		panic("ovenImpl is nil")
	}
	return &Oven{ovenImpl}
}

func (o *Oven) GetTemperature() float64 {
	return o.ovenImpl.GetTemperature()
}

func (o *Oven) GetImplementorType() string {
	return o.ovenImpl.GetOvenType()
}

func (o *Oven) ChangeImplementor(ovenImpl IOvenImplementor) error {
	if ovenImpl == nil {
		return errors.New("ovenImpl is nil")
	}
	fmt.Printf("Implementor changed from %v to %v\n",
		o.GetImplementorType(), ovenImpl.GetOvenType())
	fmt.Println("---------------------------")
	o.ovenImpl = ovenImpl
	return nil
}

func (o *Oven) prepareOven(temperature float64) {
	if o.GetTemperature() > temperature {
		o.ovenImpl.CoolDown(temperature)
	} else if o.GetTemperature() < temperature {
		o.ovenImpl.WarmUp(temperature)
	} else {
		fmt.Println("Ideal temperature")
	}
	fmt.Println("Oven prepared!")
}

func (o *Oven) CookingPizza(pizza *Pizza) {
	o.prepareOven(pizza.CookingTemperature)
	fmt.Printf("Cooking %v pizza for %v minutes at %v C\n",
		pizza.Name, pizza.CookingTime, pizza.CookingTemperature)
	o.ovenImpl.CookPizza(pizza)
	if pizza.IsCooked() {
		fmt.Println("Pizza is ready!!!")
	} else {
		fmt.Println("O_o ... some wrong ...")
	}
	fmt.Println("---------------------------")
}

// //////////////////////////////////////////

func main() {
	firstPizza := NewPizza("Margarita", 10, 220)
	secondPizza := NewPizza("Salami", 9, 180)

	implementor := NewClassicOvenImplementor(10)
	oven := NewOven(implementor)
	fmt.Printf("Implementor type: %v\n", oven.GetImplementorType())
	oven.CookingPizza(firstPizza)
	oven.CookingPizza(secondPizza)
	// change newImplementation
	newImplementor := NewElectricalOvenImplementor(oven.GetTemperature())
	firstPizza = NewPizza("Margarita", 9, 225)
	secondPizza = NewPizza("Salami", 10, 175)
	oven.ChangeImplementor(newImplementor)
	fmt.Printf("Implementor type: %v\n", oven.GetImplementorType())
	oven.CookingPizza(firstPizza)
	oven.CookingPizza(secondPizza)
}
