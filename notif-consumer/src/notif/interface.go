package notif

import "github.com/ThreeDotsLabs/watermill/message"

type IRepository interface {
	GetNotification(id int) (data NotificationData, err error)
}

type IService interface {
	ProcessNotif(payload message.Payload)
}
