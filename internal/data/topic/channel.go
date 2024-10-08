package topic

type Channel struct {
	Name   string
	Prefix []string
}

func NewChannel(name string) *Channel {
	return &Channel{Name: name}
}

// WithPrefix defines any prefix that should be used
// this will be used by the event sender to set the correct
// named channel based on the end system's naming rules.
func (c *Channel) WithPrefix(prefix ...string) *Channel {
	c.Prefix = append(c.Prefix, prefix...)
	return c
}
