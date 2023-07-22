package notif

type NotificationData struct {
	Id      int    `gorm:"id;primary_key;AUTO_INCREMENT"`
	Type    string `gorm:"type"`
	Payload string `gorm:"payload"`
}

func (*NotificationData) TableName() string {
	return "notification"
}

type NotifQueueData struct {
	Id int `json:"id"`
}
