package models

import (
	"fmt"
	"time"
)

// User users table struct
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// CreateSession create a new session for an existing user
func (u *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions(uuid, email, user_id) values (?, ?, ?)"
	stat, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stat.Close()

	uuid := createUUID()
	stat.Exec(uuid, u.Email, u.ID)

	stmtOut, err := Db.Prepare("select id, uuid, email, user_id, created_at from sessions where uuid = ?;")
	if err != nil {
		return
	}
	defer stmtOut.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtOut.QueryRow(uuid).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return
}

// Session get the session for an existing user
func (u *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("select id, uuid, email, user_id, created_id from sessions where uuid = ?;", u.UUID).
		Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return
}

// Create a new user, save user info into the database
func (u *User) Create() (err error) {
	statement := "insert into users (uuid, name, email, password) values (?, ?, ?, ?);"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, u.Name, u.Email, Sha1(u.Password))

	stmtout, err := Db.Prepare("select id, uuid, created_at from users where uuid = ?")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmtout.Close()
	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmtout.QueryRow(uuid).Scan(&u.ID, &u.UUID, &u.CreatedAt)
	return
}

// Delete user from database
func (u *User) Delete() (err error) {
	statement := "delete from users where uuid = ?;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}

	_, err = stmt.Exec(u.ID)
	return
}

// Update user information in the database
func (u *User) Update() (err error) {
	statement := "update users set name = ?, email = ? where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Email, u.ID)
	return
}

// DeleteAllUsers Delete all users from database
func DeleteAllUsers() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

// Users Get all users in the database and returns it
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// UserByEmail get a single user with the given the email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// UserByUUID get a single user with the given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = ?", uuid).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// CreateThread create a new thread
func (u *User) CreateThread(topic string) (thread Thread, err error) {
	statement := "insert into threads (uuid, topic, user_id) value (?, ?, ?);"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, topic, u.ID)

	stmtout, err := Db.Prepare("select id, uuid, topic, user_id, created_at from threads where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	// thread = Thread{}
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtout.QueryRow(uuid).Scan(&thread.ID, &thread.UUID, &thread.Topic, &thread.UserID, &thread.CreatedAt)
	return
}

// CreatePost create a new post to a thread
func (u *User) CreatePost(thread Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id) values (?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, body, u.ID, thread.ID)

	stmtout, err := Db.Prepare("select id, uuid, body, user_id, thread_id, created_at from posts where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtout.QueryRow(uuid).Scan(&post.ID, &post.UUID, &post.Body, &post.UserID, &post.ThreadID, &post.CreatedAt)
	return
}
