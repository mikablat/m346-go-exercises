package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a, b, c float64) float64{
Discirminant := -b+math.Sqrt((math.Pow(b,2)-4*a*c))
X := Discirminant/2*a
return X
}

func main() {
a := 10.0
b := 15.0
c := 20.0
fmt.Printf("Discriminant x: %v für a: %v und b: %v und c: %v",computeQuadraticFormula(a,b,c),a,b,c)
}
