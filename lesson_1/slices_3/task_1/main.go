package main

import "fmt"

// RemoveUnordered удаляет элемент по индексу без сохранения порядка.
// Если индекс выходит за границы слайса, возвращает исходный слайс.
func RemoveUnordered[T any](s []T, i int) []T {
	/* Чтобы не сдвигать все элементы, мы просто перезаписываем последний элемент
	на место i и обрубаем хвост.
	*/
	if i > len(s) {
		return s
	}
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]
	return s
}

// RemoveOrdered удаляет элемент по индексу с сохранением порядка.
// Если индекс выходит за границы слайса, возвращает исходный слайс.
func RemoveOrdered[T any](s []T, i int) []T {
	/*
		Для этого можно элементам начиная с искомого присвоить срез от искомого + 1 до конца
	*/
	s = append(s[:i], s[i+1:]...)

	return s
}

// RemoveAllByValue удаляет все вхождения указанного значения.
func RemoveAllByValue[T comparable](s []T, value T, ordered bool) []T {
	/*
			Для этого нужно
		1. найти индексы всех элементов, совпадающих с искомым
		2. вызвать функцию RemoveOrdered для каждого из них. Но если мы будем так делать, то индексы
		будут каждый раз меняться, соответственно их снова придется пересчитывать => O(n^2)

		Альтернатива

		Можно не искать элементы, а сразу двигаться по списку. Если элемент найден:
		Выполняем RemoveUnordered. Только нужно явно задавать i внутри этой функции,
		потому что RemoveUnordered меняет значения цикла
	*/

	/*
			 Если же нужно сохранить порядок и иметь при этом адекватную сложность,
		можно просто добавлять элементы в новый слайс.
	*/

	if ordered {
		res := make([]T, 0, len(s))
		for _, el := range s {
			if el != value {
				res = append(res, el)
			}
		}
		return res
	}

	for i := 0; i < len(s); {
		if s[i] == value {
			s = RemoveUnordered(s, i)
		} else {
			i++
		}
	}
	return s

}

// RemoveDuplicates оставляет только уникальные элементы (сохраняет порядок).
func RemoveDuplicates[T comparable](s []T, addNewSlice bool) []T {
	// для этого нужно создать мапу уникальности. затем сравнивать элементы из списка с мапой и
	// и добавлять в слайс
	seen := make(map[T]struct{})

	if addNewSlice {
		result := make([]T, 0, len(s))
		for _, el := range s {
			if _, ok := seen[el]; !ok {
				result = append(result, el)
				seen[el] = struct{}{}
			}
		}
		return result
	}

	writeIndex := 0
	// Либо, если нужно вернуть тот же массив,
	// можно использовать паттерн "указатель" для перезаписи элементов

	for _, el := range s {
		if _, ok := seen[el]; !ok {
			seen[el] = struct{}{}
			s[writeIndex] = el
			writeIndex++
		}
	}
	return s[:writeIndex]
}

// RemoveIf удаляет элементы, удовлетворяющие условию predicate.
func RemoveIf[T any](s []T, predicate func(T) bool, addNewSlice bool) []T {
	//  можно проверить в цикле удовлетворение условию и сложить в новый слайс
	if addNewSlice {
		result := make([]T, 0, len(s))
		for _, el := range s {
			if !predicate(el) {
				result = append(result, el)
			}
		}
		return result
	}
	// либо как и в предыдущем использовать указатель для перезаписи
	writeIndex := 0
	for _, el := range s {
		if !predicate(el) {
			s[writeIndex] = el
			writeIndex++
		}
	}
	return s[:writeIndex]
}

// RemoveOrderedWithNil удаляет элемент по индексу (для слайса указателей),
// обнуляя удаляемый элемент для предотвращения утечек памяти.
func RemoveOrderedWithNil[T any](s []*T, i int) []*T {
	//если правильно понял, то нам нужно просто сделать указатель = nil, а в остальном по классике

	s[i] = nil
	s = append(s[:i], s[i+1:]...)
	return s
}

// ShrinkCapacity сокращает вместимость слайса, если она превышает
// удвоенную длину после удаления элементов.
func ShrinkCapacity[T any](s []T) []T {
	// можно решить копированием в слайс со строго заданными размерами
	if cap(s) > len(s)*2 {
		newSlice := make([]T, len(s))
		copy(newSlice, s)
		return newSlice
	}

	return s
}

func main() {
	a := []int{10, 20, 30, 40, 50}
	a = RemoveUnordered(a, 1)

	fmt.Println(a)

	b := []int{10, 20, 30, 40}
	b = RemoveOrdered(b, 2) // удаляем 30

	fmt.Println(b)

	c := []int{1, 2, 1, 3, 1, 4, 1}

	// Удаление с сохранением порядка
	result1 := RemoveAllByValue(c, 1, true)
	fmt.Println(result1)

	// Удаление без сохранения порядка
	result2 := RemoveAllByValue(c, 1, false)
	fmt.Println(result2)

	d := []int{5, 5, 3, 5, 2, 3, 2, 1}

	// Создаём новый слайс
	r1 := RemoveDuplicates(d, true)
	fmt.Println(r1)

	// Удаляем "на месте"
	r2 := RemoveDuplicates(d, false)
	fmt.Println(r2)

	e := []int{1, 2, 3, 4, 5, 6}

	// Удаляем чётные
	res1 := RemoveIf(e, func(x int) bool {
		return x%2 == 0
	}, true)
	fmt.Println(res1)

	f, g, h := 10, 20, 30
	pointers := []*int{&f, &g, &h}

	pointers = RemoveOrderedWithNil(pointers, 1) // удаляем b
	fmt.Println(len(pointers))
	fmt.Println(pointers)

	bigCap := make([]int, 5, 200)

	bigCap = ShrinkCapacity(bigCap)
	fmt.Println(len(bigCap), cap(bigCap))
}
