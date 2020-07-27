package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{sideLength: 12}
	t := triangle{base: 2, height: 10}

	printAreaShape(s)
	printAreaShape(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printAreaShape(s shape) {
	fmt.Println(s.getArea())
}
