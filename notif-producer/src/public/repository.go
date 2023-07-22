package public

import "github.com/mnhjddn/go-notif-producer/infra"

type Repository struct {
	config infra.AppConfig
}

func (repo Repository) GetConfig() infra.AppConfig {
	return repo.config
}

func NewRepository(config infra.AppConfig) *Repository {
	return &Repository{config: config}
}
