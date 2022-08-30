package main

import "fmt"

// //////////interfaces/////////////////
type OrderItemVisitor interface {
	Visit(item ItemElement) float64
}

type ItemElement interface {
	Accept(visitor OrderItemVisitor) float64
}

// ////////////Pizza   //////////////////
type Pizza struct {
	Name  string
	price int
}

func NewPizza(name string, price int) *Pizza {
	return &Pizza{name, price}
}

func (p *Pizza) Accept(visitor OrderItemVisitor) float64 {
	return visitor.Visit(p)
}

func (p *Pizza) GetPrice() int {
	return p.price
}

// ////////////Coffee    //////////////////
type Coffee struct {
	Name     string
	price    int
	capacity float64
}

func NewCoffee(name string, price int, capacity float64) *Coffee {
	return &Coffee{name, price, capacity}
}

func (p *Coffee) Accept(visitor OrderItemVisitor) float64 {
	return visitor.Visit(p)
}

func (p *Coffee) GetPrice() float64 {
	return float64(p.price) * p.capacity
}

// ////////////WithOutDiscountVisitor //////////////////
type WithOutDiscountVisitor struct{}

func (p *WithOutDiscountVisitor) Visit(item ItemElement) float64 {
	cost := 0.0
	if val, ok := item.(*Pizza); ok {
		cost += float64(val.GetPrice())
	} else if val, ok := item.(*Coffee); ok {
		cost += val.GetPrice()
	}
	return cost
}

// ////////////OnlyPizzaDiscountVisitor //////////////////
type OnlyPizzaDiscountVisitor struct{}

func (p *OnlyPizzaDiscountVisitor) Visit(item ItemElement) float64 {
	cost := 0.0
	if val, ok := item.(*Pizza); ok {
		cost += float64(val.GetPrice()) * 0.85 // 15%
	} else if val, ok := item.(*Coffee); ok {
		cost += val.GetPrice()
	}
	return cost
}

// ////////////OnlyCoffeeDiscountVisitor //////////////////
type OnlyCoffeeDiscountVisitor struct{}

func (p *OnlyCoffeeDiscountVisitor) Visit(item ItemElement) float64 {
	cost := 0.0
	if val, ok := item.(*Pizza); ok {
		cost += float64(val.GetPrice())
	} else if val, ok := item.(*Coffee); ok {
		cost += val.GetPrice() * 0.65 // 35%
	}
	return cost
}

// ////////////AllDiscountVisitor //////////////////
type AllDiscountVisitor struct{}

func (p *AllDiscountVisitor) Visit(item ItemElement) float64 {
	cost := 0.0
	if val, ok := item.(*Pizza); ok {
		cost += float64(val.GetPrice())
	} else if val, ok := item.(*Coffee); ok {
		cost += val.GetPrice()
	}
	cost -= cost * 0.20 // 20%
	return cost
}

// ////////////Waiter //////////////////
type Waiter struct {
	discountCalculator OrderItemVisitor
	orders             []ItemElement
}

func NewWaiter(discount OrderItemVisitor) *Waiter {
	return &Waiter{
		discountCalculator: discount,
		orders:             []ItemElement{},
	}
}

func (w *Waiter) SetOrders(orders []ItemElement) {
	w.orders = orders
}

func (w *Waiter) SetDiscount(discount OrderItemVisitor) {
	w.discountCalculator = discount
}

func (w *Waiter) CalcPrice() float64 {
	price := 0.0
	for _, item := range w.orders {
		price += item.Accept(w.discountCalculator)
	}
	return price
}

// ////////////////////////////////////
func main() {
	orders := []ItemElement{
		NewPizza("Margarita", 33),
		NewCoffee("Latte", 15, 0.3),
		NewPizza("4 Cheese", 24),
		NewPizza("Salami", 40),
		NewCoffee("Cappuccino", 10, 0.5),
	}

	waiter := NewWaiter(&WithOutDiscountVisitor{})
	waiter.SetOrders(orders)
	fmt.Printf("Price without discount: %.2f\n", waiter.CalcPrice())

	waiter.SetDiscount(&OnlyPizzaDiscountVisitor{})
	fmt.Printf("Price discount on pizza only: %.2f\n", waiter.CalcPrice())

	waiter.SetDiscount(&OnlyCoffeeDiscountVisitor{})
	fmt.Printf("Price discount on coffee only: %.2f\n", waiter.CalcPrice())

	waiter.SetDiscount(&AllDiscountVisitor{})
	fmt.Printf("Price discount on all: %.2f\n", waiter.CalcPrice())
}
