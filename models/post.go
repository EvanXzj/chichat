package models

import "time"

// Post struct
type Post struct {
	ID        int
	UUID      int
	ThreadID  int
	Body      string
	UserID    int
	CreatedAt time.Time
}
