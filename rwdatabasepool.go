package rwdatabasepool

import (
	"database/sql"
	"log"
)

// RWDatabasePool represents the read-write database pool.
type RWDatabasePool struct {
	writePool    []*sql.DB
	readPool     []*sql.DB
	writeCounter int
	readCounter  int
}

// Init creates a RWDatabasePool using a write and a read pool.
func Init(write, read []*sql.DB) *RWDatabasePool {
	noDB, err := sql.Open("nodatabase", "")
	if err != nil {
		log.Fatal(err)
	}
	if len(write) == 0 {
		write = []*sql.DB{noDB}
	}
	if len(read) == 0 {
		read = write
	}
	return &RWDatabasePool{
		writePool:    write,
		readPool:     read,
		writeCounter: 0,
		readCounter:  0,
	}
}

func (dbp *RWDatabasePool) Write() *sql.DB {
	nbWrite := len(dbp.writePool)
	dbp.writeCounter = (dbp.writeCounter + 1) % nbWrite
	return dbp.writePool[dbp.writeCounter]
}

func (dbp *RWDatabasePool) Read() *sql.DB {
	nbRead := len(dbp.readPool)
	dbp.readCounter = (dbp.readCounter + 1) % nbRead
	return dbp.readPool[dbp.readCounter]
}
