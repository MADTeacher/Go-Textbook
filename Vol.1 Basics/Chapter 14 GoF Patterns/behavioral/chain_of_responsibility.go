package main

import "fmt"

type OrderType int

const (
	VeganOrder OrderType = iota
	NotVeganOrder
	BingeOrder
	NoneOrder
)

// ////////////interface////////////////////
type IHandler interface {
	Execute(order *RequestOrder)
	SetNext(next IHandler)
}

// ////////////RequestOrder //////////////////
type RequestOrder struct {
	description []string
	orderType   OrderType
}

func (r *RequestOrder) GetDescription() []string {
	return r.description
}

func (r *RequestOrder) GetOrderType() OrderType {
	return r.orderType
}

// ///////WaiterHandler /////////////////
type WaiterHandler struct {
	next IHandler
}

func (w *WaiterHandler) Execute(order *RequestOrder) {
	if order.GetOrderType() != NoneOrder {
		fmt.Println("Waiter passes the order on")
		if w.next != nil {
			w.next.Execute(order)
		}
	} else {
		fmt.Println("Waiter rejects the customer's request")
	}

}

func (w *WaiterHandler) SetNext(next IHandler) {
	w.next = next
}

// /////////KitchenHandler ////////////
type KitchenHandler struct {
	next IHandler
}

func (w *KitchenHandler) Execute(order *RequestOrder) {
	orderType := order.GetOrderType()
	if orderType == VeganOrder || orderType == NotVeganOrder {
		fmt.Println("Сhef has started to fulfill the order")
	} else {
		fmt.Println("Сhef rejects the customer's request")
		if w.next != nil {
			w.next.Execute(order)
		}
	}
}

func (w *KitchenHandler) SetNext(next IHandler) {
	w.next = next
}

// /////////BarmanHandler  /////////////
type BarmanHandler struct {
	next IHandler
}

func (w *BarmanHandler) Execute(order *RequestOrder) {
	if order.GetOrderType() == BingeOrder {
		fmt.Println("Barman pours the ordered drink")
	} else {
		fmt.Println("Barman rejects the customer's request")
		if w.next != nil {
			w.next.Execute(order)
		}
	}
}

func (w *BarmanHandler) SetNext(next IHandler) {
	w.next = next
}

//////////////////////////////////

func requestHandle(waiter *WaiterHandler, requestOrder *RequestOrder) {
	fmt.Println("**********Order processing**********")
	fmt.Println("Client request", requestOrder.GetDescription())
	waiter.Execute(requestOrder)
}

func main() {
	kitchen := &KitchenHandler{}
	bar := &BarmanHandler{next: kitchen}
	waiter := &WaiterHandler{}
	waiter.SetNext(bar)

	requestList := []string{"Borsch", "Naval macaroshki"}
	requestOrder := &RequestOrder{requestList, NotVeganOrder}
	requestHandle(waiter, requestOrder)

	requestList = []string{"Bloody Merry", "Cognac", "Whiskey"}
	requestOrder = &RequestOrder{requestList, BingeOrder}
	requestHandle(waiter, requestOrder)

	requestList = []string{"The world on a silver platter!"}
	requestOrder = &RequestOrder{requestList, NoneOrder}
	requestHandle(waiter, requestOrder)
}
