package main

import (
	"fmt"
	"os"

	serpapi "github.com/serpapi/google-search-results-golang"
)

/***
 * Demonstrate how to search on Google
 *
 * go get -u github.com/serpapi/google_search_results_golang
 */
func main() {
	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Austin,Texas",
	}

	search := serpapi.NewGoogleSearch(parameter, os.Getenv("API_KEY"))
	data, err := search.GetJSON()
	if err != nil {
		panic(err)
	}
	// decode data and display
	results := data["organic_results"].([]interface{})
	firstResult := results[0].(map[string]interface{})
	fmt.Println(firstResult["title"].(string))
}
