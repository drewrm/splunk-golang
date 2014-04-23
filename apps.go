package splunk

import (
        "fmt"
        "net/url"
)

func (conn SplunkConnection) InstallApp(path string, update bool) (string, error) {
        data := make(url.Values)
        data.Add("name", path)

        update_app := "false"
        if (update == true) {
                update_app = "true"
        }

        data.Add("update", update_app)
        response, err := conn.httpPost(fmt.Sprintf("%s/services/apps/appinstall/", conn.BaseURL), &data)
        return response, err
}
