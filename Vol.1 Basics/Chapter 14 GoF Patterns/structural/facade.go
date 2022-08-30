package main

import "fmt"

/////////////////Menu Enum/////////////////
const (
	VeganMenuType = iota
	NotVeganMenuType
	MixedMenuType
)

// /////////interfaces/////////////////
type IMenu interface {
	GetMenuName() string
}

type IClient interface {
	RequestMenu(menu IMenu)
	FormOrder() map[int]string
	EatingFood()
	GetName() string
}

// ///////Menu /////////////////
type VeganMenu struct{}

func (m *VeganMenu) GetMenuName() string {
	return "Vegan menu"
}

type NotVeganMenu struct{}

func (m *NotVeganMenu) GetMenuName() string {
	return "Not vegan menu"
}

type MixedMenu struct{}

func (m *MixedMenu) GetMenuName() string {
	return "Mixed menu"
}

////////////Kitchen and Waiter/////////
type Kitchen struct{}

func (m *Kitchen) PrepareFood() {
	fmt.Println("The ordered food is being prepared")
}

func (m *Kitchen) CallWaiter() {
	fmt.Println("Food at the waiter")
}

type Waiter struct{}

func (m *Waiter) TakeOrder(client IClient) {
	fmt.Printf("Waiter accepted the order from %v\n", client.GetName())
}

func (m *Waiter) SendToKitchen() {
	fmt.Println("Ordering in the kitchen")
}

func (m *Waiter) ServeClient(client IClient) {
	fmt.Printf("Dishes are ready, we bring them to the client with name %v\n",
		client.GetName())
}

//////////////Client//////////////////
type Client struct {
	name string
}

func (c *Client) RequestMenu(menu IMenu) {
	fmt.Printf("Client %v familiarizes with %v\n",
		c.GetName(), menu.GetMenuName())
}

func (c *Client) FormOrder() map[int]string {
	fmt.Printf("Client %v makes an order\n", c.GetName())
	return map[int]string{1: "Something"}
}

func (c *Client) EatingFood() {
	fmt.Printf("Client %v starts eating\n", c.GetName())
}

func (c *Client) GetName() string {
	return c.name
}

///////////Facade///////////////////////
type PizzeriaFacade struct {
	kitchen Kitchen
	waiter  Waiter
	menu    map[int]IMenu
}

func NewPizPizzeriaFacade() *PizzeriaFacade {
	return &PizzeriaFacade{
		kitchen: Kitchen{},
		waiter:  Waiter{},
		menu: map[int]IMenu{
			VeganMenuType:    &VeganMenu{},
			NotVeganMenuType: &NotVeganMenu{},
			MixedMenuType:    &MixedMenu{},
		},
	}
}

func (p *PizzeriaFacade) GetMenu(menuType int) IMenu {
	return p.menu[menuType]
}

func (p *PizzeriaFacade) TakeOrder(client IClient) {
	p.waiter.TakeOrder(client)
	p.waiter.SendToKitchen()
	p.kitchenWork()
	p.waiter.ServeClient(client)
}

func (p *PizzeriaFacade) kitchenWork() {
	p.kitchen.PrepareFood()
	p.kitchen.CallWaiter()
}

func main() {
	pizzeria := NewPizPizzeriaFacade()
	client1 := Client{"Ivan"}
	client2 := Client{"Alex"}
	client1.RequestMenu(pizzeria.GetMenu(MixedMenuType))
	pizzeria.TakeOrder(&client1)
	client2.RequestMenu(pizzeria.GetMenu(VeganMenuType))
	pizzeria.TakeOrder(&client2)
	client1.EatingFood()
	client2.EatingFood()
}
