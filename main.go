package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ArcGISResponse struct {
	Features []struct {
		Attributes IncidentAttributes `json:"attributes"`
	} `json:"features"`
}

type IncidentAttributes struct {
	ID          int     `json:"OBJECTID"`
	Road        string  `json:"ROADWAY"`
	Direction   string  `json:"DIRECTION"`
	Description string  `json:"DESCRIPTION"`
	County      string  `json:"COUNTY"`
	Latitude    float64 `json:"LATITUDE"`
	Longitude   float64 `json:"LONGITUDE"`
	Severity    string  `json:"SEVERITY"`
}

func main() {
	url := "https://gis.dot.state.oh.us/arcgis/rest/services/Transportation/Traffic_Incidents/FeatureServer/0/query?where=1%3D1&outFields=*&f=json"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		os.Exit(1)
	}

	fmt.Println("RAW RESPONSE:")
	fmt.Println(string(body))

	// Prevent JSON decode on HTML or empty body
	if len(body) == 0 || body[0] == '<' {
		fmt.Println("\nERROR: Server returned HTML or empty response, not JSON.")
		os.Exit(1)
	}

	var data ArcGISResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("JSON decode error:", err)
		os.Exit(1)
	}

	fmt.Println("\nParsed incidents:")
	for _, f := range data.Features {
		fmt.Printf("%+v\n", f.Attributes)
	}
}
