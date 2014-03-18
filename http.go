package splunk

import (
        "fmt"
        "bytes"
        "net/http"
        "net/url"
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
