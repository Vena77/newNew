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
	//x := [10]int{3, 6, 9, 67, 0, 1, 4, 6, 8, 9}
	for i := 1; i < len(x); i++ {
		y := x[i]
		ind := i
		for j := i - 1; j > -1; j-- {
			if x[j] > y {
				x[j+1] = x[j]
				ind = j
			}
			x[ind] = y
		}
	}
	fmt.Println(x)
}
