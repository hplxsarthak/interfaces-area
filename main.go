package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	side float64
}
type triangle struct {
	base float64
	height float64
}

func main() {
	s := square{3.0}
	t := triangle{4.0, 5.0}

	printArea(s)
	printArea(t)

}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}