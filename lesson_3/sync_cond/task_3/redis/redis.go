package redis

import (
	"math/rand"
	"sync"
	c "task_3/connector"
	"time"
)

type Redis struct {
	name          string
	cash          *Cash // А вот сюда пожалуй и запишем что-нибудь
	maxConnAmount int
	connections   []*c.Connection
	cond          *sync.Cond
}

type Cash struct {
	// Значениями пусть будут типа ID пользователя и каки-то кэшированные строки
	vault map[string][]string
	mu    sync.Mutex
}

func (r *Redis) queueIsEmpty() bool {
	return len(r.connections) == 0
}
func (r *Redis) queueIsFull() bool {
	return len(r.connections) == r.maxConnAmount
}

func NewRedisPool(maxConnectionsAmount int) *Redis {
	pool := &Redis{
		name: "Redis",
		cash: &Cash{
			vault: make(map[string][]string),
			mu:    sync.Mutex{},
		},
		maxConnAmount: maxConnectionsAmount,
		connections:   make([]*c.Connection, 0, maxConnectionsAmount),
		cond:          sync.NewCond(&sync.Mutex{}),
	}

	for i := range maxConnectionsAmount {
		pool.connections = append(pool.connections, &c.Connection{ID: i})
	}
	return pool
}

func (r *Redis) Get() *c.Connection {
	r.cond.L.Lock()
	for r.queueIsEmpty() {
		r.cond.Wait()
	}

	conn := r.connections[0]
	r.connections = r.connections[1:]
	r.cond.Signal()
	r.cond.L.Unlock()
	time.Sleep(300 * time.Millisecond) // время на подключение к Redis
	return conn
}

func (r *Redis) Release(connection *c.Connection) {
	r.cond.L.Lock()
	for r.queueIsFull() {
		r.cond.Wait()
	}

	r.connections = append(r.connections, connection)
	r.cond.Signal()
	r.cond.L.Unlock()
}

// WriteData по началу работала строго последовательно. Никак не мог понять,
// почему у меня несколько коннектов и воркеров, а пишется все с разницей 500 ms.
// В итоге обнаружил, что из-за defer на 81 строке лочится вообще вся мапа,
// и другие воркеры не могут получить к ней доступ. Так что defer не всегда гуд

func (r *Redis) WriteData(key string, values []string) {
	r.cash.mu.Lock()
	// defer r.cash.mu.Unlock()
	if _, ok := r.cash.vault[key]; !ok {
		r.cash.vault[key] = values
	}
	r.cash.mu.Unlock()
	waitTime := time.Duration(rand.Intn(1000)+50) * time.Millisecond // случайно от 100 до 600 мс
	time.Sleep(waitTime)
}

func (r *Redis) GetData() chan map[string][]string {
	out := make(chan map[string][]string)

	go func() {
		r.cash.mu.Lock()
		defer r.cash.mu.Unlock()
		defer close(out)
		for key, values := range r.cash.vault {
			outMap := map[string][]string{key: values}
			out <- outMap
		}
	}()

	return out
}
