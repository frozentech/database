package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// client ...
var (
	client *Database
)

// Database ...
type Database struct {
	connection *sqlx.DB
	isClose    bool
}

// New ...
func New() *Database {
	return &Database{}
}

// Connect ...
func (me *Database) Connect(username string, password, host string, port string, dbname string, engine string) (err error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		port,
		dbname)

	me.connection, err = sqlx.Connect(engine, connection)
	return
}

// GetConnection ...
func (me *Database) GetConnection() *sqlx.DB {
	return me.connection
}

// IsClose ...
func (me *Database) IsClose() bool {
	if me.connection == nil || me.isClose {
		return true
	}
	return false
}

// Destroy ...
func (me *Database) Destroy() {
	me.connection.Close()
	me.isClose = true
	return
}

// Destroy connection
func Destroy() {
	if client != nil {
		client.Destroy()
	}
}

// GetConnection ...
func GetConnection() *sqlx.DB {
	return client.GetConnection()
}

// Connect to database
func Connect(username, password, host, port, dbname, engine string) (err error) {
	if client == nil || client.IsClose() {
		client = New()
		err = client.Connect(username, password, host, port, dbname, engine)
		return
	}

	return nil
}
