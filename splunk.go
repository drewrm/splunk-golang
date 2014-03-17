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

type SplunkConnection struct {
        Username, Password, BaseURL string
}

type SessionKey struct {
        Value string `json:"sessionKey"`
}

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
        return key, unmarshall_error
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
        request.SetBasicAuth(conn.Username, conn.Password)
        response, err := client.Do(request)

        if err != nil {
                return "", err
        }

        body, _ := ioutil.ReadAll(response.Body)
        response.Body.Close()
        return string(body), nil
}
