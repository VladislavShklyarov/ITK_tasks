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

	// или пройтись в цикле и переложить элементы за исключением нужного в другой слайс
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
func RemoveDuplicates[T comparable](s []T) []T {
	// для этого нужно создать мапу уникальности. затем сравнивать элементы из списка с мапой и
	// и добавлять в слайс

	seen := make(map[T]struct{})
	result := make([]T, 0, len(s))
	for _, el := range s {
		if _, ok := seen[el]; !ok {
			result = append(result, el)
			seen[el] = struct{}{}
		}
	}
	return result

}

// RemoveIf удаляет элементы, удовлетворяющие условию predicate.
func RemoveIf[T any](s []T, predicate func(T) bool) []T {
	// также можно проверить в цикле удовлетворение условию и сложить в новый слайс
	return s
}

// RemoveOrderedWithNil удаляет элемент по индексу (для слайса указателей),
// обнуляя удаляемый элемент для предотвращения утечек памяти.
func RemoveOrderedWithNil[T any](s []*T, i int) []*T {
	//реализовать
	return s
}

// ShrinkCapacity сокращает вместимость слайса, если она превышает
// удвоенную длину после удаления элементов.
func ShrinkCapacity[T any](s []T) []T {
	//реализовать
	return s
}

func main() {
	s := []int{4, 2, 3, 3, 2, 1, 4, 5, 1, 4}
	s = RemoveDuplicates(s)

	fmt.Println(s)
}
