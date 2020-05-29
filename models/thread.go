package models

import "time"

// Thread struct
type Thread struct {
	ID        int
	UUID      string
	Topic     string
	UserID    int
	CreatedAt time.Time
}

// FormatCreatedAt format the CreatedAt date to display nicely on the screen
func (t *Thread) FormatCreatedAt() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// NumReplies get the number of posts in a thread
func (t *Thread) NumReplies() (count int) {
	statement := "select count(*) from posts where thread_id = ?;"
	rows, err := Db.Query(statement, t.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// Posts get posts to a thread
func (t *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = ?", t.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.ID, &post.UUID, &post.Body, &post.UserID, &post.ThreadID, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// User get the user who started this thread
func (t *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", t.UserID).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// Treads get all threads in the database and returns it
func Treads() (threads []Thread, err error) {
	rows, err := Db.Query("select id, uuid, topic, user_id, created_at from threads order created_at desc;")
	if err != nil {
		return
	}

	for rows.Next() {
		thread := Thread{}
		if err = rows.Scan(&thread.ID, &thread.UUID, &thread.Topic, &thread.UserID, &thread.CreatedAt); err != nil {
			return
		}
		threads = append(threads, thread)
	}
	return
}

// ThreadByUUID get a thread by the UUID
func ThreadByUUID(uuid string) (thread Thread, err error) {
	thread = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = ?", uuid).
		Scan(&thread.ID, &thread.UUID, &thread.Topic, &thread.UserID, &thread.CreatedAt)
	return
}
