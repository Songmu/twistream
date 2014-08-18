package twistream

import (
	"bufio"
	"encoding/json"
	"log"
)

// Stream type
type Stream struct {
	scanner *bufio.Scanner
}

// NextTweet returns new tweet
func (s *Stream) NextTweet() (tweet *Status, err error) {
	for s.scanner.Err() == nil {
		var bytes []byte
		bytes, err = func() ([]byte, error) {
			for {
				if !s.scanner.Scan() {
					return nil, s.scanner.Err()
				}
				bytes := s.scanner.Bytes()
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
		}
	}
	return nil, s.scanner.Err()
}
