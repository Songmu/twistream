package twistream_test

import "github.com/Songmu/twistream"
import "testing"
import "fmt"

func ExampleTimeline_Listen(t *testing.T) {
	timeline := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		"CONSUMER_KEY",
		"CONSUMER_SECRET",
		"ACCESS_TOKEN",
		"ACCESS_TOKEN_SECRET",
		map[string]string{},
	)
	// Practically, you would run this loop by goroutine

	ch, err := timeline.Listen()

	if err != nil {
		panic(err)
	}

	for {
		status := <-ch
		fmt.Println(status)
	}
}
