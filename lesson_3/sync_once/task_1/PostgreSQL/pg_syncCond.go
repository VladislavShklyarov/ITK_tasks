package PostgreSQL

import (
	"math/rand"
	"sync"
	"time"
)

type PostgresSQLConnPoolSC struct {
	name string
	conn *ConnectionSC
	cond *sync.Cond
}

type ConnectionSC struct {
	ID int
}

func PostgresNewPoolSC() *PostgresSQLConnPoolSC {
	return &PostgresSQLConnPoolSC{
		name: "PostgreSQL",
		conn: nil,
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (pg *PostgresSQLConnPoolSC) GetConnectionSC() *ConnectionSC {
	pg.cond.L.Lock()
	defer pg.cond.L.Unlock()

	if pg.conn == nil {
		time.Sleep(50 * time.Millisecond)
		pg.conn = &ConnectionSC{ID: rand.Intn(1000)}
	}
	return pg.conn
}
