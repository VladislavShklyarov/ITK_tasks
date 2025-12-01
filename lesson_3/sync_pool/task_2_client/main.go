package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	url := "http://localhost:8080/CreatePool"

	// Имя—возраст
	people := map[string]string{
		"Alice":   "25",
		"Bob":     "30",
		"Charlie": "22",
		"Diana":   "28",
		"Eve":     "35",
		"Frank":   "40",
		"Grace":   "27",
		"Hank":    "32",
		"Ivy":     "21",
		"Jack":    "29",
		"Kevin":   "33",
		"Laura":   "26",
		"Mike":    "31",
		"Nina":    "24",
		"Oscar":   "36",
		"Paula":   "29",
		"Quinn":   "23",
		"Rachel":  "34",
		"Steve":   "37",
		"Tina":    "25",
		"Uma":     "28",
		"Victor":  "32",
		"Wendy":   "27",
		"Xander":  "30",
		"Yara":    "22",
		"Zane":    "35",
		"Abby":    "24",
		"Ben":     "33",
		"Clara":   "26",
		"David":   "31",
		"Ella":    "29",
		"Finn":    "28",
		"Gina":    "27",
		"Carlos":  "34",
		"Isla":    "23",
		"Jake":    "36",
		"Kara":    "25",
		"Liam":    "32",
		"Mia":     "22",
		"Noah":    "30",
		"Olivia":  "29",
	}

	var wg sync.WaitGroup
	i := 0

	for name, age := range people {
		wg.Add(1)
		go func(i int, name, age string) {
			defer wg.Done()

			// Формируем JSON: {"data": {"Имя": "Возраст"}}
			body := map[string]map[string]string{
				"data": {name: age},
			}

			jsonBytes, _ := json.Marshal(body)

			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
			if err != nil {
				fmt.Printf("Request %d (%s) error: %v\n", i, name, err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Request %d (%s) status: %s\n", i, name, resp.Status)

		}(i, name, age)

		i++
	}

	wg.Wait()
	fmt.Println("All requests done")
}
