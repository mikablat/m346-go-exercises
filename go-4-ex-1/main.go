package main

import "fmt"

func computeGrade(E,M float64) float64{
	N := E/M * 5 + 1
	return N
}
func main() {
	fmt.Println(computeGrade(45, 60))
	fmt.Println(computeGrade(50, 90))
	fmt.Println(computeGrade(4, 10))
}
