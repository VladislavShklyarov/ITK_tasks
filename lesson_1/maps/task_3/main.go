package main

import (
	"fmt"
)

func FilterByValue(m map[int]string, allowedValues []string) map[int]string {

	allowedSet := make(map[string]struct{})

	for _, value := range allowedValues {
		allowedSet[value] = struct{}{}
	}

	correctMap := make(map[int]string)

	for key, value := range m {
		if _, ok := allowedSet[value]; ok {
			correctMap[key] = value
		}
	}
	return correctMap
}

func InvertMap(m map[string]int) (map[int]string, error) {
	invertedMap := make(map[int]string)

	for key, value := range m {
		if existing, exist := invertedMap[value]; exist {
			return nil, fmt.Errorf("такой элемент уже есть: " + existing + " и " + key)
		}
		invertedMap[value] = key
	}

	return invertedMap, nil
}

func main() {
	originalMap := map[int]string{
		1: "apple",
		2: "banana",
		3: "orange",
		4: "grape",
		5: "kiwi",
	}

	allowed := []string{"apple", "orange", "kiwi"}

	filtered := FilterByValue(originalMap, allowed)

	fmt.Println("Исходная мапа:", originalMap)
	fmt.Println("Разрешенные значения:", allowed)
	fmt.Println("Отфильтрованная мапа:", filtered)

	mapToInvert := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
		"grape":  4,
		"kiwi":   5,
	}

	inverted, err := InvertMap(mapToInvert)
	fmt.Println("Исходная мапа:", mapToInvert)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Перевернутая мапа:", inverted)
	}
}
