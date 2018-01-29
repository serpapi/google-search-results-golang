// This package enales to interact with SerpAPI server
package google_search_results

import(
	"net/http"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"io"
	"errors"
)

var serpApiKey string

// Create a new query
func newGoogleSearch(parameter map[string]string) SerpQuery {
	if len(serpApiKey) > 0 {
		parameter[SERP_API_KEY] = serpApiKey
	}

	return SerpQuery{ parameter: parameter}
}

// Set Serp API key
func setApiKey(key string) {
	serpApiKey = key
}

// Execute the query
func (sq *SerpQuery) execute(output string) *http.Response {
	query := url.Values{}
	for k,v := range sq.parameter {
		query.Add(k, v)
	}
	query.Add("source", "go")
	query.Add("output", output)
	endpoint := "https://" + BACKEND + "/search?" + query.Encode()
	rsp, err := http.Get(endpoint)

	if err != nil {
		panic(err.Error())
	}
	return rsp
}

// return go struct by processing the json returned by the server
func (sq *SerpQuery) json() (*SerpResponse, error) {
	rsp := sq.execute("json")
	return sq.decodeJson(rsp.Body)
}

func (sq *SerpQuery) jsonWithImages() (*SerpResponse, error) {
	rsp := sq.execute("json_with_images")
	return sq.decodeJson(rsp.Body)
}

// decode json response
func (sq *SerpQuery) decodeJson(body io.ReadCloser) (*SerpResponse, error) {
	// Decode JSON from response body
	decoder := json.NewDecoder(body)

	var serpResponse SerpResponse
	err := decoder.Decode(&serpResponse)
	if err != nil || len(serpResponse.Error) > 0 {
		if err != nil {
			panic(err)
			return nil, err
		}
		return nil, errors.New(serpResponse.Error)
	}
	return &serpResponse, nil
}

// return html as string
func (sq *SerpQuery) html() (*string, error) {
	rsp := sq.execute("html")
	return sq.decodeHtml(rsp.Body)
}

// decode html
func (sq *SerpQuery) decodeHtml(body io.ReadCloser) (*string, error) {
	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
		return nil, err
	}
	text := string(buffer)
	return &text, nil
}