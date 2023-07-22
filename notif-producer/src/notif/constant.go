package notif

import "github.com/mnhjddn/go-notif-producer/shared"

var (
	// errors
	ErrBadRequest = shared.FailedResponse{Status: 400, Message: "Bad Request"}
)
