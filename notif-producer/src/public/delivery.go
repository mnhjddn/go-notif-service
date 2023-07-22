package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Delivery struct {
	service IService
}

func NewDelivery(service IService) *Delivery {
	return &Delivery{service: service}
}

func (dlv Delivery) Ping(ctx *gin.Context) {
	data, err := dlv.service.Ping()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Next()
	ctx.Header("Content-Security-Policy", "default-src 'self'")
	ctx.JSON(http.StatusOK, data)
}
