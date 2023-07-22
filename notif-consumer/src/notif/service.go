package notif

import (
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/mnhjddn/go-notif-consumer/infra"
)

type Service struct {
	repo   IRepository
	mailer infra.IMailer
}

func NewService(repo IRepository, mailer infra.IMailer) *Service {
	return &Service{repo, mailer}
}

func (s *Service) ProcessNotif(payload message.Payload) {
	var notifQueue NotifQueueData
	err := json.Unmarshal(payload, &notifQueue)
	if err != nil {
		fmt.Println(err)
		return
	}

	notifData, err := s.repo.GetNotification(notifQueue.Id)
	if err != nil {
		fmt.Println(err)
		return
	}

	go s.processNotifFactory(notifData)
}

func (s *Service) processNotifFactory(notifData NotificationData) {
	switch notifData.Type {
	case NotificationTypeEmail:
		var emailData EmailData
		err := json.Unmarshal([]byte(notifData.Payload), &emailData)
		if err != nil {
			fmt.Println(err)
			return
		}

		s.mailer.SendEmail([]string{emailData.Recipient}, emailData.Subject, emailData.Body)
	}
}
