package main

import (
	"net/http"

	"github.com/hoochycoochie/webgolangstructure/internal/api"
	"github.com/hoochycoochie/webgolangstructure/internal/config"
	"github.com/sirupsen/logrus"
)

//Create server object and start listnener
func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("version", config.Version).Debug("Starting server")

	//Creating router
	router, err := api.NewRouter()

	if err != nil {
		logrus.WithError(err).Debug("Error building router")
	}
	const addr = "0.0.0.0:8088"
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}

	//Starting server
	if err = server.ListenAndServe(); err != nil {
		logrus.WithError(err).Debug("Error server starting ")
	} else {
		logrus.Info("Serving started")
	}
}
