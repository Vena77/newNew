package main

import (
	"fmt"
	"math"
)

func factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	var a float64
	var b float64
	fmt.Println("Введите операцию (+, -, *, /, n-для возведения в степень, !-факториал): ")
	var operat string
	fmt.Scanln(&operat)


	switch operat {
	case "+":
		fmt.Println(a + b)

	
		fmt.Printf("%.3g", math.Pow(a, b))

	case "!":
		if a < 0 {
			fmt.Println("Отрицательное число")
		} else {
			fmt.Println(factorial(uint(a)))
		}

	default:
		fmt.Println("Некорректная операция")
	}

}
