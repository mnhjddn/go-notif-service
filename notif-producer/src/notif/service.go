package notif

import (
	"encoding/json"
)

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{repo}
}

func (s *Service) SendNotif(req SendNotifRequest) (err error) {

	payloadStr, err := json.Marshal(req.Data)
	if err != nil {
		err = ErrBadRequest
		return
	}

	notifData := NotificationData{
		Type:    req.Type,
		Payload: string(payloadStr),
	}

	notifId, err := s.repo.SaveNotification(notifData)
	if err != nil {
		return
	}

	go s.repo.PublishNotification(NotifQueueData{
		Id: notifId,
	})

	return
}
