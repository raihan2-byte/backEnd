package handler

import (
	endpointcount "blog/endpointCount"
	"blog/helper"
	"blog/imagekits"
	"blog/shortvideo"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type shortVideoHandler struct {
	shortVideoService shortvideo.Service
	endpointService   endpointcount.StatisticsService
}

func NewShortVideoHandler(shortVideoService shortvideo.Service, endpointService endpointcount.StatisticsService) *shortVideoHandler {
	return &shortVideoHandler{shortVideoService, endpointService}
}

func (h *shortVideoHandler) CreateShortVideo(c *gin.Context) {
	file, _ := c.FormFile("file")
	src, err := file.Open()
	if err != nil {
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to open file")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	defer src.Close()
	if err != nil {
		fmt.Printf("error when open file %v", err)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v", err)
		return
	}

	img, err := imagekits.Base64toEncode(buf.Bytes())
	if err != nil {
		fmt.Println("error reading image %v", err)
	}

	fmt.Println("image base 64 format : %v", img)

	imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var input shortvideo.InputShortVideo

	err = c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		//inisiasi data yang tujuan dalam return hasil ke postman
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.shortVideoService.CreateShortVideo(input, imageKitURL)
	if err != nil {
		// data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

func (h *shortVideoHandler) GetAllShortVideo(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	newBerita, err := h.shortVideoService.GetAllShortVideo(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, "Eror")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// userAgent := c.GetHeader("User-Agent")

	// err = h.endpointService.IncrementCount("View-All-Short-Video")
	// if err != nil {
	//     response := helper.APIresponse(http.StatusUnprocessableEntity, err)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	response := helper.APIresponse(http.StatusOK, newBerita)
	c.JSON(http.StatusOK, response)

}

func (h *shortVideoHandler) GetOneShortVideo(c *gin.Context) {
	var input shortvideo.GetShortVideoID

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.shortVideoService.GetOneShortVideo(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	// userAgent := c.GetHeader("User-Agent")

	// err = h.endpointService.IncrementCount("View-Short-Video")
	// if err != nil {
	//     response := helper.APIresponse(http.StatusUnprocessableEntity, err)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	response := helper.APIresponse(http.StatusOK, newDel)
	c.JSON(http.StatusOK, response)
}

func (h *shortVideoHandler) DeleteShortVideo(c *gin.Context) {
	var input shortvideo.GetShortVideoID

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.shortVideoService.DeleteShortVideo(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	response := helper.APIresponse(http.StatusOK, newDel)
	c.JSON(http.StatusOK, response)
}
