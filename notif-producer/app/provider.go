package app

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/mnhjddn/go-notif-producer/infra"
	"gorm.io/gorm"
)

func ProvideConfig(conf *infra.AppConfig) infra.AppConfig {
	return *conf
}

func ProvideMySQL(conf *infra.AppConfig) *gorm.DB {
	return infra.NewMySQL(conf)
}

func ProvidePublisher(conf *infra.AppConfig) *amqp.Publisher {
	return infra.NewPublisher(conf)
}
