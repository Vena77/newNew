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
	fmt.Println("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Println("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Println("Введите операцию (+, -, *, /, n-для возведения в степень, !-факториал): ")
	var operat string
	fmt.Scanln(&operat)

	switch operat {
	case "+":
		fmt.Println(a + b)

	case "-":
		if a-b == math.Trunc(a-b) {
			fmt.Println(a - b)
		} else {
			fmt.Printf("%.2f \n", a-b)
		}
	case "*":
		fmt.Printf("%.3g", a*b)
		if a*b == math.Trunc(a*b) {
			fmt.Println(a * b)
		} else {
			fmt.Printf("%.2f \n", a*b)
		}
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
		var c = a
		var d = int(b)
		for i := 1; i < d; i++ {
			c = c * a
		}
		if c == math.Trunc(c) {
			fmt.Println(c)
		} else {
			fmt.Printf("%.2f \n", c)
		}
	case "!":
		if a < 0 {
			a = -a
		}
		var d = uint(a)
		fmt.Println(factorial(d))
	default:
		fmt.Println("Некорректная операция")
	}

}
