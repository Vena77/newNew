package main

import (
	"fmt"
	"math"
)

func main() {
	const pi float64 = 3.1415
	var s float64
	fmt.Print("Введите площадь круга: ")
	fmt.Scan(&s)
	var d = math.Sqrt(s / pi)
	fmt.Println("d =", d)
	fmt.Println("L =", pi*d)
}
