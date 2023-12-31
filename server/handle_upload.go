package server

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleUpload(c *gin.Context) {
	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	form, err := c.MultipartForm()
	if err != nil {
		logrus.Warn("err binding form: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=invalid_form")
		return
	}

	lastname := form.Value["lastname"][0]
	controlNumber := form.Value["controlNumber"][0]

	// Use loop for getting file even though it will always be one file
	// because I'm too lazy to check what file is being uploaded
	for key := range form.File {
		file := form.File[key][0]

		if key != "Photo" {
			// Only allow pdf files
			if utils.IsFilePdf(file) == false {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=not_pdf")
				return
			}
		}

		// Separate method for handling Other files
		if key == "Other" {
			filename := form.Value["filename"][0]
			file, location, err := utils.SaveOtherFile(filename, lastname, controlNumber, key, file, c)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=invalid_form")
				return
			}

			// replace backslash with forward slash for windows
			if runtime.GOOS == "windows" {
				location = strings.ReplaceAll(location, "\\", "/")
			}

			err = store.SaveOtherFile(file, location, controlNumber, key)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=invalid_form")
				return
			}

		} else {
			location, _, err := utils.FileSaver(c, file, lastname, controlNumber, key)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=invalid_form")
				return
			}

			// replace backslash with forward slash for windows
			if runtime.GOOS == "windows" {
				location = strings.ReplaceAll(location, "\\", "/")
			}

			err = store.SaveFile(location, controlNumber, key)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=invalid_form")
				return
			}
		}
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, url+"?status=success")
	return
}
