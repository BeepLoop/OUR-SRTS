package server

import (
	"encoding/gob"
	"io"
	"net/http"
	"os"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/BeepLoop/registrar-digitized/middleware"
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Router *gin.Engine

func NewServer() {
	gob.Register(types.User{})

	logLevel, err := logrus.ParseLevel(config.Env.LogLevel)
	if err != nil {
		logrus.Fatal("Error parsing logrus level in the env: ", err)
	}
	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	logFile, err := os.OpenFile("./server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		logrus.Fatal("Error creating log file: ", err)
	}

	if config.Env.GinMode == "release" {
		logrus.SetOutput(logFile)
		gin.SetMode(gin.ReleaseMode)
	} else {
		multiwriter := io.MultiWriter(logFile, os.Stdout)
		logrus.SetOutput(multiwriter)
	}

	Router = gin.New()

	sessionStore := InitSession()

	Router.Use(gin.Recovery())
	Router.Use(middleware.LogrusMiddleware)
	Router.Use(middleware.MimeType)
	Router.Use(sessions.Sessions("user", sessionStore))

	Router.NoRoute(HandleNotFoundRote)

	Router.Static("/styles", "views/styles/")
	Router.Static("/scripts", "views/scripts/")
	Router.Static("/public", "assets/public/")
	Router.Static("/fonts", "webfonts/")
	Router.StaticFS("/nas", http.Dir(config.Env.NasUrl))

	RegisterRoutes()

}
