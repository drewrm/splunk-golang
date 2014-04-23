package splunk

import (
        "fmt"
        "bytes"
        "net/http"
        "net/url"
        "io"
        "io/ioutil"
        "crypto/tls"
)

/*
 * HTTP helper methods
 */

func httpClient() *http.Client {
        tr := &http.Transport{
                TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
        client := &http.Client{Transport: tr}
        return client
}

func (conn SplunkConnection) httpGet(url string, data *url.Values) (string, error) {
        return conn.httpCall(url, "GET", data)
}

func (conn SplunkConnection) httpPost(url string, data *url.Values) (string, error) {
        return conn.httpCall(url, "POST", data)
}

func (conn SplunkConnection) httpCall(url string, method string, data *url.Values) (string, error) {
        client := httpClient()

        var payload io.Reader
        if data != nil {
          payload = bytes.NewBufferString(data.Encode())
        }

        request, err := http.NewRequest("POST", url, payload)
        conn.addAuthHeader(request)
        response, err := client.Do(request)

        if err != nil {
                return "", err
        }

        body, _ := ioutil.ReadAll(response.Body)
        response.Body.Close()
        return string(body), nil
}


func (conn SplunkConnection) addAuthHeader(request *http.Request) {
        if conn.sessionKey.Value != "" {
                request.Header.Add("Authorization", fmt.Sprintf("Splunk %s", conn.sessionKey))
        } else {
                request.SetBasicAuth(conn.Username, conn.Password)
        }
}
