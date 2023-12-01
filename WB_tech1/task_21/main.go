package main

import "fmt"

// Адаптер, который конвертирует температуру в Фаренгейтах в Цельсии

// Температура в Цельсиях
type Celsius struct {
	Temperature float64
}

func NewCelsius(t float64) *Celsius {
	return &Celsius{
		Temperature: t,
	}
}

// Температура в Фаренгейтах
type Fahrenheit struct {
	Temperature float64
}

func NewFahrenheit(t float64) *Fahrenheit {
	return &Fahrenheit{
		Temperature: t,
	}
}

// Адаптер преобразования
type FahrenheitAdapter struct {
	ft *Fahrenheit
}

func NewFahrenheitAdapter(ft *Fahrenheit) *FahrenheitAdapter {
	return &FahrenheitAdapter{
		ft: ft,
	}
}

// Преобразование в Цельсии
func (a *FahrenheitAdapter) FToC() *Celsius {
	return &Celsius{
		Temperature: float64((a.ft.Temperature - 32) * 5 / 9),
	}
}

func main() {
	ft := NewFahrenheit(100)
	adapter := NewFahrenheitAdapter(ft)
	ct := adapter.FToC()
	fmt.Printf("%f", ct.Temperature)
}
