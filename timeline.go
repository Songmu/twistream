package twistream

import "net/http"
import "regexp"

type Timeline struct {
	response *http.Response
	stream   chan Status
	client   *api
	endpoint string
	params   map[string]string
}

// New provides new reference for specified Timeline.
func New(endpoint, consumerKey, consumerSecret, accessToken, accessTokenSecret string, params map[string]string) *Timeline {

	return &Timeline{
		client: initAPI(
			consumerKey,
			consumerSecret,
			accessToken,
			accessTokenSecret,
		),
		endpoint: endpoint,
		params:   params,
	}
}

func (tl *Timeline) Init() (e error) {
	response, e := tl.client.Get(
		tl.endpoint,
		tl.params,
	)
	tl.response = response
	tl.stream = make(chan Status)
	return e
}

// Listen bytes sent from Twitter Streaming API
// and send completed status to the channel.
func (tl *Timeline) Listen() <-chan Status {
	// Delegate channel to parser.
	p := &parser{
		proxy:   tl.stream,
		trigger: regexp.MustCompile("^[0-9a-z]+\r\n$"),
	}
	go func() {
		for {
			tl.response.Write(p)
		}
	}()
	return tl.stream
}

// Tweet posts status to the timeline
func (tl *Timeline) Tweet(status Status) (e error) {
	_, e = tl.client.Post(
		"https://api.twitter.com/1.1/statuses/update.json",
		status.ToParams(),
	)
	return
}
