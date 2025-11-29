package main

import (
	"fmt"
	"sync"
	"task_2/LoadEnv"
	"time"
)

func main() {
	var wg sync.WaitGroup

	envLoader := LoadEnv.NewEnvConfigLoader()

	for id := range 10 {
		wg.Go(func() {
			fmt.Printf("Горутина %d вызывает LoadConfig...\n", id)
			envLoader.LoadConfig()
			fmt.Printf("Горутина %d завершила LoadConfig\n", id)
		})
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println("\n--- Проверяем вывод всех значений ---")
	for range 10 {
		envLoader.PrintConfig()
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println("\n--- Конкурентные запросы переменных среды ---")

	for id := range 10 {
		wg.Go(func() {
			appName := envLoader.Get("APP_NAME")
			port := envLoader.Get("APP_PORT")

			fmt.Printf("[Горутина %d] APP_NAME=%s, APP_PORT=%s\n", id, appName, port)
		})
	}
	wg.Wait()

}
