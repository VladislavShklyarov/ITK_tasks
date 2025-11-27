package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type Entry struct {
	Name string   `json:"name"`
	Info []string `json:"info"`
}

func FillIn(path string) chan map[string][]string {
	ch := make(chan map[string][]string)

	go func() {
		defer close(ch)

		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}
		defer file.Close()

		// читаем JSON
		var entries []Entry
		if err = json.NewDecoder(file).Decode(&entries); err != nil {
			fmt.Println("Ошибка чтения JSON:", err)
			return
		}

		for _, e := range entries {
			ch <- map[string][]string{
				e.Name: e.Info,
			}
		}
	}()
	return ch
}
