package app

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/mnhjddn/go-notif-consumer/infra"
	"gorm.io/gorm"
)

func ProvideConfig(conf *infra.AppConfig) infra.AppConfig {
	return *conf
}

func ProvideMySQL(conf *infra.AppConfig) *gorm.DB {
	return infra.NewMySQL(conf)
}

func ProvideSubscriber(conf *infra.AppConfig) *amqp.Subscriber {
	return infra.NewSubscriber(conf)
}
