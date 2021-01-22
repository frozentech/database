package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Database Constant
const (
	MYSQLTimestampFormat = `2006-01-02 15:04:05`
)

// Client ...
var (
	Client *Database
)

// Database ...
type Database struct {
	Connection *sqlx.DB
	IsClose    bool
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

	me.Connection, err = sqlx.Connect(engine, connection)
	return
}

// GetConnection ...
func (me *Database) GetConnection() *sqlx.DB {
	return me.Connection
}

// Destroy connection
func Destroy() {
	if Client != nil {
		Client.Connection.Close()
		Client.IsClose = true
	}
}

// Connect to database
func Connect(username, password, host, port, dbname, engine string) (err error) {
	if Client == nil || Client.IsClose || Client.Connection == nil {
		Client = New()
		err = Client.Connect(username, password, host, port, dbname, engine)
		return
	}

	return nil
}
