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
var apiKey string

// SerpQuery hold query parameter
type SerpQuery struct {
	parameter map[string]string
}

// SerpResponse hold response
type SerpResponse map[string]interface{}

// SerpResponseArray hold response array
type SerpResponseArray []interface{}

// Create a new query
func NewGoogleSearch(parameter map[string]string) SerpQuery {
	if len(apiKey) > 0 {
		parameter["api_key"] = apiKey
	}
	return SerpQuery{parameter: parameter}
}

// Set Serp API KEY
func setAPIKey(key string) {
	apiKey = key
}

// GetJSON returns SerpResponse containing
func (sq *SerpQuery) GetJSON() (SerpResponse, error) {
	rsp := sq.execute("/search", "json")
	return sq.decodeJSON(rsp.Body)
}

// GetHTML returns html as a string
func (sq *SerpQuery) GetHTML() (*string, error) {
	rsp := sq.execute("/search", "html")
	return sq.decodeHTML(rsp.Body)
}

// GetLocation returns closest location
func GetLocation(q string, limit int) (SerpResponseArray, error) {
	client := NewGoogleSearch(map[string]string{
		"q":     q,
		"limit": string(limit),
	})
	rsp := client.execute("/locations.json", "json")
	return client.decodeJSONArray(rsp.Body)
}

// GetAccount return account information
func GetAccount() (SerpResponse, error) {
	client := NewGoogleSearch(map[string]string{})
	rsp := client.execute("/account", "json")
	return client.decodeJSON(rsp.Body)
}

// GetSearchArchive retrieve search from the archive using the Search Archive API
func (sq *SerpQuery) GetSearchArchive(searchID string) (SerpResponse, error) {
	rsp := sq.execute("/searches/"+searchID+".json", "json")
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

// decodeJSONArray primitive function
func (sq *SerpQuery) decodeJSONArray(body io.ReadCloser) (SerpResponseArray, error) {
	decoder := json.NewDecoder(body)
	var rsp SerpResponseArray
	err := decoder.Decode(&rsp)
	if err != nil {
		return nil, errors.New("fail to decode array")
	}
	return rsp, nil
}

// decodeHTML primitive function
func (sq *SerpQuery) decodeHTML(body io.ReadCloser) (*string, error) {
	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	text := string(buffer)
	return &text, nil
}

// Execute the HTTP get
func (sq *SerpQuery) execute(path string, output string) *http.Response {
	query := url.Values{}
	for k, v := range sq.parameter {
		query.Add(k, v)
	}
	query.Add("source", "go")
	query.Add("output", output)
	endpoint := "https://serpapi.com" + path + "?" + query.Encode()
	var client = &http.Client{
		Timeout: time.Second * 60,
	}
	rsp, err := client.Get(endpoint)

	if err != nil {
		panic(err.Error())
	}
	return rsp
}
