package main

type RoundRobin struct {
	servers []*Server
	current int
}

func NewRoundRobin(servers []*Server) *RoundRobin {
	return &RoundRobin{
		servers: servers,
		current: 0,
	}
}

func (rr *RoundRobin) Next() *Server {
	if len(rr.servers) == 0 {
		return nil
	}

	server := rr.servers[rr.current]                // выбираем текущий сервис.
	rr.current = (rr.current + 1) % len(rr.servers) // таким образом для последнего сервера следующим будет первый
	return server
}
