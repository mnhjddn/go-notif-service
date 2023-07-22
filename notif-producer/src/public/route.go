package public

import "github.com/gin-gonic/gin"

type Routes struct {
	delivery *Delivery
}

func NewRoutes(delivery *Delivery) *Routes {
	return &Routes{delivery: delivery}
}

func (routes Routes) RegisterRoutes(r *gin.Engine) {

	route := r.Group("/public")
	{
		route.GET("/ping", routes.delivery.Ping)
	}
}
