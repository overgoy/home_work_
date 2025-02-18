package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	dbConn, err := Connect()
	assert.Nil(t, err)
	assert.NotNil(t, dbConn)
	defer dbConn.Close()
}
