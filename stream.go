package twistream

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
)

// Stream type
type Stream struct {
	response *http.Response
	scanner  *bufio.Scanner
	friends  []int64
}

type friends struct {
	Friends []int64 `json:"friends"`
}

func newStream(response *http.Response) (stream *Stream) {
	stream = &Stream{
		response: response,
	}
	stream.scanner = bufio.NewScanner(response.Body)
	return stream
}

// NextTweet returns new tweet
func (s *Stream) NextTweet() (tweet *Status, err error) {
	scanner := s.scanner

	for scanner.Err() == nil {
		var bytes []byte
		bytes, err = func() ([]byte, error) {
			for {
				if !scanner.Scan() {
					return nil, scanner.Err()
				}
				bytes := scanner.Bytes()
				if len(bytes) > 0 {
					return bytes, nil
				}
			}
		}()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bytes, &tweet)

		if err != nil {
			event := new(Event)
			err = json.Unmarshal(bytes, &event)

			if err != nil {
				log.Println(string(bytes))
				return nil, err
			} else {
				log.Println(event.Event + " event accepted!")
			}
		}
		if tweet.Id > 0 {
			return tweet, nil
		} else {
			friends := new(friends)
			_ = json.Unmarshal(bytes, &friends)
			s.friends = friends.Friends
		}
	}
	return nil, scanner.Err()
}
