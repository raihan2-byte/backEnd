package handler

import (
	"blog/artikel"
	"blog/berita"
	"blog/helper"
	"encoding/base64"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type artikelHandler struct {
	artikelService artikel.Service
}

func NewArtikelHandler(artikelService artikel.Service) *artikelHandler {
	return &artikelHandler{artikelService}
}

func (h *artikelHandler) DeleteArtikel(c *gin.Context) {
	var input artikel.GetArtikel

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.artikelService.DeleteArtikel(input.ID)
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

func (h *artikelHandler) GetOneArtikel(c *gin.Context) {
	var input berita.GetBerita

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.artikelService.GetOneArtikel(input.ID)
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

func (h *artikelHandler) CreateArtikel (c *gin.Context){
	var input artikel.CreateArtikel

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := handleBase64Upload("file")

	if err != nil {
		//inisiasi data yang tujuan dalam return hasil ke postman
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// currentUser := c.MustGet("currentUser").(user.User)
	// //ini inisiasi userID yang mana ingin mendapatkan id si user
	// // input.User = currentUser
	// userID := currentUser.ID

	// path := fmt.Sprintf("images/%d-%s", userID, file)

	// err = c.SaveUploadedFile(file, path)
	// if err != nil {
	// 	data := gin.H{"is_uploaded": false}
	// 	response := helper.APIresponse(http.StatusUnprocessableEntity, data)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	_, err = h.artikelService.CreateArtikel(input, file)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

func handleBase64Upload(base64String string) (string, error) {
	// Menghapus "data:image/png;base64," dari base64String
	base64String = strings.TrimPrefix(base64String, "data:image/png;base64,")

	// Decode base64String menjadi byte array
	imageData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return "", err
	}
	// currentUser := c.MustGet("currentUser").(user.User)
	// //ini inisiasi userID yang mana ingin mendapatkan id si user
	// // input.User = currentUser
	// userID := currentUser.ID

	// Simpan byte array menjadi file gambar
	path := "path/to/"
	err = os.WriteFile(path, imageData, 0644)
	if err != nil {
		return "", err
	}

	return path, nil
}


func (h *artikelHandler) GetAllArtikel(c *gin.Context){
	input, _ := strconv.Atoi(c.Query("id"))

	newBerita, err := h.artikelService.GetAllArtikel(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, "Eror")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, newBerita)
	c.JSON(http.StatusOK, response)
}
