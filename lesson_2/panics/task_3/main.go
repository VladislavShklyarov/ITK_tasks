package main

import "fmt"

func main() {
	Level1() // все гуд)
}

func Level1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Паника обработана на уровне 1: ошибка в Level3")
		}
	}()
	Level2()
}
func Level2() {
	defer fmt.Println("Завершаем Level2")
	Level3()
}
func Level3() {
	panic("Ошибка в Level 3")
}
