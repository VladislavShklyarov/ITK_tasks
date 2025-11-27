package main

import (
	"fmt"
	"strings"
	"sync"
	att "task_10/attachments"
	ct "task_10/commentsDB"
	ss "task_10/sessions"
	us "task_10/usersDB"
)

func main() {

	wg := sync.WaitGroup{}

	commentsConn := ct.ConnectCommentsDB()
	sessionConn := ss.ConnectSessionDB()
	usersConn := us.ConnectUserDB()
	attConn := att.ConnectAttachmentDB()

	commentsCh, sessionsCh := LoadCommentsSessions(commentsConn, sessionConn)

	usersCh := LoadUsers(commentsCh, usersConn)
	attCh := LoadAttachments(sessionsCh, attConn)

	PrintUsers(usersCh)
	PrintAttachments(attCh)
	wg.Wait()
}

func LoadCommentsSessions(cc *ct.CommentsDB, sc *ss.SessionDB) (chan *ct.Comment, chan *ss.Session) {
	wg := sync.WaitGroup{}
	// Если у нас условие, что загружаться должно одновременно,
	// то насколько правильно стартавать загрузку внутри одной функции?
	// Потому что можно то же самое сделать в main отдельно для комментариев и сессий.
	commentsCh := make(chan *ct.Comment)
	sessionsCh := make(chan *ss.Session)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, comment := range cc.GetComments() {
			commentsCh <- comment
		}
		close(commentsCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, session := range sc.GetSessions() {
			sessionsCh <- session
		}
		close(sessionsCh)
	}()

	return commentsCh, sessionsCh

}

func LoadUsers(commentsCh chan *ct.Comment, usersConn *us.UserDB) chan *us.User {
	wg := sync.WaitGroup{}
	// Сказано, что должно загружаться только после того, как получили комментарии
	// поэтому решил вычитывать из канала внутри этой функции.
	out := make(chan *us.User)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// не самое оптимальное решение с вложенным циклом)
		for comment := range commentsCh {
			users := usersConn.GetUsers() // Здесь специально засунул в цикл,
			// чтобы было несколько запросов и можно было использовать sync.Once
			for _, user := range users {
				if user.Id == comment.Id {
					out <- user
				}
			}
		}
		close(out)
	}()
	return out
}

func LoadAttachments(ssCh chan *ss.Session, attConn *att.AttDB) chan *att.Attachments {
	wg := sync.WaitGroup{}
	out := make(chan *att.Attachments)
	// Здесь вроде получилось получше, т.к. мы храним вложения в мапе
	wg.Add(1)
	go func() {
		defer wg.Done()
		for session := range ssCh {
			attachment := attConn.GetAttachments(session.ID)
			out <- attachment
		}
		close(out)
	}()

	return out
}

func PrintUsers(userCh chan *us.User) {
	for user := range userCh {
		fmt.Printf("id: %s, name: %s\n", user.Id, user.Name)
	}
}
func PrintAttachments(attCh chan *att.Attachments) {
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {

		go func() {
			defer wg.Done()
			for attachment := range attCh {
				fmt.Printf("Session ID: %s by worker%d:, Files: %s\n",
					attachment.SessionId,
					i,
					strings.Join(attachment.Files, ", "))
			}
		}()
	}
	wg.Wait()
}
