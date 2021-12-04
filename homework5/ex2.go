package main

import "fmt"

var fibb = make(map[uint]uint)

func fibbonach(n uint) uint {
	if _, ok := fibb[n]; ok {
		return fibb[n]
	} else {
		fibb[n] = fibbonach(n-1) + fibb[n-2]
	}
	return fibb[n]
}
func main() {
	fibb[0] = 0
	fibb[1] = 1
	var a uint
	fmt.Println("Введите целое число")
	fmt.Scanln(&a)
	fmt.Println(fibbonach(a))
}
