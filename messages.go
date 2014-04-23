package splunk

import (
        "fmt"
        "time"
        "net/url"
        "encoding/json"
)

type MessageSeverity string

type Message struct {
        Name string `json:"name"`
        Content MessageContent `json:"content"`
}


type MessageContent struct {
        Message string `json:"message"`
        Severity MessageSeverity `json:"severity"`
        created int64 `json:"timeCreated_epochSecs"`
}

func (mc *MessageContent) Content() (time.Time) {
        return time.Unix(mc.created, 0)
}

type Messages struct {
        Origin string `json:"origin"`
        Messages []Message `json:"entry"`
}

const (
        Info MessageSeverity = "info"
        Warn MessageSeverity = "warn"
        Error MessageSeverity = "error"
)

// SendMessage sends an informational message to Splunk
func (conn SplunkConnection) SendMessage(message *Message) (string, error) {
        data := make(url.Values)
        data.Add("name", message.Name)
        data.Add("value", message.Content.Message)
        data.Add("severity", string(message.Content.Severity))
        response, err := conn.httpPost(fmt.Sprintf("%s/services/messages", conn.BaseURL), &data)
        return response, err
}


func (conn SplunkConnection) GetMessage(name string) ([]Message, error) {
        data := make(url.Values)
        data.Add("name", name)
        data.Add("output_mode", "json")
        response, err := conn.httpGet(fmt.Sprintf("%s/services/messages/%s", conn.BaseURL, name), &data)

        if err != nil {
                return []Message{}, err
        }

        bytes := []byte(response)
        var messages Messages
        unmarshall_error := json.Unmarshal(bytes, &messages)
        return messages.Messages, unmarshall_error
}
