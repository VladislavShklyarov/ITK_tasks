package main

import (
	"sync"
)

type Server struct {
	id       string
	receiver chan *Message
	exporter chan *Message
	vault    *Vault
	wg       sync.WaitGroup

	commands chan Command
}

type Command struct {
	cmdType string
}

type Vault struct {
	internalVault map[string]string
	mu            sync.Mutex
}

func NewServer(id string) *Server {
	return &Server{
		id:       id,
		receiver: make(chan *Message),
		exporter: make(chan *Message),
		vault: &Vault{
			internalVault: map[string]string{
				"initial": "this is automatically created message",
			},
			mu: sync.Mutex{},
		},
		commands: make(chan Command),
	}
}

// Save Горутина save аггрегирует значения из канала и сохраняет в in-memory хранилище.
func (s *Server) Save() {

	go func() {
		for {
			select {
			case msg := <-s.receiver:
				if msg == nil { // канал может быть закрыт
					return
				}
				s.vault.mu.Lock()
				s.vault.internalVault[msg.id] = msg.body
				s.vault.mu.Unlock()

			case cmd := <-s.commands:
				if cmd.cmdType == "export" {
					s.exportOnce()
				}
			}
		}
	}()

}

func (s *Server) exportOnce() {
	s.vault.mu.Lock()

	for key, value := range s.vault.internalVault {
		message := &Message{
			id:   key,
			body: value,
		}
		s.exporter <- message
	}
	close(s.exporter)
}
