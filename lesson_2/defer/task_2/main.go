package main

import "fmt"

func main() {
	value := 123
	defer func() {
		fmt.Println(value)
	}()
	changeValue(&value)

}
func changeValue(value *int) {
	*value = 456
}

// defer принимает в себя значения на момент объявления и не изменяет его
// для решения можно использовать замыкание: анонимная функция сохраняет ссылку на значение value, а не его копию и при
// срабатывании defer выведет 456
