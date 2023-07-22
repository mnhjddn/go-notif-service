//go:build wireinject
// +build wireinject

package notif

import (
	"github.com/google/wire"
	"github.com/mnhjddn/go-notif-producer/app"
)

var (
	ModuleSets = wire.NewSet(
		NewRepository,
		NewService,
		wire.Bind(new(IRepository), new(*Repository)),
		wire.Bind(new(IService), new(*Service)),
		NewDelivery,
		NewRoutes)
)

func InjectRoutes() (*Routes, error) {
	panic(wire.Build(app.AppModule, ModuleSets))
}
