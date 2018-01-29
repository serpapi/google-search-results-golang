package google_search_results
// https://mholt.github.io/json-to-go/

/*
 * Constant and model declaration
 */
const BACKEND = "serpapi.com"
const SERP_API_KEY = "serp_api_key"

// Serp API query
type SerpQuery struct {
	parameter map[string]string
}

// Serp API Response
type SerpResponse struct {
	Error string `json:"error,omitempty"`

	SearchInformation struct {
	 TotalResults int     `json:"total_results"`
	 TimeTaken    float64 `json:"time_taken"`
	 Query        string  `json:"query"`
	 Location     string  `json:"location"`
  } `json:"search_information"`

	SerpAPIData struct {
	 ID struct {
		Oid string `json:"$oid"`
	 } `json:"id"`
		TotalTimeTaken float64 `json:"total_time_taken"`
	 	GoogleURL      string `json:"google_url"`
  } `json:"serp_api_data"`

	Map struct {
	 Link  string `json:"link"`
	 Image bool   `json:"image"`
 } `json:"map"`

	LocalResults         []struct {
		Position    int     `json:"position"`
		Title       string  `json:"title"`
		Rating      float64 `json:"rating"`
		Reviews     int     `json:"reviews"`
		Price       string  `json:"price,omitempty"`
		Type        string  `json:"type"`
		Address     string  `json:"address"`
		Description string  `json:"description"`
		Hours       string  `json:"hours"`
		Thumbnail   bool    `json:"thumbnail"`
	} `json:"local_results"`

	KnowledgeGraph struct {
		 Title                   string `json:"title"`
		 Type                    string `json:"type"`
		 Image                   bool   `json:"image"`
		 Description             string `json:"description"`
		 Source struct {
			 Name string `json:"name"`
			 Link string `json:"link"`
		 } `json:"source"`

		 RelatedSearches []struct {
			 Name  string `json:"name"`
			 Link  string `json:"link"`
			 Image bool   `json:"image"`
		 } `json:"related_searches"`

		 MoreRelatedSearchesLink string `json:"more_related_searches_link"`
	 } `json:"knowledge_graph"`

	AnswerBox struct {
		 Type string `json:"type"`
	 } `json:"answer_box"`

	OrganicResults       []struct {
		Position      int    `json:"position"`
		Title         string `json:"title"`
		Link          string `json:"link"`
		DisplayedLink string `json:"displayed_link"`
		Snippet       string `json:"snippet"`
		RelatedLink   string `json:"related_link,omitempty"`
		Date          string `json:"date,omitempty"`
		Sitelinks     struct {
			Inline []struct {
				Title string `json:"title"`
				Link  string `json:"link"`
			} `json:"inline"`
		} `json:"sitelinks,omitempty"`
		CachedLink    string `json:"cached_link,omitempty"`
	} `json:"organic_results"`

	RelatedPlaceSearches []struct {
		Query   string `json:"query"`
		Link    string `json:"link"`
		Snippet string `json:"snippet"`
	} `json:"related_place_searches"`

	RelatedSearches      []struct {
		Query string `json:"query"`
		Link  string `json:"link"`
	} `json:"related_searches"`

	Pagination struct {
		 Current    int    `json:"current"`
		 Next       string `json:"next"`
		 OtherPages struct {
			Num2  string `json:"2"`
			Num3  string `json:"3"`
			Num4  string `json:"4"`
			Num5  string `json:"5"`
			Num6  string `json:"6"`
			Num7  string `json:"7"`
			Num8  string `json:"8"`
			Num9  string `json:"9"`
			Num10 string `json:"10"`
		} `json:"other_pages"`
	 } `json:"pagination"`
}
