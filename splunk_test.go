package splunk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCreateConnectionFromEnvironment(t *testing.T) {

	connection, err := CreateConnectionFromEnvironment()

	assert.Nil(t, err, "error on create connection from environment")
	assert.NotNil(t, connection, "error on create connection from environment")

}

func TestCanLogin(t *testing.T) {

	connection, err := CreateConnectionFromEnvironment()

	assert.Nil(t, err, "error on create connection from environment")
	assert.NotNil(t, connection, "error on create connection from environment")

	sessionKey, err := connection.Login()

	assert.NotNil(t, sessionKey, "error while login")
	assert.NotEmpty(t, sessionKey.Value, "error while login")

}
