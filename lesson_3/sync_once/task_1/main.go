package main

import (
	"fmt"
	"sync"
	pg "task_1/PostgreSQL"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ConnectionPool := pg.PostgresNewPool()
	//for i := 0; i < 50; i++ {
	//	wg.Go(func() {
	//		start := time.Now()
	//		conn := ConnectionPool.GetConnectionOnce()
	//		fmt.Printf("Горутина %d получила подключение %d за %v\n", i, conn.ID, time.Since(start))
	//	})
	//}
	//time.Sleep(3 * time.Second)
	//for i := 0; i < 50; i++ {
	//	start := time.Now()
	//	conn := ConnectionPool.GetConnectionOnce()
	//	fmt.Printf("Цикл %d получил подключение %d за %v\n", i, conn.ID, time.Since(start))
	//}
	//time.Sleep(3 * time.Second)
	for i := range 50 {
		wg.Go(func() {
			if i != 0 {
				time.Sleep(300 * time.Millisecond)
			}
			start := time.Now()
			conn := ConnectionPool.GetConnection() // Connection без sync.Once
			fmt.Printf("[UNSAFE]Горутина %d получила подключение %d за %v\n", i, conn.ID, time.Since(start))
		})
	}
	//wg.Wait()
	//// Далее примеры с sync.Cond()
	//time.Sleep(3 * time.Second)
	//ConnectionPoolCS := pg.PostgresNewPoolSC()
	//
	//for i := 0; i < 50; i++ {
	//	wg.Go(func() {
	//		start := time.Now()
	//		conn := ConnectionPoolCS.GetConnectionSC()
	//		fmt.Printf("Горутина %d получила подключение %d за %v\n", i, conn.ID, time.Since(start))
	//	})
	//}
	//time.Sleep(3 * time.Second)
	//for i := 0; i < 50; i++ {
	//	start := time.Now()
	//	conn := ConnectionPoolCS.GetConnectionSC()
	//	fmt.Printf("Цикл %d получил подключение %d за %v\n", i, conn.ID, time.Since(start))
	//}

	wg.Wait()
}
