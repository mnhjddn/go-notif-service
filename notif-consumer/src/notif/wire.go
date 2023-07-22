//go:build wireinject
// +build wireinject

package notif

import (
	"github.com/google/wire"
	"github.com/mnhjddn/go-notif-consumer/app"
	"github.com/mnhjddn/go-notif-consumer/infra"
)

var (
	ModuleSets = wire.NewSet(
		infra.NewMailer,
		NewRepository,
		NewService,
		wire.Bind(new(infra.IMailer), new(*infra.Mailer)),
		wire.Bind(new(IRepository), new(*Repository)),
		wire.Bind(new(IService), new(*Service)),
		NewConsumer,
		NewBroker)
)

func InjectBroker() (*Broker, error) {
	panic(wire.Build(app.AppModule, ModuleSets))
}
