package reddit

type ResponseData struct {
	Data struct {
		Children []struct {
			Data struct {
				Selftext   string `json:"selftext"`
				Title      string `json:"title"`
				Url        string `json:"url"`
				CreatedUtc int    `json:"created_utc"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}
