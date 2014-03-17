package splunk

import (
        "fmt"
        "bytes"
        "encoding/json"
        "net/http"
        "io/ioutil"
        "net/url"
        "crypto/tls"
)


type MessageSeverity string

const (
        Info MessageSeverity = "info"
        Warn MessageSeverity = "warn"
        Error MessageSeverity = "error"
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

func httpClient() *http.Client {
        tr := &http.Transport{
                TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
        client := &http.Client{Transport: tr}
        return client
}

func httpPost(client *http.Client, url string, postData url.Values, conn SplunkConnection) (string, error) {
        request, err := http.NewRequest("POST", url, bytes.NewBufferString(postData.Encode()))

        if conn.sessionKey.Value != "" {
                request.Header.Add("Authorization", fmt.Sprintf("Splunk %s", conn.sessionKey))
        } else {
                request.SetBasicAuth(conn.Username, conn.Password)
        }

        response, err := client.Do(request)

        if err != nil {
                return "", err
        }

        body, _ := ioutil.ReadAll(response.Body)
        response.Body.Close()
        return string(body), nil
}
