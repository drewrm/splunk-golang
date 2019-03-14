# Splunk REST API Client

Example Usage:

    package main

    import (
        "fmt"
        "github.com/drewrm/splunk"
    )

    func main() {
        conn := splunk.SplunkConnection {
                Username: "admin",
                Password: "changeme",
                BaseURL: "https://localhost:8089",
        }

        key, err:= conn.Login()

        if err != nil {
                fmt.Println("Couldn't login to splunk: ", err)
        }

        fmt.Println("Session key: ", key.Value)
    }
