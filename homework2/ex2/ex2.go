package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Print("Введите площадь круга: ")
	fmt.Scan(&s)
	var d = 2 * math.Sqrt(s/math.Pi)
	fmt.Println("d =", d)
	fmt.Println("L =", math.Pi*d)
}
