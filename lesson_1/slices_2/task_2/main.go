package main

import (
	"fmt"
	"strings"
)

func chengeSlice(arr []string) {
	arr[0] = "Goodbye"
}

func appendSomeData(arr []string) {
	arr = append(arr, "!")
}

func main() {
	someSlice := []string{"Hello", "World"}
	chengeSlice(someSlice)
	appendSomeData(someSlice)
	fmt.Println(strings.Join(someSlice, ""))
}

/*
	Так происходит потому, что внутри функции мы увеличиваем длину слайса, и результат не попадает в main.
	Чтобы этого избежать, нам нужно или присвоить значение работы функции:

	func appendSomeData(arr []string) []string {
		arr = append(arr, "!")
		return arr //Возвращаем результат
	}

	func main() {
		someSlice := []string{"Hello", "World"}
		chengeSlice(someSlice)
		someSlice = appendSomeData(someSlice) //Присваиваем
		fmt.Println(strings.Join(someSlice, ""))
	}

	или изначально сделать len 3 и cap = 4 и передать при помощи среза.
	В таком случае во время append у нас перезапишется пустой 3-й элемент на !

	func main() {
		someSlice := make([]string, 3, 4) // Задаем длину
		someSlice[0] = "Hello"
		someSlice[1] = "World"
		chengeSlice(someSlice)
		appendSomeData(someSlice[:2]) // someSlice[:2] = ["Goodbye" "World"] l= 2 c = 4
		fmt.Println(strings.Join(someSlice, ""))
	}

*/
