package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mnhjddn/go-notif-consumer/app"
	"github.com/mnhjddn/go-notif-consumer/src/notif"
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

	// notif broker
	notifBroker, err := notif.InjectBroker()
	if err != nil {
		panic(err)
	}
	notifBroker.RegisterBrokerNotification()

	log.Fatal(router.Run(":" + config.App.Port))

	app.Listen(*config, router)
}
