package reddit

import (
	"encoding/json"
	"fmt"
	"github.com/cameronbroe/daily_programmer/cache"
	"github.com/cameronbroe/daily_programmer/utils"
	"io/ioutil"
	"net/http"
)

const UserAgent = "go:com.cameronbroe.daily_programmer:v0.0.1 (by /u/CommanderViral)"

type Client struct {
	cache      cache.Cache // Client should hold its own cache
	apiRoot    string      // Keep track of the API root for easy versioning updates
	httpClient *http.Client
}

type Post struct{}

type Response struct {
	client Client // Keep hold of the parent client to support response pagination
	count  int
	after  string
	data   []Post
}

type ClientOptions struct {
	cache   cache.Cache // Use an already created cache
	apiRoot string      // API root if wanting to override (even though this probably is useless)
}

type ClientOption func(*ClientOptions)

func clientCache(cache cache.Cache) ClientOption {
	return func(s *ClientOptions) { s.cache = cache }
}

func apiRoot(apiRoot string) ClientOption {
	return func(s *ClientOptions) { s.apiRoot = apiRoot }
}

func New(options ...ClientOption) Client {
	var defaultCache cache.Cache
	if utils.DefaultCacheExists() {
		defaultCache = cache.FromFile() // Load from default path ~/.daily_programmer_cache
	} else {
		defaultCache = cache.New() // Create a new cache
	}

	args := &ClientOptions{
		cache:   defaultCache,
		apiRoot: "https://reddit.com/",
	}

	for _, option := range options {
		option(args)
	}

	// Create HTTP client
	httpClient := &http.Client{}

	return Client{
		cache:      args.cache,
		apiRoot:    args.apiRoot,
		httpClient: httpClient,
	}
}

type RequestOptions struct {
	count int
	after string
}

type RequestOption func(*RequestOptions)

func count(c int) RequestOption {
	return func(s *RequestOptions) { s.count = c }
}

func after(a string) RequestOption {
	return func(s *RequestOptions) { s.after = a }
}

func (client *Client) GetPosts() Response {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/r/dailyprogrammer.json", client.apiRoot), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", UserAgent)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	data := ResponseData{}
	json.Unmarshal(body, &data)
	fmt.Println(data)
	return Response{}
}
