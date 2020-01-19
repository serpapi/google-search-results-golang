package main

import (
	"fmt"
	serpapi "github.com/serpapi/google-search-results-golang"
	"os"
)

/***
 * demo how to create a client for SerpApi
 *
 * go get -u github.com/serpapi/google_search_results_golang
 */
func main() {
	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Austin,Texas",
	}

	client := serpapi.NewGoogleClient(parameter, os.Getenv("API_KEY"))
	serpResponse, err := client.GetJSON()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	results := serpResponse["organic_results"].([]interface{})
	firstResult := results[0].(map[string]interface{})
	fmt.Println(firstResult["title"].(string))
}
