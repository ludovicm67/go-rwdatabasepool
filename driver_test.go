package rwdatabasepool

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDriver(t *testing.T) {
	noDB, err := sql.Open("nodatabase", "")
	assert.NoError(t, err)
	assert.IsType(t, &Driver{}, noDB.Driver())

	_, err = noDB.Prepare("INSERT joy INTO people;")
	assert.Equal(t, errNoDatabase, err)

	_, err = noDB.Begin()
	assert.Equal(t, errNoDatabase, err)

	err = noDB.Close()
	assert.NoError(t, err)
}
