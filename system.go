package splunk

import (
        "fmt"
)

func (conn SplunkConnection) RestartServer() (string, error) {
        response, err := conn.httpPost(fmt.Sprintf("%s/services/server/control/restart", conn.BaseURL), nil)
        return response, err
}
