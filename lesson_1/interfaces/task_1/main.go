package main

import "fmt"

func main() {
	circle := NewCircle(3)
	rectangle := NewRectangle(2, 4)

	fmt.Println(circle.Area())
	fmt.Println(circle.Perimeter())

	fmt.Println(rectangle.Area())
	fmt.Println(rectangle.Perimeter())
}
