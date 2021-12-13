package main

import "fmt"

type Point struct {
	x int
	p *Point
}

var head, next, cur *Point

func (n *Point) Reverse() *Point {
	var cur = n
	var prev *Point
	for cur != nil {
		next := cur.p
		cur.p = prev
		prev = cur
		cur = next
	}
	return prev
}

func (n *Point) Print() {
	var cur = n
	for cur != nil {
		fmt.Println(cur)
		cur = cur.p
	}
}

func main() {
	var a int
	fmt.Println("Введите целые числа, для выхода нажмите exit: ")
	i := 0
	for {
		_, err := fmt.Scanln(&a)
		if err != nil {
			break
		}
		if i == 0 {
			head = &Point{x: a, p: nil}
		} else if i == 1 {
			cur = &Point{x: a, p: nil}
			head.p = cur
		} else {
			next = cur
			cur = &Point{x: a, p: nil}
			next.p = cur
		}
		i++
	}
	next = head
	for {
		if next != nil {
			fmt.Println(next)
			next = next.p
			continue
		}
		break
	}
	head.Reverse().Print() //разворот списка
}
