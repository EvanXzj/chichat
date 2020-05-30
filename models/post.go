package models

import "time"

// Post struct
type Post struct {
	ID        int
	UUID      string
	ThreadID  int
	Body      string
	UserID    int
	CreatedAt time.Time
}

// FormatCreatedAt format createdAt time
func (p *Post) FormatCreatedAt() string {
	return p.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// User get the user who wrote the post
func (p *Post) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", p.UserID).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return
}
