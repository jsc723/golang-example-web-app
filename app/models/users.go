package models

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID       int
	UUID     string
	Name     string
	Email    string
	Password string
	CreateAt time.Time
	Todos    []Todo
}

type Session struct {
	ID       int
	UUID     string
	Email    string
	UserID   int
	CreateAt time.Time
}

func (u *User) CreateUser() error {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		create_at) values (?, ?, ?, ?, ?)
	`
	_, err := Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, create_at 
	from users where id = ?`
	var dt string
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&dt,
	)
	readTime(&user.CreateAt, dt)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, create_at from users where email = ?`
	var dt string
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&dt,
	)
	if err != nil {
		fmt.Println(err)
	}
	readTime(&user.CreateAt, dt)
	return
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (uuid, email, user_id, create_at) 
	values (?, ?, ?, ?)`
	_, err = Db.Exec(cmd1, u.UUID, u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}
	cmd2 := `select id, uuid, email, user_id, create_at
	from sessions where user_id = ? and email = ?`
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreateAt,
	)
	if err != nil {
		log.Println(err)
	}
	return
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, create_at
	from sessions where uuid = ?`
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreateAt,
	)
	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return
}

func (sess *Session) DeleteSessionByUUID() error {
	cmd := `delete from sessions where UUID = ?`
	_, err := Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	cmd := `select id, uuid, name, email, create_at from users where id = ?`
	var dt string
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&dt,
	)
	readTime(&user.CreateAt, dt)
	return
}
