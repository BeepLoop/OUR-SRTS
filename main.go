package main

import (
	"net/http"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/BeepLoop/registrar-digitized/server"
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/sirupsen/logrus"
)

func init() {
	err := config.Initialize()
	if err != nil {
		logrus.Fatal(err)
	}

	err = store.Initialize()
	if err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	server.NewServer()

	logrus.Infof("Server listening on %s%s", config.Env.LocalAddr, config.Env.Port)
	err := http.ListenAndServe(config.Env.Port, server.Router)
	if err != nil {
		logrus.Fatal(err)
	}
}
