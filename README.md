# Google Search Results GoLang API

![test](https://github.com/serpapi/google-search-results-golang/workflows/Go/badge.svg)

This Golang package enables to scrape and parse results from Google, Bing, Baidu, Yahoo, Yandex, Ebay, Google Schoolar and more using [SerpApi](https://serpapi.com).
 
This project is an implementation of SerpApi in Golang 1.12
There is no dependency for this project.

This Go module is meant to scrape and parse Google results using [SerpApi](https://serpapi.com).
The following services are provided:
 * [Search API](https://serpapi.com/search-api)
 * [Location API](https://serpapi.com/locations-api)
 * [Search Archive API](https://serpapi.com/search-archive-api)
 * [Account API](https://serpapi.com/account-api)

SerpApi provides a [script builder](https://serpapi.com/demo) to get you started quickly.

An implementation example is provided here.
> demo/demo.go

You take a look to our test. 
> test/google_seach_results_test.go

[The full documentation is available here.](https://serpapi.com/search-api)


## Installation

Go 1.10+ must be already installed.
```bash
go get -u github.com/serpapi/google-search-results-golang
```

## Quick start

```go
package main

import (
	"fmt"
	g "github.com/serpapi/google-search-results-golang"
)

parameter := map[string]string{
    "q":            "Coffee",
    "location":     "Portland"
}

query := NewGoogleSearch(parameter, api)
// Many search engine available: bing, yahoo, baidu, googemaps, googleproduct, googlescholar, ebay, walmart, youtube..

rsp, err := query.json()
results := rsp["organic_results"].([]interface{})
first_result := results[0].(map[string]interface{})
fmt.Println(ref["title"].(string))
```

This example runs a search about "coffee" using your secret api key.

The SerpApi service (backend)
 - searches on Google using the search: q = "coffee"
 - parses the messy HTML responses
 - return a standardizes JSON response
 - Format the request
 - Execute GET http request against SerpApi service
 - Parse JSON response into a deep hash
Et voila..

Alternatively, you can search:
 - Bing using NewBingSearch method
 - Baidu using NewBaiduSearch method

See the [playground to generate your code.](https://serpapi.com/playground)

## Example
 * [Search API capability](#search-api-capability)
 * [Example by specification](#example-by-specification)
 * [Location API](#location-api)
 * [Search Archive API](#search-archive-api)
 * [Account API](#account-api)

### Search API capability
```go
parameter = {
  "q": "query",
  "google_domain": "Google Domain",
  "location": "Location Requested",
  "device": device,
  "hl": "Google UI Language",
  "gl": "Google Country",
  "safe": "Safe Search Flag",
  "num": "Number of Results",
  "start": "Pagination Offset",
  "api_key": "Your SERP API Key",
  "tbm": "nws|isch|shop",
  "tbs": "custom to be search criteria",
  "async": true|false,  // allow async 
  "output": "json|html" // output format
}

// api_key from https://serpapi.com/dashboard
api_key := "your personal API key"

// set search engine: google|yahoo|bing|ebay|yandex
engine := "yahoo"

// define the search search
search := NewSearch(engine, parameter, api_key)

// override an existing parameter
search.parameter["location"] = "Portland,Oregon,United States"

// search format return as raw html
data, err := search.GetHTML()

// search format returns a json
data, err := search.GetJSON()
```

(the full documentation)[https://serpapi.com/search-api]

Full example: [https://github.com/serpapi/google-search-results-golang/blob/master/demo/demo.go]

see below for more hands on examples.

### Location API

```go
var locationList SerpResponseArray
var err error
locationList, err = search.GetLocation("Austin", 3)

if err != nil {
  log.Println(err)
}
log.Println(locationList)
```
rsp contains the first 3 location matching Austin (Texas, Texas, Rochester)

### Search Archive API

Run a search then get search result from the archive using the search archive API.
```go
parameter := map[string]string{
  "q":        "Coffee",
  "location": "Portland"
  }

search := NewGoogleSearch(parameter, "your user key")
rsp, err := search.GetJSON()

if err != nil {
  log.Println("unexpected error", err)
}

searchID := rsp["search_metadata"].(map[string]interface{})["id"].(string)
searchArchive, err := search.GetSearchArchive(searchID)
if err != nil {
  log.Println(err)
  return
}

searchIDArchive := searchArchive["search_metadata"].(map[string]interface{})["id"].(string)
if searchIDArchive != searchID {
  log.Println("search_metadata.id do not match", searchIDArchive, searchID)
}

log.Println(searchIDArchive)
```
it prints the search ID from the archive.

### Account API
```go
var data SearchResult
var err error
data, err = search.GetAccount()

if err != nil {
  log.Println(err)
  return
}
log.Println(data)
```
data contains the account information.

### Example by specification

We love true open source, continuous integration and Test Drive Development (TDD). 
 We are using "go test" to test our infrastructure around the clock
  to achieve the best QoS (Quality Of Service).
 
The directory test/ includes specification/examples.

To run the test from bash using make
```bash
export API_KEY="your secret key"
make test
```

### Error management

This library follows the basic error management solution provided by Go.
 A simple error is returned in case something goes wrong. 
 The error wraps a simple error message.
 
## Change log
 * 3.2
   - add naver search
   - add apple store search
 * 3.1
   - Add home depot search engine
 * 3.0
   - Naming convention change.
   - Rename Client to Search
   - Fix lint issue
   - Add walmart and youtube
 * 2.1 
    - Add support for Yandex, Ebay, Yahoo
    - create HTTP search only once per SerpApiClient
 * 2.0 Rewrite fully the implementation
        to be more scalable in order to support multiple engines.
 * 1.3 Add support for Bing and Baidu
 * 1.2 Export NewGoogleSearch outside of the package.

## Conclusion

SerpApi supports mutiple search engines and subservices all available for this Golang search.

For example: Using Google search.
To enable a type of search, the field tbm (to be matched) must be set to:

 * isch: Google Images API.
 * nws: Google News API.
 * shop: Google Shopping API.
 * any other Google service should work out of the box.
 * (no tbm parameter): regular Google search.

The field `tbs` allows to customize the search even more.

[The full documentation is available here.](https://serpapi.com/search-api)

## Contributing

Contributions are welcome, feel to submit a pull request!

To run the tests:

```bash
make test
```
