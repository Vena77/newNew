package main

import (
	"fmt"
)

func main() {
	x := make([]int, 0, 5)
	fmt.Println("Введите целые числа, для выхода нажмите exit: ")
	for {
		var a int
		_, err := fmt.Scanln(&a)
		if err != nil {
			break
		}
		x = append(x, a)
	}
	for i := 1; i < len(x); i++ {

		for j := i - 1; j > -1; j-- {
			if x[j+1] < x[j] {
				x[j+1], x[j] = x[j], x[j+1]
			}
		}
	}
	fmt.Println(x)
}
