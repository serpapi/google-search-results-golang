package search

/*
 * This package enables to interact with SerpApi search API
 */

import (
	"fmt"
	"net/http"
	"time"
)

// NewSearch create generic search which support any search engine
func NewSearch(engine string, parameter map[string]string, apiKey string) Search {
	// Create the http search
	httpSearch := &http.Client{
		Timeout: time.Second * 60,
	}
	return Search{Engine: engine, Parameter: parameter, ApiKey: apiKey, HttpSearch: httpSearch}
}

// NewGoogleSearch creates search for google
func NewGoogleSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("google", parameter, apiKey)
}

// NewBingSearch creates search for bing
func NewBingSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("bing", parameter, apiKey)
}

// NewBaiduSearch creates search for baidu
func NewBaiduSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("baidu", parameter, apiKey)
}

// NewYahooSearch creates search for yahoo
func NewYahooSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("yahoo", parameter, apiKey)
}

// NewGoogleMapsSearch creates search for google_maps
func NewGoogleMapsSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("google_maps", parameter, apiKey)
}

// NewGoogleProductSearch creates search for google_product
func NewGoogleProductSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("google_product", parameter, apiKey)
}

// NewGoogleScholarSearch creates search for google_product
func NewGoogleScholarSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("google_scholar", parameter, apiKey)
}

// NewYandexSearch creates search for yandex
func NewYandexSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("yandex", parameter, apiKey)
}

// NewEbaySearch creates search for ebay
func NewEbaySearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("ebay", parameter, apiKey)
}

// NewYoutubeSearch creates search for ebay
func NewYoutubeSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("youtube", parameter, apiKey)
}

// NewWalmartSearch creates search for ebay
func NewWalmartSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("walmart", parameter, apiKey)
}

// NewHomeDepotSearch creates search for ebay
func NewHomeDepotSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("home_depot", parameter, apiKey)
}

// NewNaverSearch creates search for Naver search engine
func NewNaverSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("naver", parameter, apiKey)
}

// NewAppleStoreSearch creates search for Apple store (itunes.apple.com)
func NewAppleStoreSearch(parameter map[string]string, apiKey string) Search {
	return NewSearch("apple_app_store", parameter, apiKey)
}

// SetApiKey globaly set api_key
func (search *Search) SetApiKey(key string) {
	search.ApiKey = key
}

// GetJSON returns SearchResult containing
func (search *Search) GetJSON() (SearchResult, error) {
	rsp, err := search.execute("/search", "json")
	if err != nil {
		return nil, err
	}
	return search.decodeJSON(rsp.Body)
}

// GetHTML returns html as a string
func (search *Search) GetHTML() (*string, error) {
	rsp, err := search.execute("/search", "html")
	if err != nil {
		return nil, err
	}
	return search.decodeHTML(rsp.Body)
}

// GetLocation returns the standardize location takes location and limit as input.
func (search *Search) GetLocation(location string, limit int) (SearchResultArray, error) {
	search.Parameter = map[string]string{
		"q":     location,
		"limit": fmt.Sprint(limit),
	}
	rsp, err := search.execute("/locations.json", "json")
	if err != nil {
		return nil, err
	}
	return search.decodeJSONArray(rsp.Body)
}

// GetAccount return account information
func (search *Search) GetAccount() (SearchResult, error) {
	search.Parameter = map[string]string{}
	rsp, err := search.execute("/account", "json")
	if err != nil {
		return nil, err
	}
	return search.decodeJSON(rsp.Body)
}

// GetSearchArchive retrieve search from the archive using the Search Archive API
func (search *Search) GetSearchArchive(searchID string) (SearchResult, error) {
	rsp, err := search.execute("/searches/"+searchID+".json", "json")
	if err != nil {
		return nil, err
	}
	return search.decodeJSON(rsp.Body)
}
