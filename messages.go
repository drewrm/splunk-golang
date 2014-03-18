package splunk

import (
        "fmt"
        "net/url"
)

type MessageSeverity string

const (
        Info MessageSeverity = "info"
        Warn MessageSeverity = "warn"
        Error MessageSeverity = "error"
)

// SendMessage sends an informational message to Splunk
func (conn SplunkConnection) SendMessage(name string, message string, severity MessageSeverity) (string, error){
        data := make(url.Values)
        data.Add("name", name)
        data.Add("value", message)
        data.Add("severity", string(severity))
        client := httpClient()
        response, err := httpPost(client, fmt.Sprintf("%s/services/messages", conn.BaseURL), data, conn)
        return response, err
}
