package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/registrar/config"
	"github.com/registrar/store"
)

func FileSaver(c *gin.Context, file *multipart.FileHeader, lastname, controlNumber, filetype string) (string, error) {
	ext := filepath.Ext(file.Filename)

	if StringInSlice(filetype) {
		count, err := store.CountFilesInMultiEntry(filetype, controlNumber)
		if err != nil {
			fmt.Println("err counting: ", err)
			return "", err
		}

		strCount := strconv.Itoa(count)
		filename := lastname + "_" + controlNumber + "_" + filetype + "_" + strCount + ext
		path := config.BASE_FOLDER + filename

		err = c.SaveUploadedFile(file, path)
		if err != nil {
			return path, err
		}

		return path, nil
	}

	filename := lastname + "_" + controlNumber + "_" + filetype + ext
	path := config.BASE_FOLDER + filename

	err := c.SaveUploadedFile(file, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
