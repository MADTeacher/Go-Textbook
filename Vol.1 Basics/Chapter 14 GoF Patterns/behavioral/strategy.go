package main

import "fmt"

type ChiefMood string

const (
	GoodMood           = "good"
	BadMood            = "bad"
	BetterStayAwayMood = "betterStayAway"
)

// //////////interfaces/////////////////
type Strategy interface {
	CheckChiefMood(mood ChiefMood) bool
	OrderProseccing(money int) string
}

// ////////////GoodStrategy  //////////////////
type GoodStrategy struct{}

func (i *GoodStrategy) CheckChiefMood(mood ChiefMood) bool {
	return mood != BetterStayAwayMood
}

func (i *GoodStrategy) OrderProseccing(money int) string {
	return "Best drink"
}

// ////////////BadStrategy  //////////////////
type BadStrategy struct{}

func (i *BadStrategy) CheckChiefMood(mood ChiefMood) bool {
	return mood != GoodMood
}

func (i *BadStrategy) OrderProseccing(money int) string {
	return "Glass of water"
}

// ////////////NormalStrategy //////////////////
type NormalStrategy struct{}

func (i *NormalStrategy) CheckChiefMood(mood ChiefMood) bool {
	return true
}

func (i *NormalStrategy) OrderProseccing(money int) string {
	if money < 5 {
		return "Politely refuse the order"
	} else if money < 10 {
		return "Prepare espresso"
	} else if money < 20 {
		return "Prepare cappuccino"
	} else if money < 50 {
		return "Make excellent coffee"
	} else {
		return "Best drink"
	}
}

// ////////////Barista  //////////////////
type Barista struct {
	strategy  Strategy
	chiefMood ChiefMood
}

func NewBarista(strategy Strategy, chiefMood ChiefMood) *Barista {
	fmt.Printf("Chief's Initial Mood: %v \n", chiefMood)
	return &Barista{
		strategy:  strategy,
		chiefMood: chiefMood,
	}
}

func (b *Barista) SetChiefMood(mood ChiefMood) {
	b.chiefMood = mood
}

func (b *Barista) SetStrategy(strategy Strategy) {
	b.strategy = strategy
}

func (b *Barista) TakeOrder(money int) {
	fmt.Printf("The client gives %v parrots for the order\n", money)
	if b.strategy.CheckChiefMood(b.chiefMood) {
		fmt.Println(b.strategy.OrderProseccing(money))
	} else {
		fmt.Println("Barista did not notice the client!")
	}
}

// ////////////////////////////////////
func main() {
	barista := NewBarista(&NormalStrategy{}, BetterStayAwayMood)
	barista.TakeOrder(20)
	barista.TakeOrder(50)
	barista.SetStrategy(&BadStrategy{})
	barista.TakeOrder(40)
	barista.TakeOrder(200)
	barista.SetStrategy(&GoodStrategy{})
	barista.TakeOrder(40)
	barista.SetChiefMood(GoodMood)
	barista.TakeOrder(0)
}
