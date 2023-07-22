//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/mnhjddn/go-notif-producer/infra"
)

var (
	configModule = wire.NewSet(ProvideAppConfig)
)

func ProvideAppConfig() *infra.AppConfig {
	return infra.NewConfig()
}

func InitializeAppConfig() (*infra.AppConfig, error) {
	panic(wire.Build(configModule))
}
