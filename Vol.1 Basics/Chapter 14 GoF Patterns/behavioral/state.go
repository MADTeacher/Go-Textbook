package main

import (
	"fmt"
)

type CoffeMachineState int

const (
	NoneSt = iota
	IdleSt
	ChooseSt
	CappuccinoSt
	LatteSt
	EspressoSt
	ChangeMoneySt
)

// //////////interfaces/////////////////
type ICoffeMachine interface {
	GetWaterValue() float64
	GetMilkValue() float64
	GetOrderMoney() int

	SetWaterValue(value float64)
	SetMilkValue(value float64)
	SetOrderMoney(money int)

	SetState(state CoffeMachineState)
	SelectedCoffee() CoffeMachineState
	ReturnMoney()
}

type State interface {
	InsertMoney(coffeMachine ICoffeMachine)
	EjectMoney(coffeMachine ICoffeMachine)
	MakeCoffe(coffeMachine ICoffeMachine)
}

// ////////////IdleState //////////////////
type IdleState struct{}

func (i *IdleState) EjectMoney(coffeMachine ICoffeMachine) {
	fmt.Println("What the money? Oo")
}

func (i *IdleState) InsertMoney(coffeMachine ICoffeMachine) {
	fmt.Println("Go to the choose state")
	coffeMachine.SetState(ChooseSt)
}

func (i *IdleState) MakeCoffe(coffeMachine ICoffeMachine) {
	fmt.Println("Get out of here, rogue")
}

// ////////////WaitChooseState //////////////////
type WaitChooseState struct{}

func (i *WaitChooseState) EjectMoney(coffeMachine ICoffeMachine) {
	fmt.Println("Order or leave your money!")
}

func (i *WaitChooseState) InsertMoney(coffeMachine ICoffeMachine) {
	fmt.Println("Enough funds uploaded to order?")
	coffeMachine.SetState(ChooseSt)
}

func (i *WaitChooseState) MakeCoffe(coffeMachine ICoffeMachine) {
	if coffeMachine.SelectedCoffee() == NoneSt {
		fmt.Println("Choose the coffee you want to make!")
	} else {
		coffeMachine.SetState(coffeMachine.SelectedCoffee())
	}
}

// ////////////ChangeState //////////////////
type ChangeState struct{}

func (i *ChangeState) EjectMoney(coffeMachine ICoffeMachine) {
	fmt.Printf("Return %v parrots\n", coffeMachine.GetOrderMoney())
	coffeMachine.SetOrderMoney(0)
	coffeMachine.SetState(IdleSt)
}

func (i *ChangeState) InsertMoney(coffeMachine ICoffeMachine) {
	i.EjectMoney(coffeMachine)
}

func (i *ChangeState) MakeCoffe(coffeMachine ICoffeMachine) {
	i.EjectMoney(coffeMachine)
}

// ////////////CappuccinoState  //////////////////
type CappuccinoState struct{}

func (i *CappuccinoState) EjectMoney(coffeMachine ICoffeMachine) {
	fmt.Printf("You will not get it!!!\n")
}

func (i *CappuccinoState) InsertMoney(coffeMachine ICoffeMachine) {
	i.MakeCoffe(coffeMachine)
}

func (i *CappuccinoState) MakeCoffe(coffeMachine ICoffeMachine) {
	cost := 32
	needWater := 0.3
	needMilk := 0.1
	waterResidues := coffeMachine.GetWaterValue() - needWater
	milkResidues := coffeMachine.GetMilkValue() - needMilk
	moneyResidues := coffeMachine.GetOrderMoney() - cost
	if moneyResidues >= 0 {
		if waterResidues >= 0 && milkResidues >= 0 {
			fmt.Println("Cooking Cappuccino!")
			coffeMachine.SetWaterValue(waterResidues)
			coffeMachine.SetMilkValue(milkResidues)
			coffeMachine.SetOrderMoney(moneyResidues)
		} else {
			fmt.Println("Not enough ingredients")
		}
		if coffeMachine.GetOrderMoney() > 0 {
			coffeMachine.SetState(ChangeMoneySt)
			coffeMachine.ReturnMoney()
		} else {
			coffeMachine.SetState(IdleSt)
		}
	} else {
		fmt.Println("Not enough funds!")
	}
}

// ////////////LatteState  //////////////////
type LatteState struct{}

func (i *LatteState) EjectMoney(coffeMachine ICoffeMachine) {
	fmt.Printf("You will not get it!!!\n")
}

func (i *LatteState) InsertMoney(coffeMachine ICoffeMachine) {
	i.MakeCoffe(coffeMachine)
}

func (i *LatteState) MakeCoffe(coffeMachine ICoffeMachine) {
	cost := 40
	needWater := 0.3
	needMilk := 0.2
	waterResidues := coffeMachine.GetWaterValue() - needWater
	milkResidues := coffeMachine.GetMilkValue() - needMilk
	moneyResidues := coffeMachine.GetOrderMoney() - cost
	if moneyResidues >= 0 {
		if waterResidues >= 0 && milkResidues >= 0 {
			fmt.Println("Cooking Latte!")
			coffeMachine.SetWaterValue(waterResidues)
			coffeMachine.SetMilkValue(milkResidues)
			coffeMachine.SetOrderMoney(moneyResidues)
		} else {
			fmt.Println("Not enough ingredients")
		}
		if coffeMachine.GetOrderMoney() > 0 {
			coffeMachine.SetState(ChangeMoneySt)
			coffeMachine.ReturnMoney()
		} else {
			coffeMachine.SetState(IdleSt)
		}
	} else {
		fmt.Println("Not enough funds!")
	}
}

// ////////////EspressoState  //////////////////
type EspressoState struct{}

func (i *EspressoState) EjectMoney(coffeMachine ICoffeMachine) {
	fmt.Printf("You will not get it!!!\n")
}

func (i *EspressoState) InsertMoney(coffeMachine ICoffeMachine) {
	i.MakeCoffe(coffeMachine)
}

func (i *EspressoState) MakeCoffe(coffeMachine ICoffeMachine) {
	cost := 25
	needWater := 0.3
	waterResidues := coffeMachine.GetWaterValue() - needWater
	moneyResidues := coffeMachine.GetOrderMoney() - cost
	if moneyResidues >= 0 {
		if waterResidues >= 0 {
			fmt.Println("Cooking Espresso!")
			coffeMachine.SetWaterValue(waterResidues)
			coffeMachine.SetOrderMoney(moneyResidues)
		} else {
			fmt.Println("Not enough ingredients")
		}
		if coffeMachine.GetOrderMoney() > 0 {
			coffeMachine.SetState(ChangeMoneySt)
			coffeMachine.ReturnMoney()
		} else {
			coffeMachine.SetState(IdleSt)
		}
	} else {
		fmt.Println("Not enough funds!")
	}
}

// ////////////CoffeeMachine  //////////////////
type CoffeeMachine struct {
	waterCapacity  float64
	milkCapacity   float64
	orderMopney    int
	selectedCoffee CoffeMachineState
	allStates      map[CoffeMachineState]State
	currentState   State
}

func NewCoffeeMachine(waterCapacity, milkCapacity float64) *CoffeeMachine {
	states := map[CoffeMachineState]State{
		IdleSt:        &IdleState{},
		ChooseSt:      &WaitChooseState{},
		ChangeMoneySt: &ChangeState{},
		CappuccinoSt:  &CappuccinoState{},
		LatteSt:       &LatteState{},
		EspressoSt:    &EspressoState{},
	}
	return &CoffeeMachine{
		waterCapacity:  waterCapacity,
		milkCapacity:   milkCapacity,
		orderMopney:    0,
		selectedCoffee: NoneSt,
		currentState:   states[IdleSt],
		allStates:      states,
	}
}

func (cm *CoffeeMachine) GetWaterValue() float64 {
	return cm.waterCapacity
}

func (cm *CoffeeMachine) GetMilkValue() float64 {
	return cm.milkCapacity
}

func (cm *CoffeeMachine) GetOrderMoney() int {
	return cm.orderMopney
}

func (cm *CoffeeMachine) SetWaterValue(value float64) {
	cm.waterCapacity = value
}

func (cm *CoffeeMachine) SetMilkValue(value float64) {
	cm.milkCapacity = value
}

func (cm *CoffeeMachine) SetOrderMoney(money int) {
	cm.orderMopney = money
}

func (cm *CoffeeMachine) SetState(state CoffeMachineState) {
	if state == IdleSt {
		cm.selectedCoffee = NoneSt
	}
	cm.currentState = cm.allStates[state]
}

func (cm *CoffeeMachine) SelectedCoffee() CoffeMachineState {
	return cm.selectedCoffee
}

func (cm *CoffeeMachine) ReturnMoney() {
	cm.currentState.EjectMoney(cm)
}

func (cm *CoffeeMachine) Cappuccino() {
	fmt.Println("Cappuccino preparation selected")
	cm.selectedCoffee = CappuccinoSt
	cm.currentState.MakeCoffe(cm)
}

func (cm *CoffeeMachine) Latte() {
	fmt.Println("Latte preparation selected")
	cm.selectedCoffee = LatteSt
	cm.currentState.MakeCoffe(cm)
}

func (cm *CoffeeMachine) Espresso() {
	fmt.Println("Espresso preparation selected")
	cm.selectedCoffee = EspressoSt
	cm.currentState.MakeCoffe(cm)
}

func (cm *CoffeeMachine) InsertMoney(money int) {
	cm.orderMopney += money
	fmt.Printf("Inserted %v parrots\n", cm.orderMopney)
	cm.currentState.InsertMoney(cm)
}

func (cm *CoffeeMachine) MakeCoffe() {
	fmt.Println("Start preparation of the selected coffee!")
	cm.currentState.MakeCoffe(cm)
}

// fmt.Printf("%v \n", c.pizza.String())

// ////////////////////////////////////
func main() {
	coffeeMachine := NewCoffeeMachine(1, 1)
	coffeeMachine.MakeCoffe()
	coffeeMachine.InsertMoney(10)
	coffeeMachine.InsertMoney(10)
	coffeeMachine.Cappuccino()
	coffeeMachine.MakeCoffe()
	coffeeMachine.InsertMoney(20)
	fmt.Println("**** When not enough products to make coffee ****")
	coffeeMachine = NewCoffeeMachine(0.1, 0.1)
	coffeeMachine.InsertMoney(100)
	coffeeMachine.MakeCoffe()
	coffeeMachine.Latte()
	coffeeMachine.MakeCoffe()
}
