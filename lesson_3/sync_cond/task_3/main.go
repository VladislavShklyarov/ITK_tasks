package main

import (
	"fmt"
	"os"
	"os/signal"
	"task_3/redis"
	"time"
)

func main() {

	//SQLitePool := sqlite.NewSQLiteConnectionPool(3) // Пул на 3 подключения

	//for i := 0; i < 50; i++ {
	//	go func() {
	//		conn := SQLitePool.Get()
	//		defer SQLitePool.Release(conn)
	//
	//		fmt.Printf("Горутина %d: SQLite подключение №%d получено\n", i, conn.ID)
	//		workTime := time.Duration(rand.Intn(100)+50) * time.Millisecond
	//		fmt.Printf("\tГорутина %d работала %v\n", i, workTime)
	//		time.Sleep(workTime)
	//	}()
	//}

	// Допустим сюда будут приходить мапы с id пользователей и каким-то кэшом.
	input := FillIn("cash.json")
	redisPool := redis.NewRedisPool(5)

	for i := 0; i < 5; i++ {

		go func() {

			conn := redisPool.Get()
			defer redisPool.Release(conn)
			for el := range input {
				start := time.Now()
				for key, values := range el {
					redisPool.WriteData(key, values)
				}

				fmt.Printf(
					"[Worker %d] использует RedisConn #%d | Записал ключи: %v за %v\n",
					i,
					conn.ID,
					getKeys(el),
					time.Since(start),
				)
			}
		}()
	}
	// записали в redis, теперь вычитаем значения так же через канал

	output := make(chan map[string][]string)

	for j := range 3 {
		go func() {
			for mapValue := range output {
				for key, value := range mapValue {
					fmt.Printf("Горутина %d: %s:%s\n", j, key, value)
				}
			}
		}()
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

}
