package notif

import "github.com/mnhjddn/go-notif-consumer/shared"

const (
	NotificationTypeEmail = "email"

	// queue
	MaxConcurrencyJobs = 10
)

var (
	// errors
	ErrBadRequest = shared.FailedResponse{Status: 400, Message: "Bad Request"}
)
