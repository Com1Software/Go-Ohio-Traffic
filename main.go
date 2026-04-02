package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type Incident struct {
    ID          string `json:"id"`
    Type        string `json:"type"`
    Description string `json:"description"`
    StartTime   string `json:"startTime"`
    Road        string `json:"road"`
    Direction   string `json:"direction"`
}

func main() {
    apiKey := os.Getenv("OHGO_API_KEY")
    url := "https://api.ohgo.com/v1/incidents?key=" + apiKey

    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    var incidents []Incident
    if err := json.NewDecoder(resp.Body).Decode(&incidents); err != nil {
        panic(err)
    }

    for _, inc := range incidents {
        fmt.Printf("%s: %s on %s %s\n",
            inc.ID, inc.Description, inc.Road, inc.Direction)
    }
}
