package public

import "github.com/mnhjddn/go-notif-producer/infra"

type IRepository interface {
	GetConfig() infra.AppConfig
}

type IService interface {
	Ping() (data interface{}, err error)
}
