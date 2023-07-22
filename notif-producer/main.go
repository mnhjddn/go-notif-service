package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mnhjddn/go-notif-producer/app"
	"github.com/mnhjddn/go-notif-producer/src/notif"
	"github.com/mnhjddn/go-notif-producer/src/public"
	log "github.com/sirupsen/logrus"
)

func main() {

	config, err := app.InitializeAppConfig()
	if err != nil {
		panic(err)
	}

	// Init gin
	gin.ForceConsoleColor()
	router := gin.New()
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// public routes
	publicRoutes, err := public.InjectRoutes()
	if err != nil {
		panic(err.Error())
	}
	publicRoutes.RegisterRoutes(router)

	// notif routes
	notifRoutes, err := notif.InjectRoutes()
	if err != nil {
		panic(err.Error())
	}
	notifRoutes.RegisterRoutes(router)

	log.Fatal(router.Run(":" + config.App.Port))

	app.Listen(*config, router)
}
