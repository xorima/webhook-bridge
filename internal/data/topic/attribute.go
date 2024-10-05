package topic

// Attribute is optional metadata which should be included with the
// event, this can be used on some systems to have filtering available
// for example
type Attribute struct {
	Key   string
	Value any
}

func NewAttribute(key string, value any) Attribute {
	return Attribute{
		Key:   key,
		Value: value,
	}
}
