package rwdatabasepool

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

// Driver is the NoDatabase driver.
type Driver struct{}
type conn struct{}

var errNoDatabase = errors.New("No database found")

// Open opens a new connection to the database.
func (d *Driver) Open(name string) (driver.Conn, error) {
	return conn{}, nil
}

func init() {
	sql.Register("nodatabase", &Driver{})
}

func (conn) Prepare(query string) (driver.Stmt, error) {
	return nil, errNoDatabase
}

func (conn) Begin() (driver.Tx, error) {
	return nil, errNoDatabase
}

func (conn) Close() error {
	return nil
}
