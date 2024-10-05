package topic

import "context"

type EventManager interface {
	EventProducer
}

// EventProducer adds an Event to the Channel
type EventProducer interface {
	Produce(ctx context.Context, channel Channel, event Event) error
}

type EventConsumer interface {
	Consume(ctx context.Context, channel Channel, consumer Consumer) (Event, error)
}
