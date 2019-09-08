package splunk

import (
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
