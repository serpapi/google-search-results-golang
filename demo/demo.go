package main

import (
	"fmt"
	g "github.com/serpapi/google-search-results-golang"
)

/***
 * demo how to create a client for SerpApi
 *
 * go get -u github.com/serpapi/google_search_results_golang
 */
func main() {
	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Portland",
		"api_key":  "2a09c3d5c663314640a0c4382bc2d6be73c233417011040f689318801bf9d328",
	}

	client := g.NewGoogleSearch(parameter)
	serpResponse, err := client.GetJSON()
	results := serpResponse["organic_results"].([]interface{})
	firstResult := results[0].(map[string]interface{})
	fmt.Println(firstResult["title"].(string))
}
