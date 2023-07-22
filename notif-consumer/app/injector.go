//go:build wireinject
// +build wireinject

package app

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/google/wire"
	"github.com/mnhjddn/go-notif-consumer/infra"
	"gorm.io/gorm"
)

var (
	AppModule = wire.NewSet(
		configModule,
		ProvideConfig,
		ProvideMySQL,
		ProvideSubscriber,
	)
)

func InjectConfig() infra.AppConfig {
	panic(wire.Build(AppModule))
}

func InjectMySQL() *gorm.DB {
	panic(wire.Build(AppModule))
}

func InjectSubscriber() *amqp.Subscriber {
	panic(wire.Build(AppModule))
}
