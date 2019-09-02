package splunk

import (
	"github.com/prometheus/common/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestICouldSearchSync(t *testing.T) {

	connection, err := CreateConnectionFromEnvironment()

	assert.NotNil(t, connection, "error on searching")

	sessionKey, err := connection.Login()

	assert.NotNil(t, sessionKey, "error on searching")

	searchString := "search index=hotels_com kubernetes.pod_id=60500b25-cbf3-11e9-8768-02b2860aca1c"
	searchResponse, err := connection.SearchSync(searchString)

	assert.Nil(t, err, "error on searching")
	log.Infof("search is %s ", searchResponse)
	assert.NotNil(t, searchResponse, "error on searching")

}
