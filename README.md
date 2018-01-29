Google search result GoLang API
===

This Golang package enables to scrape and parse Google results using [SERP API](https://serpapi.com).
 Feel free to fork this repository to add more backends.

This project is an implementation of Serp API in Golang 1.8.
There is no dependency for this project.

An implementation example is provided here.
@see google_seach_results_test.go
## Simple Example
```java
parameter := map[string]string{
    "q":            "Coffee",
    "location":     "Portland"
}

query := newGoogleSearch(parameter)
serpResponse, err = query.json()
fmt.Println(serpResponse.LocalResults[0].Title)
```

## Set parameter
```java
Map<String, String> parameter = new HashMap<>();
parameter.put("q", "Coffee");
parameter.put("location", "Portland");
```

## Set SERP API key

```java
GoogleSearchResults.serp_api_key_default = "Your Private Key"
query = GoogleSearchResults(parameter)
```
Or

```java
query = GoogleSearchResults(parameter, "Your Private Key")
```

## Example with all params and all outputs

```java
query_parameter = {
  "q": "query",
  "google_domain": "Google Domain",
  "location": "Location Requested",
  "device": device,
  "hl": "Google UI Language",
  "gl": "Google Country",
  "safe": "Safe Search Flag",
  "num": "Number of Results",
  "start": "Pagination Offset",
  "serp_api_key": "Your SERP API Key"
}

query := newGoogleSearch(query_parameter)
data, err := query.html()
data, err := query.json()
data, err := query.jsonWithImages()
```

Limitation
---
 - No wrapper method around parametrization.
  the parameters are passed by a simple Map<String,String>

Source
---
 * http://www.baeldung.com/java-http-request
 * https://github.com/google/gson

@author Victor Benarbia