package main

import "github.com/google/uuid"

func FillIn() <-chan string {
	out := make(chan string)
	go func() {
		for _, log := range appLogs {
			out <- log
		}
		close(out)
	}()
	return out
}

var appLogs = []string{
	"2024-01-15 08:50:30 INFO - Database connections closed",
	"2024-01-15 08:50:35 INFO - Cache connections closed",
	"2024-01-15 08:50:40 INFO - Application shutdown completed successfully",
}

type Message struct {
	id   string
	body string
}

func NewID() string {
	return uuid.New().String()
}

func Parse(in <-chan string) chan *Message {

	out := make(chan *Message)
	go func() {
		for value := range in {
			out <- &Message{
				id:   NewID(),
				body: value,
			}
		}
		close(out)
	}()

	return out
}
