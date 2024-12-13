package main

import (
	"fmt"
)

func convertCelsiusToFahrenheit(Celsius float64) float64{
	Fahrenheit := Celsius * 9/5 + 32
	return Fahrenheit
}
func convertFahrenheitToCelsius(Fahrenheit float64) float64{
	Celsius := (Fahrenheit -32)*5/9
	return Celsius
}
func main() {
	C1 := 20.0
	F1 := convertCelsiusToFahrenheit(C1)
	fmt.Printf("Celsius: %v -> Fahrenheit: %v\n",C1,F1)

	F2 := F1
	C2 := convertFahrenheitToCelsius(F2)
	fmt.Printf("Fahrenheit: %v -> Celsius: %v\n",F2,C2)
}
