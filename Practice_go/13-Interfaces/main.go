package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64 // anything that has area() will have shape type
}

type circle struct {
	radius float64
}

type rect struct {
	width float64
	height float64
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5, 7}


	shapes := []shape{c1, r1}

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

