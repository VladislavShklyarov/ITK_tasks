package PostgreSQL

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type PostgresSQLConnPool struct {
	name string
	conn *Connection
	once sync.Once
}

type Connection struct {
	ID int
}

func PostgresNewPool() *PostgresSQLConnPool {
	return &PostgresSQLConnPool{
		name: "PostgreSQL",
		conn: nil,
		once: sync.Once{},
	}
}

func (pg *PostgresSQLConnPool) GetConnectionOnce() *Connection {
	pg.once.Do(func() {
		fmt.Println("Подключение инициализировано")
		time.Sleep(50 * time.Millisecond)
		pg.conn = &Connection{ID: rand.Intn(1000)}
	})
	return pg.conn
}

// для сравнения тот же вариант без once.DO

func (pg *PostgresSQLConnPool) GetConnection() *Connection {
	time.Sleep(50 * time.Millisecond)
	pg.conn = &Connection{ID: rand.Intn(1000)}
	fmt.Println("[UNSAFE]Подключение инициализированно")
	return pg.conn
}
