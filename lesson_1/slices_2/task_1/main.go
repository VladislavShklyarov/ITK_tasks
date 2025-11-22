package main

import (
	"fmt"
)

func main() {
	var numbers []*int
	for _, value := range []int{10, 20, 30, 40} {
		numbers = append(numbers, &value)
	}
	for _, number := range numbers {
		fmt.Println(number)
		fmt.Println("d", *number)
	}
}

// TODO: спросить у Ромы

/* Создаем неинициализированный слайс указателей на инты

далее в цикле проходимся по значениям обычного слайса с интами, и добавляем в первый слайс адреса элементов оператором &
затем во втором цикле проходимся по этим адресам и разыменовываем их в выводе.

*/
