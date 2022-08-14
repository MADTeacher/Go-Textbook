package main

import "fmt"

const fahrenheitZero = 32.0
const celsiousToFahrenheit = 9.0 / 5.0
const fahrenheitToCelsious = 5.0 / 9.0

// /////////interfaces/////////////////
type IOven interface {
	GetTemperature() float64
	SetTemperature(t float64)
}

type ICelsiousOven interface {
	GetCelsiousTemperature() float64
	SetCelsiousTemperature(t float64)
	GetOriginalTemperature() float64
}

// ///////OriginalOven/////////////////
type OriginalOven struct {
	temperature float64
}

func (oo *OriginalOven) GetTemperature() float64 {
	return oo.temperature
}

func (oo *OriginalOven) SetTemperature(t float64) {
	if oo.temperature < fahrenheitZero {
		panic("Does the oven freeze?..")
	}
	oo.temperature = t
}

func NewOriginalOven(t float64) *OriginalOven {
	if t < fahrenheitZero {
		panic("This isn't a refrigerator")
	}
	return &OriginalOven{t}
}

// ///////OvenAdapter/////////////////
type OvenAdapter struct {
	originalOven IOven
	temperature  float64
}

func NewOvenAdapter(originalOven IOven) *OvenAdapter {
	t := fahrenheitToCelsious * (originalOven.GetTemperature() - fahrenheitZero)
	return &OvenAdapter{
		originalOven: originalOven,
		temperature:  t,
	}
}

func (oa *OvenAdapter) GetCelsiousTemperature() float64 {
	return oa.temperature
}

func (oa *OvenAdapter) SetCelsiousTemperature(t float64) {
	if t < 0 {
		panic("Oo")
	}
	oa.temperature = t
	oa.originalOven.SetTemperature(
		celsiousToFahrenheit*t + fahrenheitZero,
	)
}

func (oa *OvenAdapter) GetOriginalTemperature() float64 {
	return oa.originalOven.GetTemperature()
}

// //////////////////////////////////////////
func PrintTemperature(oven ICelsiousOven) {
	fmt.Printf("Original temperature = %v F\n",
		oven.GetOriginalTemperature())
	fmt.Printf("Celsius  temperature = %v C\n",
		oven.GetCelsiousTemperature())
}
func main() {
	fahrenheitOven := NewOriginalOven(32)
	celsiousOven := NewOvenAdapter(fahrenheitOven)
	PrintTemperature(celsiousOven)
	fmt.Println("----------------")
	fmt.Println("New temperature")
	fmt.Println("----------------")
	celsiousOven.SetCelsiousTemperature(180)
	PrintTemperature(celsiousOven)
}
