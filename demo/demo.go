package main

import (
	"fmt"
	g "github.com/serpapi/google-search-results-golang"
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
		"api_key":  os.Getenv("API_KEY"), // your api key
	}

	client := g.NewGoogleSearch(parameter)
	serpResponse, err := client.GetJSON()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	results := serpResponse["organic_results"].([]interface{})
	firstResult := results[0].(map[string]interface{})
	fmt.Println(firstResult["title"].(string))
}
