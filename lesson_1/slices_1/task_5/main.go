package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	src := arr[:1] // arr[:1] = [1] c = 3
	foo(src)
	fmt.Println(src) // [1]
	fmt.Println(arr) // [1, 5, 3]
}

func foo(src []int) {
	src = append(src, 5) // в этом append меняется длина слайса и изменения src не попадают в main
	// НО поскольку мы добавили 5 в срез a[:1], и по факту длина a не изменилась, но изменился сам подлежащий массив
	// эти изменения появятся в оригинальном слайсе
}
