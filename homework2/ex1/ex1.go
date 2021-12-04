package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Print("Введите длину прямоугольника: ")
	fmt.Scan(&a)
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scan(&b)

	fmt.Println("S =", a*b)
}
