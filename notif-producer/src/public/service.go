package public

import "github.com/mnhjddn/go-notif-producer/infra"

type Service struct {
	config infra.AppConfig
	repo   IRepository
}

func NewService(config infra.AppConfig, repo IRepository) *Service {
	return &Service{config: config, repo: repo}
}

func (svc Service) Ping() (data interface{}, err error) {
	mapValue := map[string]interface{}{
		"connection": "ok",
	}

	if svc.config.App.Env == "local" {
		mapValue["config"] = svc.repo.GetConfig()
	}

	data = mapValue

	return
}
