package notif

import "github.com/mnhjddn/go-notif-consumer/infra"

type Broker struct {
	consumer *Consumer
	config   *infra.AppConfig
}

func (b Broker) RegisterBrokerNotification() {
	b.consumer.ConsumeDataNotification(b.config.RabbitMQ.TopicNotif)
}

func NewBroker(consumer *Consumer, config *infra.AppConfig) *Broker {
	return &Broker{consumer: consumer, config: config}
}
