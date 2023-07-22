package notif

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mnhjddn/go-notif-producer/shared"
)

type Delivery struct {
	service IService
}

func NewDelivery(service IService) *Delivery {
	return &Delivery{service: service}
}

func (dlv Delivery) SendNotif(ctx *gin.Context) {
	var req SendNotifRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(ErrBadRequest.Status, ErrBadRequest)
		return
	}

	err = dlv.service.SendNotif(req)
	if err != nil {
		switch err.Error() {
		case ErrBadRequest.Error():
			ctx.JSON(ErrBadRequest.Status, ErrBadRequest)
			return
		default:
			resp := shared.FailedResponse{Status: http.StatusInternalServerError, Message: err.Error()}
			ctx.JSON(resp.Status, resp)
			return
		}
	}

	resp := shared.SuccessResponse{Status: http.StatusOK, Message: http.StatusText(http.StatusOK)}
	ctx.JSON(resp.Status, resp)
}
