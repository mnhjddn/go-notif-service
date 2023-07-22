package shared

type FailedResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (f FailedResponse) Error() string {
	return f.Message
}

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
