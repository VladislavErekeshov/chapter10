package main

import (
	"fmt"
	"math"
)

type circle struct {
	x float64
	y float64
	r float64
}
type rest struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func (r rest) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}
func main() {
	c := circle{x: 0, y: 0, r: 5}
	r := rest{0, 0, 10, 10}

	fmt.Println(r.area())
	fmt.Println(c.area())
}
