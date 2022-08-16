package main

import (
	"fmt"
	"math/rand"
)

type OrderType int
type Event int
type WorkerType int

const (
	FoodOrder  OrderType = 1
	BingeOrder OrderType = 2
	////////////////////////
	GetOrderEvent    Event = 10
	FinishOrderEvent Event = 11
	////////////////////////
	WaiterWorker WorkerType = 20
	ChiefWorker  WorkerType = 21
	BarmanWorker WorkerType = 22
	NoneWorker   WorkerType = -1
)

// ////////////interface////////////////////
type IWorker interface {
	TakeOrder(order *Order)
	FinishOrder(order *Order)
	GetWorkerType() WorkerType
	IsHasOrder(order *Order) bool
	GetID() int
}

type IMediator interface {
	Notify(worker IWorker, order *Order, event Event)
	AddWorker(worker IWorker)
	RemoveWorker(worker IWorker)
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

func (p *Order) Type() OrderType {
	return p.orderType
}

func (p *Order) String() string {
	str := ""
	if p.orderType == FoodOrder {
		str = "food"
	} else {
		str = "binge"
	}
	return fmt.Sprintf("order #:%v (%v)", p.id, str)
}

// ///////Worker/////////////////
var workerID int = 1

type Worker struct {
	orders   []*Order
	mediator IMediator
	name     string
	id       int
}

func newWorker(name string, mediator IMediator) Worker {
	worker := Worker{
		mediator: mediator,
		name:     name,
		orders:   []*Order{},
		id:       workerID,
	}
	workerID++
	return worker
}

func (w *Worker) TakeOrder(order *Order) {

}

func (w *Worker) FinishOrder(order *Order) {
	index := -1
	for idx, it := range w.orders {
		if order.id == it.id {
			index = idx
		}
	}
	if index >= 0 {
		copy(w.orders[index:], w.orders[index+1:])
		w.orders[len(w.orders)-1] = nil
		w.orders = w.orders[:len(w.orders)-1]
	} else if len(w.orders) <= 1 {
		w.orders = []*Order{}
	} else {
		panic(order.String() + "not found")
	}
}

func (w *Worker) GetWorkerType() WorkerType {
	return NoneWorker
}

func (w *Worker) IsHasOrder(order *Order) bool {
	for _, it := range w.orders {
		if it.id == order.id {
			return true
		}
	}
	return false
}

func (w *Worker) GetID() int {
	return w.id
}

// ///////Waiter /////////////////
type Waiter struct {
	Worker
}

func NewWaiter(name string, mediator IMediator) *Waiter {
	waiter := &Waiter{
		Worker: newWorker(name, mediator),
	}
	mediator.AddWorker(waiter)
	return waiter
}

func (w *Waiter) TakeOrder(order *Order) {
	w.orders = append(w.orders, order)
	fmt.Printf("Waiter %v take  %v \n", w.name, order)
	w.mediator.Notify(w, order, GetOrderEvent)
}

func (w *Waiter) FinishOrder(order *Order) {
	fmt.Printf("Waiter %v carried %v  to client\n", w.name, order)
	w.Worker.FinishOrder(order)
}

func (w *Waiter) GetWorkerType() WorkerType {
	return WaiterWorker
}

// ///////Barman /////////////////
type Barman struct {
	Worker
}

func NewBarman(name string, mediator IMediator) *Barman {
	barman := &Barman{
		Worker: newWorker(name, mediator),
	}
	mediator.AddWorker(barman)
	return barman
}

func (w *Barman) TakeOrder(order *Order) {
	w.orders = append(w.orders, order)
	fmt.Printf("Barman %v take  %v \n", w.name, order)
}

func (w *Barman) FinishOrder(order *Order) {
	fmt.Printf("Barman %v finish  %v \n", w.name, order)
	w.mediator.Notify(w, order, FinishOrderEvent)
}

func (w *Barman) GetWorkerType() WorkerType {
	return BarmanWorker
}

func (w *Barman) ProcessingOrder() {
	if len(w.orders) > 0 {
		order := w.orders[len(w.orders)-1]
		w.orders = w.orders[:len(w.orders)-1]
		fmt.Printf("Barman %v fulfills %v \n", w.name, order)
		w.FinishOrder(order)
	} else {
		fmt.Printf("Barman %v makes a sad gesture \n", w.name)
	}
}

// ///////Chief /////////////////
type Chief struct {
	Worker
}

func NewChief(name string, mediator IMediator) *Chief {
	chief := &Chief{
		Worker: newWorker(name, mediator),
	}
	mediator.AddWorker(chief)
	return chief
}

func (w *Chief) TakeOrder(order *Order) {
	w.orders = append(w.orders, order)
	fmt.Printf("Chief %v take  %v \n", w.name, order)
}

func (w *Chief) FinishOrder(order *Order) {
	fmt.Printf("Chief %v finish  %v \n", w.name, order)
	w.mediator.Notify(w, order, FinishOrderEvent)
}

func (w *Chief) GetWorkerType() WorkerType {
	return ChiefWorker
}

func (w *Chief) ProcessingOrder() {
	if len(w.orders) > 0 {
		order := w.orders[len(w.orders)-1]
		w.orders = w.orders[:len(w.orders)-1]
		fmt.Printf("Chief %v fulfills %v \n", w.name, order)
		w.FinishOrder(order)
	} else {
		fmt.Printf("Chief %v makes a sad gesture \n", w.name)
	}
}

// ///////WorkersMediator //////////////
type WorkersMediator struct {
	workers map[WorkerType][]IWorker
}

func NewMediator() *WorkersMediator {
	return &WorkersMediator{
		workers: map[WorkerType][]IWorker{
			BarmanWorker: {},
			ChiefWorker:  {},
			WaiterWorker: {},
		},
	}
}

func (w *WorkersMediator) AddWorker(worker IWorker) {
	wt := worker.GetWorkerType()
	w.workers[wt] = append(w.workers[wt], worker)
}

func (w *WorkersMediator) RemoveWorker(worker IWorker) {
	wt := worker.GetWorkerType()
	index := -1
	for idx, it := range w.workers[wt] {
		if worker.GetID() == it.GetID() {
			index = idx
		}
	}
	if index > 0 {
		copy(w.workers[wt][index:], w.workers[wt][index+1:])
		w.workers[wt][len(w.workers[wt])-1] = nil
		w.workers[wt] = w.workers[wt][:len(w.workers[wt])-1]
	} else if index == 0 {
		w.workers[wt] = []IWorker{}
	} else {
		panic("worker with id = " + fmt.Sprint(worker.GetID()) + "not found")
	}
}

func (w *WorkersMediator) Notify(worker IWorker, order *Order, event Event) {
	if event == GetOrderEvent && worker.GetWorkerType() == WaiterWorker {
		if order.Type() == FoodOrder {
			chiefList := w.workers[ChiefWorker]
			if len(chiefList) > 0 {
				randomIndex := rand.Intn(len(chiefList))
				chief := chiefList[randomIndex]
				chief.TakeOrder(order)
			} else {
				fmt.Println("Chief is missing")
			}
		} else if order.Type() == BingeOrder {
			barmanList := w.workers[BarmanWorker]
			if len(barmanList) > 0 {
				randomIndex := rand.Intn(len(barmanList))
				barman := barmanList[randomIndex]
				barman.TakeOrder(order)
			} else {
				fmt.Println("Barman is missing")
			}
		} else {
			fmt.Println("O_o")
		}
	} else if event == FinishOrderEvent &&
		worker.GetWorkerType() != WaiterWorker {
		for _, waiter := range w.workers[WaiterWorker] {
			if waiter.IsHasOrder(order) {
				waiter.FinishOrder(order)
				return
			}
		}
		fmt.Printf("%v wasn't delivered to the client !!! \n", order.String())
	} else {
		panic("event cannot be processing")
	}
	// fmt.Printf("Chief %v makes a sad gesture \n", w.name)
}

// ////////////////////////////////////
func main() {
	mediator := NewMediator()

	waiter1 := NewWaiter("Alexander", mediator)
	waiter2 := NewWaiter("George", mediator)
	waiter3 := NewWaiter("Maksim", mediator)

	barman1 := NewBarman("Hermann", mediator)
	barman2 := NewBarman("Alexey", mediator)

	chief := NewChief("Stanislav", mediator)

	orderTypeList := []OrderType{FoodOrder, BingeOrder}
	orders := []*Order{}
	for it := 0; it < 10; it++ {
		orderType := orderTypeList[rand.Intn(len(orderTypeList))]
		orders = append(orders, NewOrder(orderType))
	}

	waiterList := []*Waiter{waiter1, waiter2, waiter3}
	for _, order := range orders {
		waiter := waiterList[rand.Intn(len(waiterList))]
		waiter.TakeOrder(order)
	}

	fmt.Println("*********************************")
	fmt.Println("******Chef prepares dishes*******")
	fmt.Println("*********************************")
	for it := 0; it < 10; it++ {
		chief.ProcessingOrder()
		fmt.Println("*********************************")
	}

	fmt.Println("*********************************")
	fmt.Println("******Barmens mix cocktails******")
	fmt.Println("*********************************")
	barmanList := []*Barman{barman1, barman2}
	for it := 0; it < 10; it++ {
		barman := barmanList[rand.Intn(len(barmanList))]
		barman.ProcessingOrder()
		fmt.Println("*********************************")
	}
}
