package main

import (
	"fmt"
	"sync"
	cDB "task_10/commentsDB"
	ss "task_10/sessions"
)

func main() {

	wg := sync.WaitGroup{}

	commentsConn := cDB.ConnectCommentsDB()
	sessionConn := ss.ConnectSessionDB()

	commentsCh, sessionsCh := LoadCommentsSessions(commentsConn, sessionConn)

	wg.Go(func() {

		for comment := range commentsCh {
			fmt.Println(comment)
		}
	})

	wg.Go(func() {
		for session := range sessionsCh {
			fmt.Println(session)
		}
	})

	wg.Wait()
}

func LoadCommentsSessions(cc *cDB.CommentsDB, sc *ss.SessionDB) (chan *cDB.Comment, chan *ss.Session) {
	wg := sync.WaitGroup{}

	commentsCh := make(chan *cDB.Comment)
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
