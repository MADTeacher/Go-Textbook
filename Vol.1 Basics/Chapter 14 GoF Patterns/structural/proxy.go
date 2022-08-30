package main

import "fmt"

// ///////PizzaOrderFlyWeight /////////////////
type PizzaOrderFlyWeight struct {
	PizzaSize string
	Diameter  int
}

func NewPizzaOrderFlyWeight(size string, diameter int) *PizzaOrderFlyWeight {
	return &PizzaOrderFlyWeight{
		PizzaSize: size,
		Diameter:  diameter,
	}
}

func SharedPizzaOrderFlyWeight(shared PizzaOrderFlyWeight) *PizzaOrderFlyWeight {
	return &PizzaOrderFlyWeight{
		PizzaSize: shared.PizzaSize,
		Diameter:  shared.Diameter,
	}
}

//////////////PizzaOrderContext //////////////////
type PizzaOrderContext struct {
	Name        string
	sharedState *PizzaOrderFlyWeight
}

func (p *PizzaOrderContext) String() string {
	return fmt.Sprintf("unique state: %v || shared state: %v",
		p.Name, p.sharedState)
}

func NewPizzaOrderContext(name string,
	sharedState *PizzaOrderFlyWeight) *PizzaOrderContext {
	return &PizzaOrderContext{
		Name:        name,
		sharedState: sharedState,
	}
}

///////////FlyWeightFactory /////////////
type FlyWeightFactory struct {
	flyweights []*PizzaOrderFlyWeight
}

func (f *FlyWeightFactory) GetFlyWeight(st PizzaOrderFlyWeight) *PizzaOrderFlyWeight {
	if len(f.flyweights) > 0 {
		for _, it := range f.flyweights {
			if *it == st {
				return it
			}
		}
	}
	f.flyweights = append(f.flyweights, &st)
	return &st
}

func (f *FlyWeightFactory) Total() int {
	return len(f.flyweights)
}

//////////////interface////////////////////
type IOrder interface {
	MakePizzaOrder(uState string, sState PizzaOrderFlyWeight) *PizzaOrderContext
	TotalSharedStates() int
}

///////////PizzaOrderMaker  /////////////
type PizzaOrderMaker struct {
	flyWeightFactory *FlyWeightFactory
}

func (p *PizzaOrderMaker) MakePizzaOrder(
	uState string, sState PizzaOrderFlyWeight) *PizzaOrderContext {
	flyWeight := p.flyWeightFactory.GetFlyWeight(sState)
	return &PizzaOrderContext{
		Name:        uState,
		sharedState: flyWeight,
	}
}

func (p *PizzaOrderMaker) TotalSharedStates() int {
	return p.flyWeightFactory.Total()
}

//////////////Proxy////////////////////
type ProxyOrderMaker struct {
	subject *PizzaOrderMaker
}

func (p *ProxyOrderMaker) MakePizzaOrder(
	uState string, sState PizzaOrderFlyWeight) *PizzaOrderContext {
	p.logging(uState, sState)
	return p.subject.MakePizzaOrder(uState, sState)
}

func (p *ProxyOrderMaker) TotalSharedStates() int {
	return p.subject.TotalSharedStates()
}

func (p *ProxyOrderMaker) logging(uState string, sState PizzaOrderFlyWeight) {
	fmt.Printf("Logging: unique state: %v || shared state: %v\n",
		uState, sState)
}

func main() {
	pizzaMaker := PizzaOrderMaker{
		flyWeightFactory: &FlyWeightFactory{
			flyweights: []*PizzaOrderFlyWeight{},
		},
	}

	pizzaMakerProxy := ProxyOrderMaker{
		subject: &pizzaMaker,
	}

	sharedState := []PizzaOrderFlyWeight{
		{"Big pizza", 30},
		{"Medium pizza", 20},
		{"Small pizza", 15},
	}
	uniqueStates := []string{"Margarita", "Salami", "4 Cheese"}
	orders := []*PizzaOrderContext{}
	for _, shared := range sharedState {
		for _, unique := range uniqueStates {
			orders = append(orders, pizzaMakerProxy.MakePizzaOrder(
				unique, shared,
			))
		}
	}
	fmt.Printf("Number of pizzas: %v\n", len(orders))
	fmt.Printf("Number of sharedState: %v\n", pizzaMakerProxy.TotalSharedStates())
	for idx, it := range orders {
		fmt.Println("----------------------")
		fmt.Printf("Pizza number in the list: %v\n", idx)
		fmt.Println(it)
	}
}
