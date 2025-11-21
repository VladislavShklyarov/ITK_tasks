package main

import "fmt"

var userStorage map[string]int

func init() {
	userStorage = make(map[string]int)
}

func AddPerson(name string, age int) {
	userStorage[name] = age
}

func GetAge(name string) int {
	return userStorage[name]
}

func DeletePerson(name string) {
	delete(userStorage, name)
}

func PrintAll() {
	for name, age := range userStorage {
		fmt.Printf("%s: %d\n", name, age)
	}
}

func main() {
	AddPerson("Alice", 25)
	AddPerson("Bob", 30)
	AddPerson("Charlie", 35)

	PrintAll()

	fmt.Println(GetAge("Alice"))

	DeletePerson("Bob")
	PrintAll()
}
