package main

import "fmt"

var fibb = make(map[uint]uint)

func fibbonach1(n uint) uint {
	if _, ok := fibb[n]; ok {
		return fibb[n]
	} else {
		fibb[n] = fibbonach1(n-1) + fibb[n-2]
	}
	return fibb[n]
}
func main() {
	fibb[0] = 0
	fibb[1] = 1
	var a uint
	fmt.Println("Введите целое число")
	fmt.Scanln(&a)
	fmt.Println(fibbonach1(a))
}
