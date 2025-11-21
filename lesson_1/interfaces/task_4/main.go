package main

import "fmt"

type MyStruct struct {
	MyInt int
}

func func1() MyStruct {
	return MyStruct{MyInt: 1}
}

func func2() *MyStruct {
	return &MyStruct{}
}

func func3(s *MyStruct) {
	s.MyInt = 333
}

func func4(s MyStruct) {
	s.MyInt = 923
}

func func5() *MyStruct {
	return nil
}

func main() {
	ms1 := func1()         // Здесь создаем структуру
	fmt.Println(ms1.MyInt) // 1

	ms2 := func2()         // Создаем структуру, но возвращаем УКАЗАТЕЛЬ на нее
	fmt.Println(ms2.MyInt) // 0 т.к. значение не задано.

	func3(ms2)             // Функция принимает указатель на структуру, и изменяет ее значение
	fmt.Println(ms2.MyInt) // Из-за этого выведет 333

	func4(ms1) // Эта функция принимает саму структуру, поэтому создается ее КОПИЯ
	// И значение мы меняем непосредственно для копии, оригинальная ms1 остается без изменений
	fmt.Println(ms1.MyInt) // 1 как на строчке 31

	ms5 := func5()         // Возвращает пустой указатель.
	fmt.Println(ms5.MyInt) // Свалимся в панику, т.к. пытаемся обратиться к полю через nil. Go пытается разыменовать
	// и паникует
}
