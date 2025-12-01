package PostgreSQL

import (
	"fmt"
	"time"
)

func (pg *Postgres) Create(key string, value string) {
	pg.data.mu.Lock()
	if _, ok := pg.data.vault[key]; !ok {
		pg.data.vault[key] = value
	}
	pg.data.mu.Unlock()

}

func (pg *Postgres) GetData(key string) (string, error) {
	pg.data.mu.Lock()
	defer pg.data.mu.Unlock()
	fmt.Println(pg.data.vault[key])
	value := ""
	ok := false
	if value, ok = pg.data.vault[key]; !ok {
		return "", fmt.Errorf("no data found")
	}
	time.Sleep(200 * time.Millisecond) // имитация задержки работы с БД
	return value, nil
}

func (pg *Postgres) GetAll() map[string]string {
	pg.data.mu.Lock()
	defer pg.data.mu.Unlock()

	return pg.data.vault
}
