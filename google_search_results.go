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
	Engine    string
	Parameter map[string]string
	ApiKey    string
}

// SerpApiResponse hold response
type SerpApiResponse map[string]interface{}

// SerpApiResponseArray hold response array
type SerpApiResponseArray []interface{}

// NewSerpApiClient create generic SerpApi client
func NewSerpApiClient(engine string, parameter map[string]string, apiKey string) SerpApiClient {
	return SerpApiClient{Engine: engine, Parameter: parameter, ApiKey: apiKey}
}

// NewGoogleSearch create client for google
func NewGoogleClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("google", parameter, apiKey)
}

// NewBingSearch create client for bing
func NewBingClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("bing", parameter, apiKey)
}

// NewBaiduSearch create client for baidu
func NewBaiduClient(parameter map[string]string, apiKey string) SerpApiClient {
	return NewSerpApiClient("baidu", parameter, apiKey)
}

// NewBaiduSearch create client for yahoo
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

// Set your API KEY
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
	var serpApiResponse SerpApiResponse
	err := decoder.Decode(&serpApiResponse)
	if err != nil {
		return nil, errors.New("fail to decode")
	}

	// check error message
	errorMessage, derror := serpApiResponse["error"].(string)
	if derror {
		return nil, errors.New(errorMessage)
	}
	return serpApiResponse, nil
}

// decodeJSONArray primitive function
func (client *SerpApiClient) decodeJSONArray(body io.ReadCloser) (SerpApiResponseArray, error) {
	decoder := json.NewDecoder(body)
	var rsp SerpApiResponseArray
	err := decoder.Decode(&rsp)
	if err != nil {
		return nil, errors.New("fail to decode array")
	}
	return rsp, nil
}

// decodeHTML primitive function
func (client *SerpApiClient) decodeHTML(body io.ReadCloser) (*string, error) {
	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	text := string(buffer)
	return &text, nil
}

// Execute the HTTP get
func (client *SerpApiClient) execute(path string, output string) (*http.Response, error) {
	query := url.Values{}
	if client.Parameter != nil {
		for k, v := range client.Parameter {
			query.Add(k, v)
		}
	}
	if len(client.ApiKey) != 0 {
		query.Add("api_key", client.ApiKey)
	}
	query.Add("source", "go")
	query.Add("output", output)
	endpoint := "https://serpapi.com" + path + "?" + query.Encode()
	var httpClient = &http.Client{
		Timeout: time.Second * 60,
	}
	rsp, err := httpClient.Get(endpoint)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
