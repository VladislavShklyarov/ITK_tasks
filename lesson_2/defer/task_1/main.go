package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
	// defer хранит отложенный вызов функции в виде стэка LIFO и срабатывает после завершения остальных функций
	// start
	// end
	// 3
	// 2
	// 1
}
