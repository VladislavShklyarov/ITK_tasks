package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 3, 4)
	fmt.Println(slice)

	appendSlice(slice)
	fmt.Println(slice)

	mutareSlice(slice)
	fmt.Println(slice)
}

func appendSlice(slice []string) {
	slice = append(slice, "privet")
}
func mutareSlice(slice []string) {
	slice[0] = "vasya"
}

/* создаем слайс строк длиной 3 емкостью 4, выводим его
Далее аппендим в слайс слово "привет". Мы его не возвращаем из функции, при этом длина слайса
меняется, соответственно в main изменения не видны.

Далее внутри функции изменяем 0 элемент слайса. В этом случае, так как мы ссылаемся на тот же самый базовый массив
изменения видны

Если мы хотим добавить privet то нужно либо возвращать значение и присваивать, либо использовать срез [:]

*/
