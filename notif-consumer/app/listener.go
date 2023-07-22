package app

import (
	"net/http"

	"github.com/mnhjddn/go-notif-consumer/infra"
	log "github.com/sirupsen/logrus"
)

// Listen is a func to start http server
func Listen(conf infra.AppConfig, handler http.Handler) {
	addr := conf.App.Host + ":" + conf.App.Port
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatal(err)
	}
}
