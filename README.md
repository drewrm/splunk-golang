# Splunk REST API Client

Example Usage:

    package main

    import (
        "fmt"
        "github.com/drewrm/splunk"
    )

    func main() {
        splunk := splunk.SplunkConnection {"admin", "changeme", "https://localhost:8089"}
        key, err:= splunk.Login()

        if err != nil {
            fmt.Println("Couldn't login to splunk: %s", err)
        }

        fmt.Println("Session key: ", key.Value)
    }
