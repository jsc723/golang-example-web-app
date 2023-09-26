package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"learn_go/todo_app/config"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id Integer Primary Key Autoincrement,
		uuid String Not Null Unique,
		name String,
		email String,
		password String,
		create_at Datatime
	)`, tableNameUser)
	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		create_at Datatime
	)`, tableNameTodo)
	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE	,
		email STRING,
		user_id INTEGER,
		create_at Datetime
	)`, tableNameSession)
	Db.Exec(cmdS)
}

func createUUID() uuid.UUID {
	uuidobj, _ := uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) string {
	cryptext := fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

func readTime(t *time.Time, dt string) error {
	*t, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", dt)
	return err
}
