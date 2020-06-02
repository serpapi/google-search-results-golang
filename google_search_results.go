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

// SerpApiClient hold query parameter
type SerpApiClient struct {
	Engine     string
	Parameter  map[string]string
	ApiKey     string
	HttpClient *http.Client
}

// SerpApiResponse hold response
type SerpApiResponse map[string]interface{}

// SerpApiResponseArray hold response array
type SerpApiResponseArray []interface{}

const (
	// Version
	VERSION = "2.1.0"
)

// NewSerpApiClient create generic SerpApi client which support any search engine
func NewSerpApiClient(engine string, parameter map[string]string, apiKey string) SerpApiClient {
	// Create the http client
	httpClient := &http.Client{
		Timeout: time.Second * 60,
	}
	return SerpApiClient{Engine: engine, Parameter: parameter, ApiKey: apiKey, HttpClient: httpClient}
}

// NewGoogleClient create client for google
func NewGoogleClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("google", parameter, apiKey)
}

// NewBingClient create client for bing
func NewBingClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("bing", parameter, apiKey)
}

// NewBaiduClient create client for baidu
func NewBaiduClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("baidu", parameter, apiKey)
}

// NewYahooClient create client for yahoo
func NewYahooClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("yahoo", parameter, apiKey)
}

// NewGoogleMapsClient create client for google_maps
func NewGoogleMapsClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("google_maps", parameter, apiKey)
}

// NewGoogleProductClient create client for google_product
func NewGoogleProductClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("google_product", parameter, apiKey)
}

// NewGoogleScholarClient create client for google_product
func NewGoogleScholarClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("google_scholar", parameter, apiKey)
}

// NewYandexClient create client for yandex
func NewYandexClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("yandex", parameter, apiKey)
}

// NewEbayClient create client for ebay
func NewEbayClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("ebay", parameter, apiKey)
}

// SetApiKey globaly set api_key
func (client *SerpApiClient) SetApiKey(key string) {
	client.ApiKey = key
}

// GetJSON returns SerpApiResponse containing
func (client *SerpApiClient) GetJSON() (SerpApiResponse, error) {
	rsp, err := client.execute("/search", "json")
	if err != nil {
		return nil, err
	}
	return client.decodeJSON(rsp.Body)
}

// GetHTML returns html as a string
func (client *SerpApiClient) GetHTML() (*string, error) {
	rsp, err := client.execute("/search", "html")
	if err != nil {
		return nil, err
	}
	return client.decodeHTML(rsp.Body)
}

// GetLocation returns closest location
func (client *SerpApiClient) GetLocation(q string, limit int) (SerpApiResponseArray, error) {
	client.Parameter = map[string]string{
		"q":     q,
		"limit": string(limit),
	}
	rsp, err := client.execute("/locations.json", "json")
	if err != nil {
		return nil, err
	}
	return client.decodeJSONArray(rsp.Body)
}

// GetAccount return account information
func (client *SerpApiClient) GetAccount() (SerpApiResponse, error) {
	client.Parameter = map[string]string{}
	rsp, err := client.execute("/account", "json")
	if err != nil {
		return nil, err
	}
	return client.decodeJSON(rsp.Body)
}

// GetSearchArchive retrieve search from the archive using the Search Archive API
func (client *SerpApiClient) GetSearchArchive(searchID string) (SerpApiResponse, error) {
	rsp, err := client.execute("/searches/"+searchID+".json", "json")
	if err != nil {
		return nil, err
	}
	return client.decodeJSON(rsp.Body)
}

// decodeJson response
func (client *SerpApiClient) decodeJSON(body io.ReadCloser) (SerpApiResponse, error) {
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
func (client *SerpApiClient) decodeJSONArray(body io.ReadCloser) (SerpApiResponseArray, error) {
	decoder := json.NewDecoder(body)
	var rsp SerpApiResponseArray
	err := decoder.Decode(&rsp)
	if err != nil {
		return nil, errors.New("fail to decode array")
	}
	return rsp, nil
}

// decodeHTML decodes response body to html string
func (client *SerpApiClient) decodeHTML(body io.ReadCloser) (*string, error) {
	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	text := string(buffer)
	return &text, nil
}

// execute HTTP get reuqest and returns http response
func (client *SerpApiClient) execute(path string, output string) (*http.Response, error) {
	query := url.Values{}
	if client.Parameter != nil {
		for k, v := range client.Parameter {
			query.Add(k, v)
		}
	}

	// api_key
	if len(client.ApiKey) != 0 {
		query.Add("api_key", client.ApiKey)
	}

	// engine
	if len(query.Get("engine")) == 0 {
		query.Set("engine", client.Engine)
	}

	// source programming language
	query.Add("source", "go")

	// set output
	query.Add("output", output)

	endpoint := "https://serpapi.com" + path + "?" + query.Encode()
	rsp, err := client.HttpClient.Get(endpoint)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
