package UserDB

import "time"

type User struct {
	ID   string
	Name string
}

type UserDB struct {
	users []*User
}

func ConnectUserDB() *UserDB {
	users := []*User{
		{ID: "user_1", Name: "Лютик"},
		{ID: "user_2", Name: "Йенифер"},
		{ID: "user_3", Name: "Ламберт"},
		{ID: "user_4", Name: "Трис Меригольд"},
		{ID: "user_6", Name: "Геральт из Ривии"},
		{ID: "user_7", Name: "Вернон Роше"},
		{ID: "user_8", Name: "Цирилла"},
		{ID: "user_9", Name: "Эмгыр Вар Эмрейс"},
	}
	return &UserDB{users: users}
}

func (db *UserDB) GetUsers() []*User {
	time.Sleep(600 * time.Millisecond)
	return db.users
}
