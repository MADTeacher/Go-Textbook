package main

import (
	"fmt"
	"math/rand"
)

type OrderType string

const (
	Cappuccino = "cappuccino"
	Latte      = "latte"
	Espresso   = "espresso "
)

// //////////interfaces/////////////////
type Observer interface {
	Update(orderId int)
	ObserverID() int
}

type BaseBarista interface {
	Attach(observer Observer)
	Detach(observer Observer)
	TakeOrder(order Order)
	GetOrder(orderId int) Order
}

// ////////////Order //////////////////
var orderID int = 1

type Order struct {
	id        int
	orderType OrderType
}

func NewOrder(orderType OrderType) *Order {
	order := &Order{
		id:        orderID,
		orderType: orderType,
	}
	orderID++
	return order
}

func (p *Order) Id() int {
	return p.id
}

func (p *Order) String() string {
	return fmt.Sprintf("order (%v):(%v)", p.id, p.orderType)
}

// ////////////Subject  //////////////////
type Subject struct {
	observers      []Observer
	detachObserver Observer
}

func (s *Subject) Attach(observer Observer) {
	isNotContains := s.isConstains(observer)
	if !isNotContains {
		s.observers = append(s.observers, observer)
	}
}

func (s *Subject) isConstains(observer Observer) bool {
	for _, it := range s.observers {
		if it.ObserverID() == observer.ObserverID() {
			return true
		}
	}
	return false
}

func (s *Subject) Detach(observer Observer) {
	isContains := s.isConstains(observer)
	if isContains {
		s.detachObserver = observer
	}
}

func (s *Subject) Notify(orderID int) {
	for _, it := range s.observers {
		it.Update(orderID)
	}
	if s.detachObserver != nil {
		index := -1
		for idx, it := range s.observers {
			if s.detachObserver.ObserverID() == it.ObserverID() {
				index = idx
				break
			}
		}
		if index > 0 {
			copy(s.observers[index:], s.observers[index+1:])
			s.observers[len(s.observers)-1] = nil
			s.observers = s.observers[:len(s.observers)-1]
		} else if len(s.observers) <= 1 && index >= 0 {
			s.observers = []Observer{}
		}
		s.detachObserver = nil
	}
}

// ////////////Barista  //////////////////
type Barista struct {
	Subject
	orders       []Order
	finishOrders []Order
}

func NewBarista() *Barista {
	return &Barista{
		Subject: Subject{
			observers:      []Observer{},
			detachObserver: nil,
		},
		orders:       []Order{},
		finishOrders: []Order{},
	}
}

func (b *Barista) Attach(observer Observer) {
	b.Subject.Attach(observer)
}

func (b *Barista) Detach(observer Observer) {
	b.Subject.Detach(observer)
}

func (b *Barista) TakeOrder(order Order) {
	fmt.Printf("Barista accepted %v \n", &order)
	b.orders = append(b.orders, order)
}

func (b *Barista) GetOrder(orderId int) Order {
	var clientOrder *Order
	index := -1
	for idx, it := range b.finishOrders {
		if it.Id() == orderId {
			clientOrder = &it
			index = idx
			break
		}
	}
	if clientOrder != nil {
		if index >= 0 {
			copy(b.finishOrders[index:], b.finishOrders[index+1:])
			b.finishOrders = b.finishOrders[:len(b.finishOrders)-1]
		} else if len(b.finishOrders) <= 1 && index >= 0 {
			b.finishOrders = []Order{}
		}
		return *clientOrder
	} else {
		panic("O_o")
	}
}

func (b *Barista) ProcessingOrder() {
	if len(b.orders) > 0 {
		order := b.orders[rand.Intn(len(b.orders))]
		index := -1
		for idx, it := range b.orders {
			if it.Id() == order.Id() {
				index = idx
				break
			}
		}
		if index >= 0 {
			copy(b.orders[index:], b.orders[index+1:])
			b.orders = b.orders[:len(b.orders)-1]
		} else if len(b.orders) <= 1 && index >= 0 {
			b.orders = []Order{}
		}
		b.finishOrders = append(b.finishOrders, order)
		fmt.Printf("Barista has completed %v \n", &order)
		b.Notify(order.Id())
	} else {
		fmt.Println("Barista rubs the coffee machine")
	}
}

// ////////////Client   //////////////////
var clientID int = 1

type Client struct {
	id      int
	name    string
	barista BaseBarista
	order   *Order
}

func NewClient(name string, barista BaseBarista) *Client {
	client := &Client{
		id:      clientID,
		name:    name,
		barista: barista,
		order:   nil,
	}
	clientID++
	return client
}

func (c *Client) Update(orderId int) {
	if c.order != nil {
		if c.order.Id() == orderId {
			order := c.barista.GetOrder(orderId)
			fmt.Printf("Client %v took  %v \n", c.name, &order)
			c.barista.Detach(c)
		}
	}
}

func (c *Client) ObserverID() int {
	return c.id
}

func (c *Client) CreateOrder() {
	orderTypes := []OrderType{Cappuccino, Latte, Espresso}
	ot := orderTypes[rand.Intn(len(orderTypes))]
	c.order = NewOrder(ot)
	fmt.Printf("Client %v made   %v \n", c.name, c.order)
	c.barista.Attach(c)
	c.barista.TakeOrder(*c.order)
}

// ////////////////////////////////////
func main() {
	names := []string{
		"Alexander",
		"George",
		"Maksim",
		"Hermann",
		"Oleg",
		"Alexey",
		"Stanislav",
	}

	barista := NewBarista()
	clients := []*Client{}
	for _, it := range names {
		clients = append(clients, NewClient(it, barista))
	}

	for _, client := range clients {
		fmt.Println("*********************************")
		client.CreateOrder()
	}
	fmt.Println("*********************************")
	fmt.Println("****Barista starts to fill orders****")

	for it := 0; it < 10; it++ {
		fmt.Println("*********************************")
		barista.ProcessingOrder()
	}
}
