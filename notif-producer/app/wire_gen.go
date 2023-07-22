// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/google/wire"
	"github.com/mnhjddn/go-notif-producer/infra"
	"gorm.io/gorm"
)

// Injectors from config_wire.go:

func InitializeAppConfig() (*infra.AppConfig, error) {
	appConfig := ProvideAppConfig()
	return appConfig, nil
}

// Injectors from injector.go:

func InjectConfig() infra.AppConfig {
	appConfig := ProvideAppConfig()
	infraAppConfig := ProvideConfig(appConfig)
	return infraAppConfig
}

func InjectMySQL() *gorm.DB {
	appConfig := ProvideAppConfig()
	db := ProvideMySQL(appConfig)
	return db
}

func InjectPublisher() *amqp.Publisher {
	appConfig := ProvideAppConfig()
	publisher := ProvidePublisher(appConfig)
	return publisher
}

// config_wire.go:

var (
	configModule = wire.NewSet(ProvideAppConfig)
)

func ProvideAppConfig() *infra.AppConfig {
	return infra.NewConfig()
}

// injector.go:

var (
	AppModule = wire.NewSet(
		configModule,
		ProvideConfig,
		ProvideMySQL,
		ProvidePublisher,
	)
)