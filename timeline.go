package twistream

import (
	"bufio"
	"log"
)

type Timeline struct {
	client   *api
	endpoint string
	params   map[string]string
	stream   *Stream
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

// Listen bytes sent from Twitter Streaming API
// and send completed status to the channel.
func (tl *Timeline) Listen() (chan *Status, error) {
	response, e := tl.client.Get(
		tl.endpoint,
		tl.params,
	)
	if e != nil {
		return nil, e
	}

	tl.stream = &Stream{
		scanner: bufio.NewScanner(response.Body),
	}
	tweets_chan := make(chan *Status)

	go func() {
		for {
			tweet, err := tl.stream.NextTweet()
			if err != nil {
				log.Fatal(err)
			}
			tweets_chan <- tweet
		}
	}()

	return tweets_chan, nil
}

// Tweet posts status to the timeline
func (tl *Timeline) Tweet(status Status) (e error) {
	_, e = tl.client.Post(
		"https://api.twitter.com/1.1/statuses/update.json",
		status.ToParams(),
	)
	return
}
