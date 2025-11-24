package main

import (
	"fmt"
	"reflect"
)

type MyError struct {
	data string
}

func (m *MyError) Error() string {
	return m.data
}
func foo(i int) error {
	var err *MyError
	if i > 5 {
		err = &MyError{data: "i>5"}
	}
	return err
}
func main() {
	err := foo(4)
	fmt.Println(err, reflect.TypeOf(err))
	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok")
	}
}

// Вернувшаяся ошибка не равна nil, потому что MyError реализует интерфейс Error.
// Когда мы возвращаем err как error, происходит упаковка в интерфейс.
// Интерфейсы состоят из двух itab и data. Первое - тип, второе - конкретные методы. У нас уже есть конкретный тип: структура MyError
// Если хотя бы одно из этих значений != nil, То и весь интерфейс != nil
//
