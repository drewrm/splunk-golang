package splunk

import (
        "fmt"
        "net/url"
        "encoding/json"
)

type SplunkConnection struct {
        Username, Password, BaseURL string
        sessionKey SessionKey
}

// SessionKey represents the JSON object returned from the Splunk authentication REST call
type SessionKey struct {
        Value string `json:"sessionKey"`
}

// Login connects to the Splunk server and retrieves a session key
func (conn SplunkConnection) Login() (SessionKey, error){

        data := make(url.Values)
        data.Add("username", conn.Username)
        data.Add("password", conn.Password)
        data.Add("output_mode", "json")
        client := httpClient()
        response, err := httpPost(client, fmt.Sprintf("%s/services/auth/login", conn.BaseURL), data, conn)

        if err != nil {
          return SessionKey{}, err
        }

        bytes := []byte(response)
        var key SessionKey
        unmarshall_error := json.Unmarshal(bytes, &key)
        conn.sessionKey = key
        return key, unmarshall_error
}

