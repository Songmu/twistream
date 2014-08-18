# Twitter Streaming API

The very simplest interface to use Twitter Streaming API by golang.
This is forked from <https://github.com/otiai10/twistream>.

# Usage

```go
timeline := twistream.New(
    "https://userstream.twitter.com/1.1/user.json",
    CONSUMERKEY,
    CONSUMERSECRET,
    ACCESSTOKEN,
    ACCESSTOKENSECRET,
    map[string]string{
        with: "followers",
    },
)

// Listen timeline
ch, _ := timeline.Listen()
for {
    status := <-ch
    fmt.Println(status)
}

// Tweet to timeline
status := twistream.Status{
    Text: "@otiai10 How does Go work?",
    InReplyToStatusId: 493324823926288386,
}
_ := timeline.Tweet(status)
```

# TODOs

- [x] GET user
- [ ] GET site
- [ ] GET statuses/sample
- [ ] GET status/firehose
- [ ] POST statuses/filter
- [x] POST statuses/update
- [ ] POST statuses/update_with_media
