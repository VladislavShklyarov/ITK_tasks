package main

import (
	"fmt"
	"time"
)

func main() {
	cache := NewCache[any]()

	// Устанавливаем элементы с разными TTL
	cache.Set("user", User{Name: "Alice"}, time.Hour) // Хранится 1 час
	cache.Set("temp_data", 42, time.Second)           // Хранится 1 секунду

	// Ждем 2 секунды, чтобы temp_data истек
	time.Sleep(2 * time.Second)

	// Проверка получения элементов
	val, ok := cache.Get("user")
	if ok {
		fmt.Println("user:", val)
	}

	val, ok = cache.Get("temp_data") // Должно показать, что TTL истек
	if !ok {
		fmt.Println("temp_data уже удалился")
	}

	// Принудительное удаление
	cache.Delete("user")
	val, ok = cache.Get("user")
	if !ok {
		fmt.Println("user удален")
	}

	// Добавляем новые элементы
	cache.Set("user_1", User{Name: "Mark"}, time.Hour)
	cache.Set("user_2", User{Name: "Bill"}, 30*time.Minute)

	// Проверка существования
	fmt.Println("user_1 exists?", cache.Exists("user_1"))       // true
	fmt.Println("temp_data exists?", cache.Exists("temp_data")) // false

	// Сериализация кэша в JSON
	jsonData, _ := cache.ToJSON()
	fmt.Println("Кэш в JSON:", string(jsonData))

	// Типизированное получение
	valAs, err := cache.GetAs("user_2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("user_2 как конкретный тип: %#v\n", valAs)
	}

	// Очистка кэша
	cache.Clear()
	jsonData, _ = cache.ToJSON()
	fmt.Println("Кэш после очистки:", string(jsonData))

	// Добавляем новый элемент
	cache.Set("temp_data", 42, time.Second)
	val, ok = cache.Get("temp_data")
	if ok {
		fmt.Println("temp_data после очистки:", val)
	}
}
