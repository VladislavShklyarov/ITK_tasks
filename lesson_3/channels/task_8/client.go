package main

import (
	"fmt"
	"log"
	"sync"
)

type Client struct {
	id      string
	servers []*Server
	rr      *RoundRobin
}

func NewClient(id string, servers []*Server) *Client {
	c := &Client{
		id:      id,
		servers: servers,
		rr:      NewRoundRobin(servers),
	}
	return c
}

func (c *Client) StartServers() {
	for _, s := range c.servers {
		s.Save()
	}
	log.Println("Сервера запущены и ждут данные")
}

func (c *Client) SendTo(server *Server, message *Message) {
	alteredBody := fmt.Sprintf("send from %s: \n\t%s", c.id, message.body)
	msg := &Message{
		id:   message.id,
		body: alteredBody,
	}
	server.receiver <- msg
}

func (c *Client) BroadcastData(message *Message) {
	alteredBody := fmt.Sprintf("broadcasted from %s: \n\t%s", c.id, message.body)
	msg := &Message{
		id:   message.id,
		body: alteredBody,
	}
	for _, s := range c.servers {
		s.receiver <- msg
	}
}

func (c *Client) SplitSend(in <-chan *Message) {
	for data := range in {
		server := c.rr.Next()
		c.SendTo(server, data)
	}
	c.BroadcastData(&Message{
		id:   NewID(),
		body: "end of splitting",
	})

}

func (c *Client) Receive() chan *Message {

	out := make(chan *Message)

	for _, server := range c.servers {
		server.commands <- Command{"export"}
	}

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(c.servers))

		for _, s := range c.servers {
			go func(server *Server) {
				defer wg.Done()
				for msg := range server.exporter {
					out <- msg
				}
			}(s)
		}

		wg.Wait()
		close(out)
	}()
	return out
}
