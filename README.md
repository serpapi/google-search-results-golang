Google Search Results GoLang API
===

[![Build Status](https://travis-ci.org/serpapi/google-search-results-golang.svg?branch=master)](https://travis-ci.org/serpapi/google-search-results-golang)

This Golang package enables to scrape and parse Google, Bing and Baidu results using [SERP API](https://serpapi.com).
 Feel free to fork this repository to add more backends.

This project is an implementation of Serp API in Golang 1.12
There is no dependency for this project.

This Go module is meant to scrape and parse Google results using [SerpApi](https://serpapi.com).
The following services are provided:
 * [Search API](https://serpapi.com/search-api)
 * [Location API](https://serpapi.com/locations-api)
 * [Search Archive API](https://serpapi.com/search-archive-api)
 * [Account API](https://serpapi.com/account-api)

Serp API provides a [script builder](https://serpapi.com/demo) to get you started quickly.

An implementation example is provided here.
> demo/demo.go

You take a look to our test. 
> test/google_seach_results_test.go

[The full documentation is available here.](https://serpapi.com/search-api)

Feel free to fork this repository to add more backends.

## Installation

Go 1.12 must be already installed.
```bash
go get -u github.com/serpapi/google_search_results_golang
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

query := NewGoogleSearch(parameter)
serpResponse, err := query.json()
results := serpResponse["organic_results"].([]interface{})
first_result := results[0].(map[string]interface{})
fmt.Println(ref["title"].(string))
```

This example runs a search about "coffee" using your secret api key.

The Serp API service (backend)
 - searches on Google using the client: q = "coffee"
 - parses the messy HTML responses
 - return a standardizes JSON response
 - Format the request to Serp API server
 - Execute GET http request
 - Parse JSON into Ruby Hash using JSON standard library provided by Ruby
Et voila..

Alternatively, you can search:
 - Bing using NewBingSearch method
 - Baidu using NewBaiduSearch method

See the [playground to generate your code.](https://serpapi.com/playground)

## Example
 * [How to set SERP API key](#how-to-set-serp-api-key)
 * [Search API capability](#search-api-capability)
 * [Example by specification](#example-by-specification)
 * [Location API](#location-api)
 * [Search Archive API](#search-archive-api)
 * [Account API](#account-api)

### How to set SERP API key
The Serp API key can be set globally using a singleton pattern.

```go
GoogleSearchResults.serp_api_key_default = "Your Private Key"
client = GoogleSearchResults(parameter)
```

The Serp API key can be provided for each client.
```go
client = GoogleSearchResults(parameter, "Your Private Key")
```

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
  "serp_api_key": "Your SERP API Key",
  "tbm": "nws|isch|shop",
  "tbs": "custom to be search criteria",
  "async": true|false, # allow async 
  "output": "json|html" # output format
}

# define the search client
client := newGoogleSearch(parameter)

# override an existing parameter
client.parameter["location"] = "Portland,Oregon,United States"

# search format return as raw html
data, err := client.GetHTML()

# search format returns a json
data, err := client.GetJSON()
```

(the full documentation)[https://serpapi.com/search-api]

see below for more hands on examples.

### Example by specification

We love true open source, continuous integration and Test Drive Development (TDD). 
 We are using RSpec to test [our infrastructure around the clock](https://travis-ci.org/serpapi/google-search-results-ruby) to achieve the best QoS (Quality Of Service).
 
The directory test/ includes specification/examples.

Set your api key.
```bash
export API_KEY="your secret key"
```

```bash
make test
```

### Location API

```go
var locationList SerpResponseArray
var err error
locationList, err = GetLocation("Austin", 3)

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
  "api_key":  "your user key",
  "q":        "Coffee",
  "location": "Portland"
  }

client := NewGoogleSearch(parameter)
rsp, err := client.GetJSON()

if err != nil {
  log.Println("unexpected error", err)
}

searchID := rsp["search_metadata"].(map[string]interface{})["id"].(string)
searchArchive, err := client.GetSearchArchive(searchID)
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
var data SerpResponse
var err error
data, err = GetAccount()

if err != nil {
  log.Println(err)
  return
}
log.Println(data)
```
data contains the account information.

## Change log

 * 1.3 Add support for Bing and Baidu
 * 1.2 Export NewGoogleSearch outside of the package.

## Conclusion
Serp API supports Google Images, News, Shopping and more..
To enable a type of search, the field tbm (to be matched) must be set to:

 * isch: Google Images API.
 * nws: Google News API.
 * shop: Google Shopping API.
 * any other Google service should work out of the box.
 * (no tbm parameter): regular Google client.

The field `tbs` allows to customize the search even more.

[The full documentation is available here.](https://serpapi.com/search-api)

## Contributing

Contributions are welcome, feel to submit a pull request!

To run the tests:

```bash
make test
```
