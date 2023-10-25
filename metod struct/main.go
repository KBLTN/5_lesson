package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func movePoint(p Point, x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func (p *Point) movePointPtr(x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	p := Point{1, 2}
	//fmt.Println(p)
	//fmt.Println(movePoint(p, 1, 1))
	//fmt.Println(p)
	//p.movePointPtr(1, 1)
	//fmt.Println(p)
	v := &p
	v.movePointPtr(2, 3)
	fmt.Println(p)
	fmt.Println(v)
}
