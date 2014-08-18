package twistream

// see. https://dev.twitter.com/docs/streaming-apis/messages#User_stream_messages
type Event struct {
	Event string `json:"event"`
}
