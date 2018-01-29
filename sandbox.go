package google_search_results


import (
	"fmt"
	"net/url"
)

func main(){
	parameter := map[string]string{
		"app_id": "you_api",
		"app_sign":"md5_base_16",
		"timestamp":"1473655478000"}

	fmt.Println(parameter)

	q := url.Values{}
		for k,v := range parameter{
			q.Add(k, v)
		}

	fmt.Println(q.Encode())
}