package search

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

var apiKey string

func TestMain(m *testing.M) {
	apiKey = os.Getenv("API_KEY")
	if len(apiKey) == 0 {
		panic("API_KEY must be defined!")
	}
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func shoulSkip() bool {
	return len(apiKey) == 0
}

func TestGoogleQuickStart(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"q":             "Coffee",
		"location":      "Portland, Oregon, United States",
		"hl":            "en",
		"gl":            "us",
		"google_domain": "google.com",
		"safe":          "active",
		"start":         "10",
		"num":           "10",
		"device":        "desktop",
	}

	search := NewGoogleSearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error(err)
		return
	}
	result := rsp["organic_results"].([]interface{})[0].(map[string]interface{})
	if len(result["title"].(string)) == 0 {
		t.Error("empty title in local results")
		return
	}
}

// basic use case
func TestGoogleJSON(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Portland"}

	search := NewGoogleSearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
		t.Error("bad status")
		return
	}

	if len(rsp["organic_results"].([]interface{})) < 5 {
		t.Error("less than 5 organic result")
		return
	}
}

func TestBaiduJSON(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"q": "Coffee",
	}

	search := NewBaiduSearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
		t.Error("bad status")
		return
	}

	if len(rsp["organic_results"].([]interface{})) < 5 {
		t.Error("less than 5 organic result")
		return
	}
}

func TesNewtSearch(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"q": "Coffee",
	}

	search := NewSearch("baidu", parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
		t.Error("bad status")
		return
	}

	if len(rsp["organic_results"].([]interface{})) < 5 {
		t.Error("less than 5 organic result")
		return
	}
}

func TestBingJSON(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Portland"}

	search := NewBingSearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
		t.Error("bad status")
		return
	}

	// bing results are mostly advertising. not so great!
	if len(rsp["organic_results"].([]interface{})) < 3 {
		t.Error("less than 5 organic result")
		t.Error("response:")
		t.Error(rsp["search_metadata"])
		return
	}
}

func TestYandexJSON(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"text": "Coffee",
	}

	search := NewYandexSearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
		t.Error("bad status")
		return
	}

	if len(rsp["organic_results"].([]interface{})) < 5 {
		t.Error("less than 5 organic result")
		return
	}
}

func TestYahooJSON(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"p": "Coffee",
	}

	search := NewYahooSearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
		t.Error("bad status")
		return
	}

	if len(rsp["organic_results"].([]interface{})) < 5 {
		t.Error("less than 5 organic result")
		return
	}
}

func TestHomeDepotJSON(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"q": "Coffee",
	}

	search := NewHomeDepotSearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
		t.Error("bad status")
		return
	}

	if len(rsp["products"].([]interface{})) < 5 {
		t.Error("less than 5 organic result")
		return
	}
}

func TestEbayJSON(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"_nkw": "Coffee",
	}

	search := NewEbaySearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
		t.Error("bad status")
		return
	}

	if len(rsp["organic_results"].([]interface{})) < 5 {
		t.Error("less than 5 organic result")
		return
	}
}

func TestAppleStoreSearch(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"term": "Coffee",
	}

	search := NewAppleStoreSearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
		t.Error("bad status")
		return
	}

	if len(rsp["organic_results"].([]interface{})) < 5 {
		t.Error("less than 5 organic result")
		return
	}
}

func TestNaverSearch(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"query": "Coffee",
	}

	search := NewNaverSearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
		t.Error("bad status")
		return
	}

	if len(rsp["ads_results"].([]interface{})) < 5 {
		t.Error("less than 5 ads")
		return
	}
}

func TestGoogleJSONwithGlobalKey(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Portland"}

	search := NewGoogleSearch(parameter, apiKey)
	rsp, err := search.GetJSON()
	if err != nil {
		t.Error("unexpected error", err)
		return
	}
	result := rsp["organic_results"].([]interface{})[0].(map[string]interface{})
	if len(result["title"].(string)) == 0 {
		t.Error("empty title in local results")
		return
	}
}

func TestGoogleGetHTML(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Portland"}

	search := NewGoogleSearch(parameter, apiKey)
	data, err := search.GetHTML()
	if err != nil {
		t.Error("err must be nil")
		return
	}
	if !strings.Contains(*data, "</html>") {
		t.Error("data does not contains <html> tag")
	}
}

func TestGoogleDecodeJson(t *testing.T) {
	reader, err := os.Open("./data/search_coffee_sample.json")
	if err != nil {
		panic(err)
	}
	var search Search
	rsp, err := search.decodeJSON(reader)
	if err != nil {
		t.Error("error should be nil", err)
		return
	}

	results := rsp["organic_results"].([]interface{})
	ref := results[0].(map[string]interface{})
	if ref["title"] != "Portland Roasting Coffee" {
		t.Error("empty title in local results")
		return
	}
}

func TestGoogleDecodeJsonPage20(t *testing.T) {
	t.Log("run test")
	reader, err := os.Open("./data/search_coffee_sample_page20.json")
	if err != nil {
		panic(err)
	}
	var search Search
	rsp, err := search.decodeJSON(reader)
	if err != nil {
		t.Error("error should be nil")
		t.Error(err)
	}
	t.Log(reflect.ValueOf(rsp).MapKeys())
	results := rsp["organic_results"].([]interface{})
	ref := results[0].(map[string]interface{})
	t.Log(ref["title"].(string))
	if ref["title"].(string) != "Coffee | HuffPost" {
		t.Error("fail decoding the title ")
	}
}

func TestGoogleDecodeJsonError(t *testing.T) {
	reader, err := os.Open("./data/error_sample.json")
	if err != nil {
		panic(err)
	}
	var search Search
	rsp, err := search.decodeJSON(reader)
	if rsp != nil {
		t.Error("response should not be nil")
		return
	}

	if err == nil {
		t.Error("unexcepted err is nil")
	} else if strings.Compare(err.Error(), "Your account credit is too low, plesae add more credits.") == 0 {
		t.Error("empty title in local results")
		return
	}
}

func TestGoogleGetLocation(t *testing.T) {

	var rsp SearchResultArray
	var err error
	search := NewSearch("google", map[string]string{}, apiKey)
	rsp, err = search.GetLocation("Austin", 3)

	if err != nil {
		t.Error(err)
	}

	//log.Println(rsp[0])
	first := rsp[0].(map[string]interface{})
	googleID := first["google_id"].(float64)
	if googleID != float64(200635) {
		t.Error(googleID)
		return
	}
}

func TestGoogleGetAccount(t *testing.T) {
	// Skip this test
	if len(apiKey) == 0 {
		t.Skip("API_KEY required")
		return
	}

	var rsp SearchResult
	var err error
	search := NewSearch("google", map[string]string{}, apiKey)
	rsp, err = search.GetAccount()

	if err != nil {
		t.Error("fail to fetch data")
		t.Error(err)
		return
	}

	if rsp["account_id"] == nil {
		t.Error("no account_id found")
		return
	}
}

// Search archive API
func TestGoogleSearchArchive(t *testing.T) {
	if len(apiKey) == 0 {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Portland"}

	search := NewGoogleSearch(parameter, apiKey)
	rsp, err := search.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	searchID := rsp["search_metadata"].(map[string]interface{})["id"].(string)
	if len(searchID) == 0 {
		t.Error("search_metadata.id must be defined")
		return
	}

	searchArchive, err := search.GetSearchArchive(searchID)
	if err != nil {
		t.Error(err)
		return
	}

	searchIDArchive := searchArchive["search_metadata"].(map[string]interface{})["id"].(string)
	if searchIDArchive != searchID {
		t.Error("search_metadata.id do not match", searchIDArchive, searchID)
	}
}
