package splunk

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestICouldSearchSync(t *testing.T) {

	connection, _ := CreateConnectionFromEnvironment()
	require.NotNil(t, connection, "error on searching")

	sessionKey, _ := connection.Login()
	require.NotNil(t, sessionKey, "error on searching")

	searchString := "search index=_internal linecount=1"
	searchResponse, _ := connection.SearchSync(searchString)
	require.NotNil(t, searchResponse, "error on searching")
}

func TestICouldSearchSyncLongQuery(t *testing.T) {

	connection, _ := CreateConnectionFromEnvironment()
	require.NotNil(t, connection, "error on searching")

	sessionKey, _ := connection.Login()
	require.NotNil(t, sessionKey, "error on searching")

	searchString := "search index=_internal|head 100"
	searchResponse, _ := connection.SearchSync(searchString)
	assert.Contains(t, searchResponse, "linecount", "error on searching")
	require.NotNil(t, searchResponse, "error on searching")
}
