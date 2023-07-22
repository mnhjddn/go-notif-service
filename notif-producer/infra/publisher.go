package infra

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	log "github.com/sirupsen/logrus"
)

const (
	ErrInitializingAMQP = "Error initializing AMQP"
)

func NewPublisher(conf *AppConfig) *amqp.Publisher {
	amqpConfig := amqp.NewDurableQueueConfig(fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.RabbitMQ.Usr, conf.RabbitMQ.Pw, conf.RabbitMQ.Host, conf.RabbitMQ.Port))
	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		log.Panic(ErrInitializingAMQP)
	}

	return publisher
}
