package topic

// Event defines the entire item to be put on the Channel
// every event should be versioned correctly following a semver
// policy to ensure consumers can rapidly understand what is
// available in the Event
type Event struct {
	Version    string
	Body       string
	Attributes []Attribute
}

func NewEvent(version, body string, attributes ...Attribute) *Event {
	return &Event{
		Version:    version,
		Body:       body,
		Attributes: attributes,
	}
}
