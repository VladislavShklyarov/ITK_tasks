package main

import "fmt"

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func returnError(flag bool) error {
	if flag {
		return &CustomError{"Что-то пошло не так"}
	}
	var err *CustomError
	return err
}

func main() {
	err1 := returnError(true)
	err2 := returnError(false)

	fmt.Println("err1 == nil:", err1 == nil) //false
	fmt.Println("err2 == nil:", err2 == nil) //false
	// В обоих случаях type != nil. Внутри функции мы обе этих переменные упаковываем в интерфейс,
	// и эти интерфейсы != nil
}
