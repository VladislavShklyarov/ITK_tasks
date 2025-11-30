package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	url := "http://localhost:8080/CreatePool"

	// JSON, который будем отправлять
	payloads := []string{
		`{"name":"Alice","age":"25"}`,
		`{"name":"Bob","age":"30"}`,
		`{"name":"Charlie","age":"22"}`,
		`{"name":"Diana","age":"28"}`,
		`{"name":"Eve","age":"35"}`,
		`{"name":"Frank","age":"40"}`,
		`{"name":"Grace","age":"27"}`,
		`{"name":"Hank","age":"32"}`,
		`{"name":"Ivy","age":"21"}`,
		`{"name":"Jack","age":"29"}`,
		`{"name":"Kevin","age":"33"}`,
		`{"name":"Laura","age":"26"}`,
		`{"name":"Mike","age":"31"}`,
		`{"name":"Nina","age":"24"}`,
		`{"name":"Oscar","age":"36"}`,
		`{"name":"Paula","age":"29"}`,
		`{"name":"Quinn","age":"23"}`,
		`{"name":"Rachel","age":"34"}`,
		`{"name":"Steve","age":"37"}`,
		`{"name":"Tina","age":"25"}`,
		`{"name":"Uma","age":"28"}`,
		`{"name":"Victor","age":"32"}`,
		`{"name":"Wendy","age":"27"}`,
		`{"name":"Xander","age":"30"}`,
		`{"name":"Yara","age":"22"}`,
		`{"name":"Zane","age":"35"}`,
		`{"name":"Abby","age":"24"}`,
		`{"name":"Ben","age":"33"}`,
		`{"name":"Clara","age":"26"}`,
		`{"name":"David","age":"31"}`,
		`{"name":"Ella","age":"29"}`,
		`{"name":"Finn","age":"28"}`,
		`{"name":"Gina","age":"27"}`,
		`{"name":"Harry","age":"34"}`,
		`{"name":"Isla","age":"23"}`,
		`{"name":"Jake","age":"36"}`,
		`{"name":"Kara","age":"25"}`,
		`{"name":"Liam","age":"32"}`,
		`{"name":"Mia","age":"22"}`,
		`{"name":"Noah","age":"30"}`,
		`{"name":"Olivia","age":"29"}`,
	}

	var wg sync.WaitGroup
	for i, payload := range payloads {
		wg.Add(1)
		go func(i int, data string) {
			defer wg.Done()
			resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
			if err != nil {
				fmt.Printf("Request %d error: %v\n", i, err)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("Request %d status: %s\n", i, resp.Status)
		}(i, payload)
	}

	wg.Wait()
	fmt.Println("All requests done")
}
