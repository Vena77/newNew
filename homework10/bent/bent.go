package main

import "fmt"

var fibb = make(map[uint]uint)

func fibbonach(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonach(n-1) + fibbonach(n-2)
}

func main() {
	var a uint
	fibb[0] = 0
	fibb[1] = 1
	fmt.Println("Введите целое число: ")
	fmt.Scanln(&a)
	fmt.Println(fibbonach(a))
}
