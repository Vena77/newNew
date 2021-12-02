package main

import (
	"fmt"
)

func main() {
	var s int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scan(&s)
	fmt.Println(s/100, s%100/10, s%100%10)
}
