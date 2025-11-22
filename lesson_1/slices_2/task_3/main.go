package main

import "fmt"

func test(testSlice []string) {
	testSlice = append(testSlice, "Пока")
}
func main() {
	testSlice := make([]string, 0, 3)
	testSlice = append(testSlice, "Привет")
	testSlice = append(testSlice, "Привет")
	test(testSlice[:1])
	fmt.Println(testSlice)
}

/*Анализ: создали слайс с len 0 cap 3 и cделали два аппенда. Затем передали слайс в test и попытались добавить "Пока"
Не добавилось, поскольку эта операция привела к увеличению длины массива и мы не вернули значения из функции

Исправления зависят от необходимого результат

Если цель добавить "Пока" в конец, то можно просто присвоить результат работы функции

	func test(testSlice []string) []string {
		testSlice = append(testSlice, "Пока")
		return testSlice
	}
	func main() {
		testSlice := make([]string, 0, 3)
		testSlice = append(testSlice, "Привет")
		testSlice = append(testSlice, "Привет")
		testSlice = test(testSlice)
		fmt.Println(testSlice)
	}

Если перезаписать последнее "Привет" на "Пока", то можно либо просто присвоить по индексу:

	func test(testSlice []string) {
	testSlice[1] = "Пока"
	}

Или также затереть при помощи среза:

func main() {
	testSlice := make([]string, 0, 3)
	testSlice = append(testSlice, "Привет")
	testSlice = append(testSlice, "Привет")
	test(testSlice[:1])
	fmt.Println(testSlice)
}


*/
