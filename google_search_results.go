package google_search_results

/*
 * This package enables to interact with SerpApi server
 */

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Hold SerpApi user key
var serpApiKey string

// SerpQuery holds query parameter
type SerpQuery struct {
	parameter map[string]string
}

// SerpResponse holds response
type SerpResponse map[string]interface{}

// Create a new query
func newGoogleSearch(parameter map[string]string) SerpQuery {
	if len(serpApiKey) > 0 {
		parameter["serp_api_key"] = serpApiKey
	}

	return SerpQuery{parameter: parameter}
}

// Set Serp API key
func setAPIKey(key string) {
	serpApiKey = key
}

// Execute the query
func (sq *SerpQuery) execute(output string) *http.Response {
	query := url.Values{}
	for k, v := range sq.parameter {
		query.Add(k, v)
	}
	query.Add("source", "go")
	query.Add("output", output)
	endpoint := "https://serpapi.com/search?" + query.Encode()
	var client = &http.Client{
		Timeout: time.Second * 60,
	}
	rsp, err := client.Get(endpoint)

	if err != nil {
		panic(err.Error())
	}
	return rsp
}

// return go struct by processing the json returned by the server
func (sq *SerpQuery) json() (SerpResponse, error) {
	rsp := sq.execute("json")
	return sq.decodeJSON(rsp.Body)
}

// decodeJson response
func (sq *SerpQuery) decodeJSON(body io.ReadCloser) (SerpResponse, error) {
	// Decode JSON from response body
	decoder := json.NewDecoder(body)
	//var serpResponse SerpResponse
	var serpResponse SerpResponse
	err := decoder.Decode(&serpResponse)
	if err != nil {
		return nil, errors.New("fail to decode")
	}

	// check error message
	errorMessage, derror := serpResponse["error"].(string)
	if derror {
		return nil, errors.New(errorMessage)
	}
	return serpResponse, nil
}

// return html as a string
func (sq *SerpQuery) html() (*string, error) {
	rsp := sq.execute("html")
	return sq.decodeHTML(rsp.Body)
}

// decodeHTML
func (sq *SerpQuery) decodeHTML(body io.ReadCloser) (*string, error) {
	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	text := string(buffer)
	return &text, nil
}
