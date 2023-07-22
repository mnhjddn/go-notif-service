//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/mnhjddn/go-notif-producer/infra"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"gorm.io/gorm"
)

var (
	AppModule = wire.NewSet(
		configModule,
		ProvideConfig,
		ProvideMySQL,
		ProvidePublisher,
	)
)

func InjectConfig() infra.AppConfig {
	panic(wire.Build(AppModule))
}

func InjectMySQL() *gorm.DB {
	panic(wire.Build(AppModule))
}

func InjectPublisher() *amqp.Publisher {
	panic(wire.Build(AppModule))
}