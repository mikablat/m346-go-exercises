package main

import (
	"fmt"
	"math"
)

func computeHypotenuse (a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2)+ math.Pow(b, 2))
}

func main() {
	a := 10.0
	b := 12.0
	fmt.Printf("Hypotenuse für a=%v, b=%v, %v\n", a, b, computeHypotenuse(a, b))
}
