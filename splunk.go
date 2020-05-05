package splunk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
)

type SplunkConnection struct {
	Username, Password, BaseURL string
	sessionKey                  SessionKey
}

// SessionKey represents the JSON object returned from the Splunk authentication REST call
type SessionKey struct {
	Value string `json:"sessionKey,omitempty"`
}

// Login connects to the Splunk server and retrieves a session key
func (conn *SplunkConnection) Login() (SessionKey, error) {

	data := make(url.Values)
	data.Add("username", conn.Username)
	data.Add("password", conn.Password)
	data.Add("output_mode", "json")
	response, err := conn.httpPost(fmt.Sprintf("%s/services/auth/login", conn.BaseURL), &data)

	if err != nil {
		return SessionKey{}, err
	}

	bytes := []byte(response)
	var key SessionKey
	unmarshall_error := json.Unmarshal(bytes, &key)

	if key.Value == "" {
		return SessionKey{}, errors.New(response)
	}

	conn.sessionKey.Value = key.Value
	return conn.sessionKey, unmarshall_error
}

func CreateConnectionFromEnvironment() (*SplunkConnection, error) {

	var splunkUsername string
	var splunkPassword string
	var splunkUrl string

	if splunkUsername = os.Getenv("SPLUNK_USERNAME"); splunkUsername == "" {
		return nil, fmt.Errorf("Invalid value for environment variable SPLUNK_USERNAME: %v", splunkUsername)
	}

	if splunkPassword = os.Getenv("SPLUNK_PASSWORD"); splunkPassword == "" {
		return nil, fmt.Errorf("Invalid value for environment variable SPLUNK_PASSWORD: %v", splunkPassword)
	}

	if splunkUrl = os.Getenv("SPLUNK_URL"); splunkUrl == "" {
		return nil, fmt.Errorf("Invalid value for environment variable SPLUNK_URL: %v", splunkUrl)
	}

	return &SplunkConnection{
		Username: splunkUsername,
		Password: splunkPassword,
		BaseURL:  splunkUrl,
	}, nil
}
