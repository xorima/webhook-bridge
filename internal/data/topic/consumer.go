package topic

type Consumer struct {
	Name            string
	Group           string
	MessageStrategy ConsumerStrategy
}

func NewConsumer(name, group string) *Consumer {
	return &Consumer{Name: name, Group: group, MessageStrategy: ConsumerStrategyUndelivered}
}
func (c *Consumer) WithConsumeNewFromAllTime() *Consumer {
	c.MessageStrategy = ConsumerStrategyAllTime
	return c
}

type ConsumerStrategy uint

const (
	ConsumerStrategyAllTime ConsumerStrategy = iota
	ConsumerStrategyUndelivered
)
