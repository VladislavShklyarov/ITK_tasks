package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 1, 3)
	fmt.Println(nums) // [0]
	appendSlice(nums, 1)
	fmt.Println(nums) // [0]
	copySlice(nums, []int{2, 3})
	fmt.Println(nums) // [2] // т.к copy копирует минимальное по длине
	mutateSlice(nums, 1, 4)
	fmt.Println(nums) // ну и так как у нас только 1 элемент, обращение к индексу 1 вызовет панику
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func copySlice(sl, cp []int) {
	copy(sl, cp)
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}
