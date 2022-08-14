package main

import "fmt"

// /////////interfaces/////////////////
type IProduct interface {
	Cost() int
	Name() string
}

// ///////Product /////////////////
type Product struct {
	name string
	cost int
}

func (p *Product) Cost() int {
	return p.cost
}

func (p *Product) Name() string {
	return p.name
}

func NewProduct(name string, cost int) *Product {
	return &Product{
		name: name,
		cost: cost,
	}
}

// ///////CompoundProduct  /////////////////
type CompoundProduct struct {
	name        string
	productList []IProduct
}

func NewCompoundProduct(name string) *CompoundProduct {
	return &CompoundProduct{
		name: name,
	}
}

func (cp *CompoundProduct) Cost() int {
	costProduct := 0
	for _, it := range cp.productList {
		costProduct += it.Cost()
	}
	return costProduct
}

func (cp *CompoundProduct) Name() string {
	return cp.name
}

func (cp *CompoundProduct) AddProduct(product IProduct) {
	cp.productList = append(cp.productList, product)
}

func (cp *CompoundProduct) RemoveProduct(product IProduct) {
	if len(cp.productList) > 0 {
		cp.productList = cp.productList[:len(cp.productList)-1]
	}
}

// /////////Pizza ////////////////////
type Pizza struct {
	CompoundProduct
}

func NewPizza(name string) *Pizza {
	return &Pizza{
		CompoundProduct: CompoundProduct{
			name: name,
		},
	}
}

func (p *Pizza) Cost() int {
	costProduct := 0
	for _, it := range p.productList {
		currentCost := it.Cost()
		fmt.Printf("Cost %v = %v parrots\n", it.Name(), currentCost)
		costProduct += it.Cost()
	}
	fmt.Printf("Cost pizza '%v' = %v parrots\n", p.Name(), costProduct)
	return costProduct
}

func main() {
	dough := NewCompoundProduct("Dough")
	dough.AddProduct(NewProduct("flour", 3))
	dough.AddProduct(NewProduct("egg", 23))
	dough.AddProduct(NewProduct("salt", 1))
	dough.AddProduct(NewProduct("sugar", 2))

	sauce := NewProduct("Barbecue", 12)

	topping := NewCompoundProduct("Topping")
	topping.AddProduct(NewProduct("dorblue", 14))
	topping.AddProduct(NewProduct("parmesan", 12))
	topping.AddProduct(NewProduct("mozzarella", 94))
	topping.AddProduct(NewProduct("maasdam", 77))

	pizza := NewPizza("4 Cheese")
	pizza.AddProduct(dough)
	pizza.AddProduct(sauce)
	pizza.AddProduct(topping)
	pizza.Cost()
}
