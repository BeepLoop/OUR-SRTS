package server

import (
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleGetLogin(c *gin.Context) {
	logrus.Info("Hit Get login route")
	html := utils.HtmlParser("login.tmpl", "components/head.tmpl", "components/header.tmpl")

	html.Execute(c.Writer, nil)
}
