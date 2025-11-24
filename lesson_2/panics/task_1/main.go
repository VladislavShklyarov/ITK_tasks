package main

import "fmt"

func CausePanic() {
	panic("Что-то пошло не так!")
}

func HandlePanic() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\"Паника перехвачена: ...\"")
		}
	}()
	CausePanic()

}

func main() {
	HandlePanic() // Ловим панику и продолжаем выполнение
}
