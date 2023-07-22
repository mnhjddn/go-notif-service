package notif

type SendNotifRequest struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type SendNotifResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
