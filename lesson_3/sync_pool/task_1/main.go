package main

import (
	"fmt"
	"strings"
	"task_1/syncPool"
)

func main() {
	examples := []string{
		"hello, world!",
		"gopher",
		"lorem ipsum dolor sit amet",
	}

	pool := &syncPool.SyncPool{}

	for _, s := range examples {
		processed := ProcessString(s)
		fmt.Printf("Original: %q\nProcessed: %q\n\n", s, processed)
	}
}

func ProcessString(s string) string {
	buf := syncPool.BuffersPool.Get().(*[]byte) // берем буфер из пула
	*buf = append(*buf, strings.ToUpper(s)...)  // добавляем в него обработанную строку
	result := string(*buf)                      // сохраняем в локальной переменной

	syncPool.BuffersPool.(buf) // возвращаем буфер обратно в пул
	return result
}
