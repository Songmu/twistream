package twistream_test

// test
import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/Songmu/twistream"
	. "github.com/otiai10/mint"
)

var (
	CONSUMER_KEY        string
	CONSUMER_SECRET     string
	ACCESS_TOKEN        string
	ACCESS_TOKEN_SECRET string
)

func TestNew(t *testing.T) {
	timeline := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		CONSUMER_KEY,
		CONSUMER_SECRET,
		ACCESS_TOKEN,
		ACCESS_TOKEN_SECRET,
		map[string]string{},
	)
	Expect(t, timeline).TypeOf("*twistream.Timeline")

	ch, e := timeline.Listen()
	Expect(t, e).ToBe(nil)

	var terminate = make(chan bool)
	go func() {
		for {
			status := <-ch
			fmt.Printf("%+v\n", status)
			terminate <- true
			break
		}
	}()

	if <-terminate {
		return
	}
}

func TestTimeline_Tweet(t *testing.T) {
	timeline := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		CONSUMER_KEY,
		CONSUMER_SECRET,
		ACCESS_TOKEN,
		ACCESS_TOKEN_SECRET,
		map[string]string{},
	)

	status := twistream.Status{
		Text:              "This is test!!" + time.Now().String(),
		InReplyToStatusId: 493324823926288386,
	}
	e := timeline.Tweet(status)
	Expect(t, e).ToBe(nil)
}

func init() {
	buffer, _ := ioutil.ReadFile("test.conf")
	lines := strings.Split(string(buffer), "\n")
	if len(lines) < 4 {
		panic("test.conf requires at least 4 lines")
	}
	CONSUMER_KEY = lines[0]
	CONSUMER_SECRET = lines[1]
	ACCESS_TOKEN = lines[2]
	ACCESS_TOKEN_SECRET = lines[3]
}
