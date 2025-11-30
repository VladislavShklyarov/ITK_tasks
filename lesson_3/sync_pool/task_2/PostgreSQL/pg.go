package PostgreSQL

import (
	"fmt"
	"sync"
	"time"
)

type Postgres struct {
	data          *Data
	name          string
	connections   []*Connection
	maxConnAmount int
	cond          *sync.Cond
}

type Connection struct {
	ID   int
	name string
}

func (pg *Postgres) queueIsEmpty() bool {
	return len(pg.connections) == 0
}
func (pg *Postgres) queueIsFull() bool {
	return len(pg.connections) == pg.maxConnAmount
}

type Data struct {
	// Значениями будут имя пользователя и возраст
	vault map[string]string
	mu    sync.Mutex
}

func NewPostgresPool(maxConnectionsAmount int) *Postgres {
	pool := &Postgres{
		name: "Redis",
		data: &Data{
			vault: make(map[string]string),
			mu:    sync.Mutex{},
		},
		maxConnAmount: maxConnectionsAmount,
		connections:   make([]*Connection, 0, maxConnectionsAmount),
		cond:          sync.NewCond(&sync.Mutex{}),
	}

	for i := range maxConnectionsAmount {
		pool.connections = append(pool.connections, &Connection{ID: i})
	}
	return pool
}

func (pg *Postgres) GetConn() *Connection {
	pg.cond.L.Lock()
	for pg.queueIsEmpty() {
		fmt.Println("All connections are busy, waiting...")
		pg.cond.Wait()
	}

	conn := pg.connections[0]
	pg.connections = pg.connections[1:]

	fmt.Printf("Connection %d acquired\n", conn.ID)
	pg.cond.L.Unlock()
	time.Sleep(100 * time.Millisecond) // имитация задержки на выдачу подключения к БД
	return conn
}

func (pg *Postgres) Release(connection *Connection) {
	pg.cond.L.Lock()
	for pg.queueIsFull() {
		pg.cond.Wait()
	}

	pg.connections = append(pg.connections, connection)
	fmt.Printf("Connection %d released\n", connection.ID)
	pg.cond.Signal()
	pg.cond.L.Unlock()
}
