package Sessions

import (
	"time"
)

type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type SessionDB struct {
	sessions []*Session
}

func ConnectSessionDB() *SessionDB {
	sessions := []*Session{
		{
			ID:        "sess_1",
			UserID:    "user_1",
			CreatedAt: time.Now().Add(-1 * time.Hour),
			ExpiresAt: time.Now().Add(2 * time.Hour),
		},
		{
			ID:        "sess_2",
			UserID:    "user_2",
			CreatedAt: time.Now().Add(-30 * time.Minute),
			ExpiresAt: time.Now().Add(3 * time.Hour),
		},
		{
			ID:        "sess_3",
			UserID:    "user_3",
			CreatedAt: time.Now(),
			ExpiresAt: time.Now().Add(1 * time.Hour),
		},
	}
	return &SessionDB{sessions: sessions}
}

func (db *SessionDB) GetSessions() []*Session {
	time.Sleep(1000 * time.Millisecond)
	return db.sessions
}
