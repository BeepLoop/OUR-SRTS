package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com.BeepLoop/registrar-digitized/store"
	"github.com.BeepLoop/registrar-digitized/types"
	"github.com.BeepLoop/registrar-digitized/utils"
)

func HandleManageStaff(c *gin.Context) {
	html := utils.HtmlParser(
		"admin/manage-staff.html",
		"components/head.html",
		"components/header.html",
		"components/sidebar.html",
		"components/add-staff.html",
	)

	user := utils.GetUserInSession(c)

	users, err := store.GetUsers()
	if err != nil {
		fmt.Println("err getting users: ", err)

		html.Execute(c.Writer, gin.H{
			"user":  user.Fullname,
			"users": []types.User{},
		})
		return
	}

	html.Execute(c.Writer, gin.H{
		"user":  user,
		"users": users,
	})
}
