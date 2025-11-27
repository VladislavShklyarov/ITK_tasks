package commentsDB

import (
	"time"
)

type CommentsDB struct {
	comments []*Comment
}

type Comment struct {
	Id   string
	Text []string
}

func ConnectCommentsDB() *CommentsDB {
	comments := []*Comment{
		{
			Id: "user_1",
			Text: []string{
				"Не верь всему, что говорят в таверне… особенно, если это про тебя.",
				"Музыка — лучшее оружие против скуки и уныния.",
			},
		},
		{
			Id: "user_4",
			Text: []string{
				"Не бывает слишком много магии… только недостаточно осторожности.",
			},
		},
		{
			Id: "user_6",
			Text: []string{
				"Зараза!",
				"Если нужно выбирать между меньшим и большим злом...",
				"Ламберт-ламберт...",
				"Сыграем в Гвинт?",
			},
		},
	}
	return &CommentsDB{comments: comments}
}
func (db *CommentsDB) GetComments() []*Comment {
	time.Sleep(500 * time.Millisecond)
	return db.comments
}
