package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 3, 4) // [0 0 0] c = 4
	appendingSlice(slice[:1])
	fmt.Println(slice) // [0 1 0]
}

func appendingSlice(slice []int) {
	slice = append(slice, 1) // длина нашего слайса не изменилась, соответственно изменения будут отображены
}
