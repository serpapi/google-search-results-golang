package google_search_results

import (
	"testing"
	"os"
	"strings"
	"reflect"
)

func TestRealWorldExample(t *testing.T) {
	parameter := map[string]string{
		"q": "Coffee",
		"location": "Portland, Oregon, United States",
		"hl": "en",
		"gl": "us",
		"google_domain": "google.com",
		"api_key": "demo",
		"safe": "active",
		"start": "10",
		"num": "10",
		"device": "desktop",
	}

	query := newGoogleSearch(parameter)
	serpResponse, err := query.json()

	if err != nil {
		t.Error("unexpected error")
	}
	result := serpResponse["organic_results"].([]interface{})[0].(map[string]interface{})
	if len(result["title"].(string)) == 0 {
		t.Error("empty title in local results")
	}
}

// basic use case
func TestJSON(t *testing.T) {
	parameter := map[string]string{
		"serp_api_key": "demo",
		"q":            "Coffee",
		"location":     "Portland"}

	query := newGoogleSearch(parameter)
	serpResponse, err := query.json()

	if err != nil {
		t.Error("unexpected error")
	}
	result := serpResponse["organic_results"].([]interface{})[0].(map[string]interface{})
	if len(result["title"].(string)) == 0 {
		t.Error("empty title in local results")
	}
}

func TestJSONwithGlobalKey(t *testing.T) {
	parameter := map[string]string{
		"q":            "Coffee",
		"location":     "Portland"}

	setApiKey("demo")

	query := newGoogleSearch(parameter)

	serpResponse, err := query.json()
	if err != nil {
		t.Error("unexpected error")
	}

	result := serpResponse["organic_results"].([]interface{})[0].(map[string]interface{})
	if len(result["title"].(string)) == 0 {
		t.Error("empty title in local results")
	}
}

func TestGetHTML(t *testing.T) {
	parameter := map[string]string{
		"q":            "Coffee",
		"location":     "Portland"}

	setApiKey("demo")

	query := newGoogleSearch(parameter)
	data, err  := query.html()
	if err != nil {
		t.Error("err must be nil")
	}
	if !strings.Contains(*data, "</html>") {
		t.Error("data does not contains <html> tag")
	}
}

func TestDecodeJson(t *testing.T) {
	reader, err := os.Open("./data/search_coffee_sample.json")
	if err != nil {
		panic(err)
	}
	var sq SerpQuery
	serpResponse, serpError := sq.decodeJson(reader)
	if serpError != nil {
		t.Error("error should be nil")
	}

	results := serpResponse["organic_results"].([]interface{})
	ref := results[0].(map[string]interface{})
	if ref["title"] != "Portland Roasting Coffee"{
		t.Error("empty title in local results")
	}
}

func TestDecodeJsonPage20(t *testing.T) {
	t.Log("run test")
	reader, err := os.Open("./data/search_coffee_sample_page20.json")
	if err != nil {
		panic(err)
	}
	var sq SerpQuery
	serpResponse, serpError := sq.decodeJson(reader)
	if serpError != nil {
		t.Error("error should be nil")
		t.Error(serpError)
	}
	t.Log(reflect.ValueOf(serpResponse).MapKeys())
	results := serpResponse["organic_results"].([]interface{})
	ref := results[0].(map[string]interface{})
	t.Log(ref["title"].(string))
	if ref["title"].(string) != "Coffee | HuffPost" {
		t.Error("fail decoding the title ")
	}
}

func TestDecodeJsonError(t *testing.T) {
	reader, err := os.Open("./data/error_sample.json")
	if err != nil {
		panic(err)
	}
	var sq SerpQuery
	serpResponse, serpError := sq.decodeJson(reader)
	if serpResponse != nil {
		t.Error("response should be nil")
	}

	if serpError == nil {
		t.Error("unexcepted serpError is nil")
	} else if strings.Compare(serpError.Error(), "Your account credit is too low, plesae add more credits.") == 0 {
		t.Error("empty title in local results")
	}
}
