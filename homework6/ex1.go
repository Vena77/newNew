package main

import "fmt"

type Point struct {
	x int
	p *Point
}

func main() {
	var vv = Point{}
	y := make([]Point, 0, 5)
	fmt.Println(len(y))
	fmt.Println("Введите целые числа, для выхода нажмите exit: ")
	for {
		_, err := fmt.Scanln(&vv.x)
		if err != nil {
			break
		}
		y = append(y, vv)
		if len(y) > 1 {
			y[len(y)-2].p = &y[len(y)-1]
		}
	}
	fmt.Println(y)

	for i, j := 0, len(y)-1; i < j; i, j = i+1, j-1 {
		y[i].x, y[j].x = y[j].x, y[i].x
	}
	fmt.Println(y)

}
