package main

import (
	"errors"
	"fmt"
)

func main() {
	println("Case 1")
	case1()
	println()
	println()

	println("Case 2")
	case2()
	println()
	println()

	println("Case 3")
	case3()
	println()
	println()

}

func case1() {
	helperWithDefer := func(isError bool) error {
		var retVal error

		defer func() {
			retVal = errors.New("Extra error")
		}()

		if isError {
			retVal = errors.New("Default error")
		}

		return retVal
	}

	helperWithoutDefer := func(isError bool) error {
		var retVal error

		if isError {
			retVal = errors.New("Default error")
		}

		return retVal
	}

	fmt.Println("\twithout:")
	fmt.Println(helperWithoutDefer(false)) // nil
	fmt.Println(helperWithoutDefer(true))  // Default error
	fmt.Println("\twith:")
	fmt.Println(helperWithDefer(false)) // nil
	fmt.Println(helperWithDefer(true))  // Default error

	/* Если функция возвращает неименованный результат, то по окончании работы функция вычисляет итоговое значение
	возвращаемой переменной и складывает его в return slot.

	После этого выполняется defer, который локально меняет restVal, но это не изменяет значение в return-slot,
	так как там уже лежит "Default error".

	Соответственно выводится дефолтное значение
	*/
}

func case2() {
	helperWithDefer := func(isError bool) (retVal error) {
		defer func() {
			retVal = errors.New("Extra error")
		}()

		if isError {
			retVal = errors.New("Default error")
		}

		return
	}

	helperWithoutDefer := func(isError bool) (retVal error) {
		if isError {
			retVal = errors.New("Default error")
		}

		return
	}

	fmt.Println("\twithout:")
	fmt.Println(helperWithoutDefer(false)) // nil
	fmt.Println(helperWithoutDefer(true))  // Default error
	fmt.Println("\twith:")
	fmt.Println(helperWithDefer(false)) // Extra error
	fmt.Println(helperWithDefer(true))  // Extra error
	/* В случае, если мы возвращаем именованное значение, то она и является областью возврата,
	поэтому, если defer туда что-то запишет перед выводом, то это значение и вернется
	*/
}

func case3() {
	helperWithDefer := func(isError bool) (retVal error) {
		defer func() {
			retVal = errors.New("First Error")
		}()

		defer func() {
			retVal = errors.New("Second Error")
		}()

		if isError {
			retVal = errors.New("Default error")
		}

		return
	}

	helperWithoutDefer := func(isError bool) (retVal error) {
		if isError {
			retVal = errors.New("Default error")
		}

		return
	}

	fmt.Println("\twithout:")
	fmt.Println(helperWithoutDefer(false)) // nil
	fmt.Println(helperWithoutDefer(true))  // Default error
	fmt.Println("\twith:")
	fmt.Println(helperWithDefer(false)) // First error
	fmt.Println(helperWithDefer(true))  // First error

	// Ситуация аналогична предыдущей, только теперь вызывается несколько defer. Вызываются они "снизу вверх",
	// Поэтому в области возврата остается последнее значение.
}
