package twistream

import (
	"bufio"
	"log"
	"net/http"
)

type Timeline struct {
	response    *http.Response
	tweets_chan chan *Status
	client      *api
	endpoint    string
	params      map[string]string
	stream      *Stream
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
	tl.stream = &Stream{
		scanner: bufio.NewScanner(response.Body),
	}

	return e
}

// Listen bytes sent from Twitter Streaming API
// and send completed status to the channel.
func (tl *Timeline) Listen() <-chan *Status {
	// Delegate channel to parser.

	tl.tweets_chan = make(chan *Status)

	go func() {
		for {
			tweet, err := tl.stream.NextTweet()
			if err != nil {
				log.Fatal(err)
			}
			tl.tweets_chan <- tweet
		}
	}()

	return tl.tweets_chan
}

// Tweet posts status to the timeline
func (tl *Timeline) Tweet(status Status) (e error) {
	_, e = tl.client.Post(
		"https://api.twitter.com/1.1/statuses/update.json",
		status.ToParams(),
	)
	return
}
