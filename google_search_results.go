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

// SerpApiSearch hold query parameter
type SerpApiSearch struct {
	Engine     string
	Parameter  map[string]string
	ApiKey     string
	HttpSearch *http.Client
}

// SerpApiResponse hold response
type SerpApiResponse map[string]interface{}

// SerpApiResponseArray hold response array
type SerpApiResponseArray []interface{}

const (
	// Version
	VERSION = "2.1.0"
)

// NewSerpApiSearch create generic SerpApi search which support any search engine
func NewSerpApiSearch(engine string, parameter map[string]string, apiKey string) SerpApiSearch {
	// Create the http search
	httpSearch := &http.Client{
		Timeout: time.Second * 60,
	}
	return SerpApiSearch{Engine: engine, Parameter: parameter, ApiKey: apiKey, HttpSearch: httpSearch}
}

// NewGoogleSearch create search for google
func NewGoogleSearch(parameter map[string]string, apiKey string) SerpApiSearch {
	return NewSerpApiSearch("google", parameter, apiKey)
}

// NewBingSearch create search for bing
func NewBingSearch(parameter map[string]string, apiKey string) SerpApiSearch {
	return NewSerpApiSearch("bing", parameter, apiKey)
}

// NewBaiduSearch create search for baidu
func NewBaiduSearch(parameter map[string]string, apiKey string) SerpApiSearch {
	return NewSerpApiSearch("baidu", parameter, apiKey)
}

// NewYahooSearch create search for yahoo
func NewYahooSearch(parameter map[string]string, apiKey string) SerpApiSearch {
	return NewSerpApiSearch("yahoo", parameter, apiKey)
}

// NewGoogleMapsSearch create search for google_maps
func NewGoogleMapsSearch(parameter map[string]string, apiKey string) SerpApiSearch {
	return NewSerpApiSearch("google_maps", parameter, apiKey)
}

// NewGoogleProductSearch create search for google_product
func NewGoogleProductSearch(parameter map[string]string, apiKey string) SerpApiSearch {
	return NewSerpApiSearch("google_product", parameter, apiKey)
}

// NewGoogleScholarSearch create search for google_product
func NewGoogleScholarSearch(parameter map[string]string, apiKey string) SerpApiSearch {
	return NewSerpApiSearch("google_scholar", parameter, apiKey)
}

// NewYandexSearch create search for yandex
func NewYandexSearch(parameter map[string]string, apiKey string) SerpApiSearch {
	return NewSerpApiSearch("yandex", parameter, apiKey)
}

// NewEbaySearch create search for ebay
func NewEbaySearch(parameter map[string]string, apiKey string) SerpApiSearch {
	return NewSerpApiSearch("ebay", parameter, apiKey)
}

// SetApiKey globaly set api_key
func (search *SerpApiSearch) SetApiKey(key string) {
	search.ApiKey = key
}

// GetJSON returns SerpApiResponse containing
func (search *SerpApiSearch) GetJSON() (SerpApiResponse, error) {
	rsp, err := search.execute("/search", "json")
	if err != nil {
		return nil, err
	}
	return search.decodeJSON(rsp.Body)
}

// GetHTML returns html as a string
func (search *SerpApiSearch) GetHTML() (*string, error) {
	rsp, err := search.execute("/search", "html")
	if err != nil {
		return nil, err
	}
	return search.decodeHTML(rsp.Body)
}

// GetLocation returns closest location
func (search *SerpApiSearch) GetLocation(q string, limit int) (SerpApiResponseArray, error) {
	search.Parameter = map[string]string{
		"q":     q,
		"limit": string(limit),
	}
	rsp, err := search.execute("/locations.json", "json")
	if err != nil {
		return nil, err
	}
	return search.decodeJSONArray(rsp.Body)
}

// GetAccount return account information
func (search *SerpApiSearch) GetAccount() (SerpApiResponse, error) {
	search.Parameter = map[string]string{}
	rsp, err := search.execute("/account", "json")
	if err != nil {
		return nil, err
	}
	return search.decodeJSON(rsp.Body)
}

// GetSearchArchive retrieve search from the archive using the Search Archive API
func (search *SerpApiSearch) GetSearchArchive(searchID string) (SerpApiResponse, error) {
	rsp, err := search.execute("/searches/"+searchID+".json", "json")
	if err != nil {
		return nil, err
	}
	return search.decodeJSON(rsp.Body)
}

// decodeJson response
func (search *SerpApiSearch) decodeJSON(body io.ReadCloser) (SerpApiResponse, error) {
	// Decode JSON from response body
	decoder := json.NewDecoder(body)

	// Response data
	var rsp SerpApiResponse
	err := decoder.Decode(&rsp)
	if err != nil {
		return nil, errors.New("fail to decode")
	}

	// check error message
	errorMessage, derror := rsp["error"].(string)
	if derror {
		return nil, errors.New(errorMessage)
	}
	return rsp, nil
}

// decodeJSONArray decodes response body to SerpApiResponseArray
func (search *SerpApiSearch) decodeJSONArray(body io.ReadCloser) (SerpApiResponseArray, error) {
	decoder := json.NewDecoder(body)
	var rsp SerpApiResponseArray
	err := decoder.Decode(&rsp)
	if err != nil {
		return nil, errors.New("fail to decode array")
	}
	return rsp, nil
}

// decodeHTML decodes response body to html string
func (search *SerpApiSearch) decodeHTML(body io.ReadCloser) (*string, error) {
	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	text := string(buffer)
	return &text, nil
}

// execute HTTP get reuqest and returns http response
func (search *SerpApiSearch) execute(path string, output string) (*http.Response, error) {
	query := url.Values{}
	if search.Parameter != nil {
		for k, v := range search.Parameter {
			query.Add(k, v)
		}
	}

	// api_key
	if len(search.ApiKey) != 0 {
		query.Add("api_key", search.ApiKey)
	}

	// engine
	if len(query.Get("engine")) == 0 {
		query.Set("engine", search.Engine)
	}

	// source programming language
	query.Add("source", "go")

	// set output
	query.Add("output", output)

	endpoint := "https://serpapi.com" + path + "?" + query.Encode()
	rsp, err := search.HttpSearch.Get(endpoint)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
