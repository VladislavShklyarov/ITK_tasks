package main

import (
	"fmt"
)

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func checkErr(err error) {
	fmt.Println(err == nil)
}

func main() {
	var e1 error
	checkErr(e1) // true
	// в данном случае e1 это пустой неинициализированный интерфейс типа error,
	// его iface = nil, так как и type = nil и data = nil.
	var e *errorString
	checkErr(e) // false
	// e это указатель на структуру errorString, которая удовлетворяет интерфейсу Error.
	// т.к. есть конкретный тип, iface (type = *errorString), проверка не проходит.
	e = &errorString{}
	checkErr(e) // false
	// здесь e это разыменованный указатель. Аналогично предыдущему type= *errorString,
	// только теперь еще и data = &errorString{}

	e = nil
	// сам по себе указатель e = nil. Но так как внутри функции мы приводим ео к интерфейсу,
	// то в type iface кладется *errorString

	checkErr(e) //false
}
