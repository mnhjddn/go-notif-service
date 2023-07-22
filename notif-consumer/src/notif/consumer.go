package notif

import (
	"context"

	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Consumer struct {
	service    IService
	subscriber *amqp.Subscriber
}

func NewConsumer(service IService, subscriber *amqp.Subscriber) *Consumer {
	return &Consumer{
		service:    service,
		subscriber: subscriber,
	}
}

func (c *Consumer) ConsumeDataNotification(topicName string) {
	messages, _ := c.subscriber.Subscribe(context.Background(), topicName)

	go c.process(messages)
}

func (c *Consumer) process(messages <-chan *message.Message) {
	waitChan := make(chan struct{}, MaxConcurrencyJobs)
	for msg := range messages {
		waitChan <- struct{}{}
		go func(msg *message.Message) {
			msg.Ack()
			c.service.ProcessNotif(msg.Payload)
		}(msg)
	}
}
