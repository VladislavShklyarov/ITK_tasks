package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const NumWorkers int = 4

func main() {
	timeStart := time.Now()
	fileNamesCh := make(chan string)
	resultCh := make(chan int)
	result := 0
	wg := sync.WaitGroup{}

	files, err := filepath.Glob("*.txt")
	if err != nil {
		fmt.Printf("Ошибка: %s", err)
		return
	}

	go func() { // Заполняем канал именами файлов
		for _, file := range files {
			fileNamesCh <- file
		}
		close(fileNamesCh)
	}()

	for i := 0; i < NumWorkers; i++ { // Вызываем воркеров на каждый файл.
		wg.Go(func() {
			for file := range fileNamesCh {
				count, err := processFile(file)
				if err != nil {
					fmt.Printf("Ошибка: %s\n", err)
					continue
				}
				resultCh <- count
			}
		})
	}

	go func() {
		wg.Wait()       // ждём всех воркеров
		close(resultCh) // после завершения закрываем канал
	}()

	for count := range resultCh {
		result += count
	}

	fmt.Println("Общее количество слов:", result)
	fmt.Printf("%v", time.Since(timeStart))
}

func processFile(filename string) (int, error) {
	time.Sleep(100 * time.Millisecond) // предположим что файлы большие, и на обработку каждого нужно 100ms
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Ошибка: %s", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		total += len(words)
	}
	return total, nil
}
