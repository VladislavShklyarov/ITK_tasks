package sqlite

import (
	"math/rand"
	"sync"
	c "task_3/connector"
	"time"
)

type SQLite struct {
	name        string
	data        map[string]struct{} // Обращаться к ней не будем в коде
	MaxConnNum  int
	connections []*c.Connection
	cond        *sync.Cond
}

func (s *SQLite) queueIsEmpty() bool {
	return len(s.connections) == 0
}

func (s *SQLite) queueIsFull() bool {
	return len(s.connections) == s.MaxConnNum
}

func NewSQLiteConnectionPool(maxConenctions int) *SQLite {
	pool := &SQLite{
		name:        "SQLite conncetor pool",
		data:        nil,
		MaxConnNum:  maxConenctions,
		connections: make([]*c.Connection, 0, maxConenctions),
		cond:        sync.NewCond(&sync.Mutex{}),
	}

	for i := 1; i <= maxConenctions; i++ {
		pool.connections = append(pool.connections, &c.Connection{ID: i})
	}
	return pool
}

func (sql *SQLite) Get() *c.Connection {
	sql.cond.L.Lock()

	for len(sql.connections) == 0 {
		sql.cond.Wait()
	}

	conn := sql.connections[0]
	sql.connections = sql.connections[1:]
	sql.cond.Signal()
	sql.cond.L.Unlock()
	waitTime := time.Duration(rand.Intn(500)+100) * time.Millisecond // случайно от 100 до 600 мс
	time.Sleep(waitTime)
	return conn
}

func (sql *SQLite) Release(connection *c.Connection) {
	sql.cond.L.Lock()
	for sql.queueIsFull() {
		sql.cond.Wait()
	}
	sql.connections = append(sql.connections, connection)
	sql.cond.Signal() // Сигнализируем другим горутинам что подключение освободилось
	sql.cond.L.Unlock()
}
