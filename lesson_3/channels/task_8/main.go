package main

import (
	"fmt"
)

func main() {
	sourceCh := FillIn()
	parsedCh := Parse(sourceCh)

	s1 := NewServer("alpha")
	s2 := NewServer("beta")
	s3 := NewServer("delta")

	ss := []*Server{s1, s2, s3}

	client := NewClient("client 1", ss)

	client.StartServers() // начинаем слушать каналы receivers

	client.SplitSend(parsedCh)

	messagesCh := client.Receive()
	fmt.Println("=== Первичная выгрузка сообщений ===")
	for msg := range messagesCh {
		fmt.Println(msg.id, "->", msg.body)
	}

	client.BroadcastData(&Message{
		id:   NewID(),
		body: "Выгрузка завершена, всем спасибо",
	})

	anotherMessagesCh := client.Receive()
	fmt.Println("=== Первичная выгрузка сообщений ===")
	for msg := range anotherMessagesCh {
		fmt.Println(msg.id, "->", msg.body)

	}
}
