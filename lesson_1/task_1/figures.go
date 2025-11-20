package main

import "math"

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, length float64
}

func NewCircle(radius float64) Circle {
	return Circle{radius: radius}
}

func NewRectangle(width, length float64) Rectangle {
	return Rectangle{
		width:  width,
		length: length,
	}
}

func (c *Circle) Area() float64 {
	area := c.radius * math.Pi
	return math.Round(area*100) / 100
}

func (c *Circle) Perimeter() float64 {
	perimeter := 2 * c.radius * math.Pi
	return math.Round(perimeter*100) / 100
}

func (r *Rectangle) Area() float64 {
	area := r.width * r.length
	return math.Round(area*100) / 100
}

func (r *Rectangle) Perimeter() float64 {
	perimeter := 2 * (r.width + r.length)
	return math.Round(perimeter*100) / 100
}
