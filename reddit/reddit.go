package reddit

import (
	"../cache"
	"../utils"
)

type Client struct {
	cache   cache.Cache // Client should hold its own cache
	apiRoot string      // Keep track of the API root for easy versioning updates
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

	return Client{
		cache:   args.cache,
		apiRoot: args.apiRoot,
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

func (client *Client) GetPosts() *Response {

	return &Response{}
}
