package main

import (
	"fmt"
	"math"
)

func Circle(ss float64) float64 {
	return 2 * math.Sqrt(ss/math.Pi)
}

func main() {
	var s float64
	fmt.Print("Введите площадь круга: ")
	fmt.Scan(&s)
	d := Circle(s)
	fmt.Println("d =", d)
	fmt.Println("L =", math.Pi*d)
}
