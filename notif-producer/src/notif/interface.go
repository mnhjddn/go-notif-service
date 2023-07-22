package notif

type IRepository interface {
	SaveNotification(data NotificationData) (id int, err error)
	PublishNotification(data NotifQueueData)
}

type IService interface {
	SendNotif(req SendNotifRequest) (err error)
}
