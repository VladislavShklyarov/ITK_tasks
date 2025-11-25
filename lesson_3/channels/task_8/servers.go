package main

import (
	"sort"
	"sync"
)

type Server struct {
	id       int
	receiver chan string
	vault    *Vault
}

type Vault struct {
	internalVault map[string]struct{}
	mu            sync.Mutex
}

func NewServer(id int) *Server {
	return &Server{
		id:       id,
		receiver: make(chan string),
		vault: &Vault{
			internalVault: make(map[string]struct{}),
			mu:            sync.Mutex{},
		},
	}
}

//
//type MultiError struct {
//	msgs []string
//}

//func (m *MultiError) Error() string {
//	return strings.Join(m.msgs, ";\n")
//}

// Горутина save аггрегирует значения из канала и сохраняет в in-memory хранилище.

func (s *Server) Save() {
	go func() {
		for data := range s.receiver {
			s.vault.mu.Lock()
			if _, ok := s.vault.internalVault[data]; !ok {
				s.vault.internalVault[data] = struct{}{}
			}
			s.vault.mu.Unlock()
		}
	}()
}

func (s *Server) GetData() []string {
	s.vault.mu.Lock()
	defer s.vault.mu.Unlock()

	data := make([]string, 0, len(s.vault.internalVault))
	for key := range s.vault.internalVault {
		data = append(data, key)
	}

	sort.Strings(data) // сортируем для красивого и предсказуемого вывода
	return data
}
