package main

import "fmt"

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

func main() {
	err := handle()
	if err != nil {
		fmt.Println(err)
	}
}

func handle() error {
	return MyError{
		msg: "error occurred"}
}

// Я так понял суть в том, что тип error это просто интерфейс c методом Error() string
// И любой с таким методом ему удовлетворяет
