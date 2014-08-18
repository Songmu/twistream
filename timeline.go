package twistream

import "log"

type Timeline struct {
	client    *api
	endpoint  string
	params    map[string]string
	stream    *Stream
	Reconnect bool
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
		endpoint:  endpoint,
		params:    params,
		Reconnect: false,
	}
}

func (tl *Timeline) Connect() error {
	response, e := tl.client.Get(
		tl.endpoint,
		tl.params,
	)
	tl.stream = newStream(response)
	return e
}

// TODO: revise reconnect strategy
func (tl *Timeline) reconnect() error {
	return tl.Connect()
}

// Listen bytes sent from Twitter Streaming API
// and send completed status to the channel.
func (tl *Timeline) Listen() (chan *Status, error) {
	if e := tl.Connect(); e != nil {
		return nil, e
	}

	tweets_chan := make(chan *Status)

	go func() {
		for {
			tweet, err := tl.stream.NextTweet()
			if err != nil {
				if tl.Reconnect && tl.stream.response.Close {
					log.Println("connection closed. try reconnect")
					err = tl.reconnect()
				}
				if err != nil {
					log.Fatal(err)
				}
			} else {
				tweets_chan <- tweet
			}
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
