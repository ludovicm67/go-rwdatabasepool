package rwdatabasepool

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	// create two database connections
	noDB1, err := sql.Open("nodatabase", "")
	assert.NoError(t, err)
	noDB2, err := sql.Open("nodatabase", "")
	assert.NoError(t, err)

	// create two arrays
	dbWrite := []*sql.DB{noDB1}
	dbRead := []*sql.DB{noDB2}

	// check database pool
	pool := Init(dbWrite, dbRead)
	assert.Equal(t, pool.writePool, dbWrite)
	assert.Equal(t, pool.readPool, dbRead)
	assert.Zero(t, pool.writeCounter)
	assert.Zero(t, pool.readCounter)

	// try some connections
	for i := 0; i < 5; i++ {
		w := pool.Write()
		assert.Equal(t, w, noDB1)
		assert.Zero(t, pool.writeCounter)

		r := pool.Read()
		assert.Equal(t, r, noDB2)
		assert.Zero(t, pool.readCounter)
	}
}

func TestMultpileWrite(t *testing.T) {
	// create two database connections
	noDB1, err := sql.Open("nodatabase", "")
	assert.NoError(t, err)
	noDB2, err := sql.Open("nodatabase", "")
	assert.NoError(t, err)

	// create two arrays
	dbWrite := []*sql.DB{noDB1, noDB2}
	dbRead := []*sql.DB{}

	// check database pool
	pool := Init(dbWrite, dbRead)
	assert.Equal(t, pool.writePool, dbWrite)
	assert.Equal(t, pool.readPool, dbWrite) // read = write
	assert.Zero(t, pool.writeCounter)
	assert.Zero(t, pool.readCounter)

	// try some connections
	w1 := pool.Write()
	assert.Equal(t, w1, noDB2)
	assert.Equal(t, 1, pool.writeCounter)
	assert.Zero(t, pool.readCounter)

	w2 := pool.Write()
	assert.Equal(t, w2, noDB1)
	assert.Equal(t, 0, pool.writeCounter)
	assert.Zero(t, pool.readCounter)

	w3 := pool.Write()
	assert.Equal(t, w3, noDB2)
	assert.Equal(t, 1, pool.writeCounter)
	assert.Zero(t, pool.readCounter)
}

func TestMultpileRead(t *testing.T) {
	// create two database connections
	noDB1, err := sql.Open("nodatabase", "")
	assert.NoError(t, err)
	noDB2, err := sql.Open("nodatabase", "")
	assert.NoError(t, err)

	// create two arrays
	dbWrite := []*sql.DB{}
	dbRead := []*sql.DB{noDB1, noDB2}

	// check database pool
	pool := Init(dbWrite, dbRead)
	assert.Equal(t, 1, len(pool.writePool)) // one write connection will be created
	assert.Equal(t, pool.readPool, dbRead)
	assert.Zero(t, pool.writeCounter)
	assert.Zero(t, pool.readCounter)

	// try some connections
	r1 := pool.Read()
	assert.Equal(t, r1, noDB2)
	assert.Equal(t, 1, pool.readCounter)
	assert.Zero(t, pool.writeCounter)

	r2 := pool.Read()
	assert.Equal(t, r2, noDB1)
	assert.Equal(t, 0, pool.readCounter)
	assert.Zero(t, pool.writeCounter)

	r3 := pool.Read()
	assert.Equal(t, r3, noDB2)
	assert.Equal(t, 1, pool.readCounter)
	assert.Zero(t, pool.writeCounter)
}
