package models

import "time"

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
func (u *User) CreateSession(session Session, err error) {
	statement := "insert into users(uuid, email, user_id, created_at) values (?, ?, ?, ?);"
	stat, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stat.Close()

	uuid := createUUID()
	stat.Exec(uuid, u.Name, u.Email, u.ID, u.CreatedAt)

	stmtOut, err := Db.Prepare("select id, uuid, email, user_id, created_id from sessions where uuid = ?;")
	if err != nil {
		return
	}
	defer stmtOut.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtOut.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Session get the session for an existing user
func (u *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("select id, uuid, email, user_id, created_id from sessions where uuid = ?;", u.UUID).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Create a new user, save user info into the database
func (u *User) Create() (err error) {
	statement := "insert into users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?);"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, user.Name, user.Email, Sha1(user.Password), time.Now())

	stmtout, err := Db.Prepare("select id, uuid, created_at from users where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmtout.QueryRow(uuid).Scan(&u.Id, &u.Uuid, &u.CreatedAt)
	return
}

// Delete user from database
func (u *User) Delete() (err error) {
	statement := "delete from users where uuid = ?;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}

	_, err = stmt.Exec(u.Id)
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

	_, err = stmt.Exec(u.Name, u.Email, u.Id)
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
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
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
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// UserByUUID get a single user with the given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
