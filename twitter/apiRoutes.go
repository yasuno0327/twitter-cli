package twitter

type ApiRoute struct {
	Method string
	Path string
}

type ApiRoutes = map[string]ApiRoute

const ResourceURLPrefix = "https://api.twitter.com/1.1/"

var Routes = ApiRoutes {
	// Tweets
	"statuses/update": ApiRoute{Method: "POST", Path: "statuses/update"},
	"statuses/destroy": ApiRoute{Method: "POST", Path: "statuses/destroy/:id"},
	"statuses/show": ApiRoute{Method: "GET", Path: "statuses/show/:id"},
	"statuses/oembed": ApiRoute{Method: "GET", Path: "statuses/oembed"},
	"statuses/lookup": ApiRoute{Method: "GET", Path: "statuses/lookup"},
	"statuses/home_timeline": ApiRoute{Method: "GET", Path: "statuses/home_timeline"},
	"statuses/user_timeline": ApiRoute{Method: "GET", Path: "statuses/user_timeline"},
	"statuses/mentions_timeline": ApiRoute{Method: "GET", Path: "statuses/mentions_timeline"},
}