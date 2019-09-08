package splunk

import (
	"fmt"
	"net/url"
)

func (conn SplunkConnection) SearchSync(searchString string) (string, error) {
	data := make(url.Values)
	data.Add("search", searchString)
	data.Add("max_count", "1")
	data.Add("max_time", "60")
	data.Add("earliest_time", "-10m")
	data.Add("latest_time", "now")
	data.Add("output_mode", "json")

	response, err := conn.httpPost(fmt.Sprintf("%s/services/search/jobs/export", conn.BaseURL), &data)
	return response, err
}
