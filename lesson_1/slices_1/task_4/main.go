package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a[:2]       // [a:2] = [1, 2] c = 3
	b = append(b, 4) // [1, 2, 4] и 2 элемент в a затирается
	fmt.Println(b)   //  [1, 2, 4]
	fmt.Println(a)   // [1, 2, 4]
}
