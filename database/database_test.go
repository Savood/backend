package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConnectDatabase(t *testing.T) {
	assert.NoError(t, ConnectDatabase(nil, nil))
}
