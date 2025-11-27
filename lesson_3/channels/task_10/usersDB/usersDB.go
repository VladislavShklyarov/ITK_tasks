package UserDB

import (
	"sync"
	"time"
)

type User struct {
	Id   string
	Name string
}

type UserDB struct {
	users []*User
	once  sync.Once
}

func ConnectUserDB() *UserDB {
	users := []*User{
		{Id: "user_1", Name: "Лютик"},
		{Id: "user_2", Name: "Йенифер"},
		{Id: "user_3", Name: "Ламберт"},
		{Id: "user_4", Name: "Трис Меригольд"},
		{Id: "user_6", Name: "Геральт из Ривии"},
		{Id: "user_7", Name: "Вернон Роше"},
		{Id: "user_8", Name: "Цирилла"},
		{Id: "user_9", Name: "Эмгыр Вар Эмрейс"},
	}
	return &UserDB{users: users}
}

func (db *UserDB) GetUsers() []*User {
	db.once.Do(func() {
		time.Sleep(600 * time.Millisecond) // имитируем тяжелую загрузку
	}) // грузим один раз, а дальнейшие вызовы просто возвращают то же самое
	return db.users
}
