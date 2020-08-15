package search

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// decodeJson response
func (search *Search) decodeJSON(body io.ReadCloser) (SearchResult, error) {
	// Decode JSON from response body
	decoder := json.NewDecoder(body)

	// Response data
	var rsp SearchResult
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

// decodeJSONArray decodes response body to SearchResultArray
func (search *Search) decodeJSONArray(body io.ReadCloser) (SearchResultArray, error) {
	decoder := json.NewDecoder(body)
	var rsp SearchResultArray
	err := decoder.Decode(&rsp)
	if err != nil {
		return nil, errors.New("fail to decode array")
	}
	return rsp, nil
}

// decodeHTML decodes response body to html string
func (search *Search) decodeHTML(body io.ReadCloser) (*string, error) {
	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	text := string(buffer)
	return &text, nil
}

// execute HTTP get reuqest and returns http response
func (search *Search) execute(path string, output string) (*http.Response, error) {
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
