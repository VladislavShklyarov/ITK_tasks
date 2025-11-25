package main

import "fmt"

func main() {

	s1 := NewServer(1)
	s2 := NewServer(2)
	s3 := NewServer(3)

	servers := []*Server{s1, s2, s3}

	dataCh := FillIn()

	parseCh := Parse(dataCh)
	serversSplit := Split(parseCh, servers)
	Send(serversSplit)

	for i, server := range serversSplit {
		fmt.Printf("Server %d:\n", i+1)
		data := server.GetData()
		if len(data) == 0 {
			fmt.Println("  (пусто)")
			continue
		}
		for _, info := range data {
			fmt.Printf("  - %s\n", info)
		}
	}

}

func Parse(in <-chan string) chan string {

	out := make(chan string)
	go func() {
		for value := range in {
			out <- "parsed - " + value
		}
		close(out)
	}()

	return out
}

func Split(in <-chan string, servers []*Server) []*Server {
	rr := NewRoundRobin(servers)

	for _, server := range servers {
		server.Save() // запускаем горутину, которая читает и сохраняет данные
	}

	for data := range in {
		server := rr.Next()
		server.receiver <- data
	}

	//for _, server := range servers {
	//	close(server.receiver)
	//}

	return servers
}

func Send(servers []*Server) {
	for _, server := range servers {
		server.Save() // запускаем горутину, которая читает и сохраняет данные
		server.receiver <- "Данные записанные методом Send"
	}
}

/* TODO: записи из мапы вернуть обратно в канал reciever (либо создать новый канал)
Из этого канала мы получаем записи в Send, дополняем каждую из них и снова сохраняем

*/
