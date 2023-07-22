package notif

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/mnhjddn/go-notif-producer/infra"
	"gorm.io/gorm"
)

type Repository struct {
	db        *gorm.DB
	publisher *amqp.Publisher
	config    *infra.AppConfig
}

func NewRepository(db *gorm.DB, publisher *amqp.Publisher, config *infra.AppConfig) *Repository {
	return &Repository{
		db:        db,
		publisher: publisher,
		config:    config,
	}
}

func (r *Repository) SaveNotification(data NotificationData) (id int, err error) {
	err = r.db.Create(&data).Error
	if err != nil {
		return
	}

	id = data.Id
	return
}

func (r *Repository) PublishNotification(data NotifQueueData) {
	payload, err := json.Marshal(data)
	if err != nil {
		return
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)
	err = r.publisher.Publish(r.config.RabbitMQ.TopicNotif, msg)
	if err != nil {
		return
	}
}
