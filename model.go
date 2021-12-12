package search

import "net/http"

// Search holds query
type Search struct {
	Engine     string
	Parameter  map[string]string
	ApiKey     string
	HttpSearch *http.Client
}

// SearchResult holds response
type SearchResult map[string]interface{}

// SearchResultArray hold response array
type SearchResultArray []interface{}

const (
	// Current version
	VERSION = "3.2.0"
)
