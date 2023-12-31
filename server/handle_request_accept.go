package server

import (
	"net/http"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func HandleRequestFulfill(c *gin.Context) {

	user := utils.GetUserInSession(c)
	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	type Accept struct {
		RequestId   string `form:"requestId" binding:"required"`
		Password    string `form:"password" binding:"required"`
		NewPassword string `form:"newPassword" binding:"required"`
	}

	var accept Accept
	err := c.ShouldBindWith(&accept, binding.Form)
	if err != nil {
		logrus.Warn("err binding form: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=invalid_request")
		return
	}

	credential, err := store.GetCredentials(user.Username)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=cant_get_credentials")
		return
	}

	err = utils.ValidateCredentials(accept.Password, credential.Password)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=Invalid_admin_password")
		return
	}

	if !utils.ValidPassword(accept.NewPassword) {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=Password_does_not_meet_requirements")
		return
	}

	hash, err := utils.HashPassword(accept.NewPassword)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=hashing_error")
		return
	}

	err = store.FulfillRequest(accept.RequestId, hash)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=database_error")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, url+"?status=success")
}
