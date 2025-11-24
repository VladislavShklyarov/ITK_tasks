package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(SimpleError())
	err := MyError{Code: 418}
	fmt.Println(FormattedError(21))
	fmt.Println(StructError(err.Error()))
}

func SimpleError() error {
	return errors.New("простая ошибка")
}

func FormattedError(age int) error {
	basicError := fmt.Errorf("возраст %d недопустим", age)
	wrappedError := fmt.Errorf("обернутая ошибка: %w", basicError)
	return wrappedError
}

type MyError struct {
	Code int
}

func (e MyError) Error() int {
	return e.Code
}

func StructError(code int) error {
	return fmt.Errorf("MyError{Code: %d, Msg: \"%s\"}", code, "Не найдено")
}
