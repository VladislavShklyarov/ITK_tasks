package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

var (
	ErrNotFound  = errors.New("ресурс не найден")
	TimeoutError = errors.New("таймаут операции")
	UnknownError = errors.New("неизвестная ошибка")
)

func main() {
	wg := sync.WaitGroup{}
	for range 10 {
		wg.Go(func() {
			ProcessError(SimulateRequest())
		})
	}
	wg.Wait()
}

func ProcessError(err error) {
	if errors.Is(err, TimeoutError) {
		fmt.Println("Требуется повторная попытка")
	} else if errors.Is(err, ErrNotFound) {
		fmt.Println("Ресурс не найден")
	} else {
		fmt.Println("Неизвестная ошибка")
	}
}

func SimulateRequest() error {
	chance := rand.Intn(10)
	switch {
	case chance >= 5:
		return fmt.Errorf("запрос не выполнен: ,%w", TimeoutError)
	case chance >= 2:
		return fmt.Errorf("ошибка: ,%w", ErrNotFound)
	default:
		return UnknownError
	}
}
