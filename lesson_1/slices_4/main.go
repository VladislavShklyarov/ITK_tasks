package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1
	first := []int{1, 2, 3, 4, 5}
	first = nil
	fmt.Println("first:", first, ":", len(first), ":", cap(first))
	// first: [], 0, 0. Мы просто обнуляем слайс

	//2
	second := []int{1, 2, 3, 4, 5}
	second = second[:0]
	fmt.Println("second:", second, ":", len(second), ":", cap(second))
	// second: [], 0, 5. Когда берем срез, capacity сохраняется от элемента среза до конца слайса

	//3
	third := []int{1, 2, 3, 4, 5}
	clear(third)
	fmt.Println("third:", third, ":", len(third), ":", cap(third))
	// Clear зануляет элементы, но сохраняет len и cap
	// [0 0 0 0 0 ] 5 5

	//4
	fourth := []int{1, 2, 3, 4, 5}
	clear(fourth[1:3])
	fmt.Println("fourth:", fourth, ":", len(fourth), ":", cap(fourth))
	// Я думаю что занулим только 1 и 2 индексы
	// [1 0 0 4 5] 5 5

	//5
	slice := make([]int, 3, 6)
	array := [3]int(slice[:3])
	slice[0] = 10
	// массивы копируются при присваивании, так что это независящие друг от друга объекты
	fmt.Println("slice = ", slice, len(slice), cap(slice)) // [10 0 0] 3 6
	fmt.Println("array =", array, len(array), cap(array))  // [0 0 0 ] 3 3

	//6 В каких случаях Slice пустой или нулевой
	//1
	var data []string
	fmt.Println("var data []string:")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// empty и nil покажут true, т.к. слайс пустой и неинициализированный
	// Под капотом в слайсе лежит структура из 3-х элементов - указаель на массив, длина и capacity. На каждый по 8 байт
	// => SizeOf = 24. SliceData возвращает указатель на первый элемент слайса. Но поскольку слайс nil то и адрес будет пустой

	//2
	data = []string(nil)
	fmt.Println("data = []string(nil):")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// Аналогично предыдущему, только мы специально сделали слайс nil. true, true, 24 0x0
	//3
	data = []string{}
	fmt.Println("data = []string{}")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// Здесь уже слайс инициализирован, но базовый массив все еще пуст. То есть как бы структура уже есть, но нет данных
	// empty=true, nil=false (мы инициализировали слайс) размер 24, адрес уже не пустой
	//4
	data = make([]string, 0)
	fmt.Println("data =make([]string,0)")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// ситуация аналогичная предыдущей, только по-другому инициализировали слайс

	empty := struct{}{}
	fmt.Println("empty struct address ", unsafe.Pointer(&empty))
	// пустая структура не занимает места в памяти, но адрес у нее все равно есть) просто он виртуальный
}
