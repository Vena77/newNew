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
	fmt.Println("Введите первое число: ")
	fmt.Scanln(&a)
	if operat != "!" {
		fmt.Println("Введите второе число: ")
		fmt.Scanln(&b)
	}

	switch operat {
	case "+":
		fmt.Println(a + b)

	case "-":
		fmt.Printf("%.3g", a-b)

	case "*":
		fmt.Printf("%.3g", a*b)

	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль")
		} else {
			if a/b == math.Trunc(a/b) {
				fmt.Println(a / b)
			} else {
				fmt.Printf("%.2f \n", a/b)
			}
		}
	case "n":
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
