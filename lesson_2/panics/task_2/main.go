package main

import "fmt"

func main() {
	fmt.Println(SafeDivide(10, 2)) // Ожидаемый результат: 5
	fmt.Println(SafeDivide(10, 0)) // Ожидаемый результат: 0 (без паники)
}

func SafeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = 0
		}
	}()

	if b == 0 {
		panic("деление на 0")
	}
	return a / b
}
