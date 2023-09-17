package handler

import (
	barang "blog/Barang"
	"blog/helper"
	"blog/imagekits"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type barangHandler struct {
	barangService barang.Service
}

func NewBarangHandler(barangService barang.Service) *barangHandler {
	return &barangHandler{barangService}
}

func (h *barangHandler) CreateBarang(c *gin.Context){
	file, _ := c.FormFile("file")
	src,err:=file.Open()
	defer	src.Close()
	if err!=nil{
		fmt.Printf("error when open file %v",err)
	}
	
	buf:=bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v",err)
		return 
	}	

	img,err:=imagekits.Base64toEncode(buf.Bytes())
	if err!=nil{
		fmt.Println("error reading image %v",err)
	}

	fmt.Println("image base 64 format : %v",img)

	imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var input barang.InputBarang

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

	_, err = h.barangService.CreateBarang(input, imageKitURL)
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

func (h *barangHandler) GetAllBarang (c *gin.Context){
	input, _ := strconv.Atoi(c.Query("id"))

	data, err := h.barangService.GetAllBarang(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, barang.FormatterGetCategory(data))
	c.JSON(http.StatusOK, response)
}

func (h *barangHandler) GetOneBarang (c *gin.Context){
	var input barang.GetBarang

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.barangService.GetOneBarang(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, barang.FormatterBarang(data))
	c.JSON(http.StatusOK, response)
}

func (h *barangHandler) UpdateBarang (c *gin.Context){

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("error when open file: %v", err)
		return
	}
	
	src, err := file.Open()
	if err != nil {
		fmt.Printf("error when open file: %v", err)
		return
	}
	defer src.Close()
	
	buf:=bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v",err)
		return 
	}	

	img,err:=imagekits.Base64toEncode(buf.Bytes())
	if err!=nil{
		fmt.Println("error reading image %v",err)
	}

	fmt.Println("image base 64 format : %v",img)

	imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var inputID barang.GetBarang
	err = c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input barang.InputBarang

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

	_, err = h.barangService.UpdateBarang(inputID, input, imageKitURL)
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

func (h *barangHandler) GetCategoryBarang (c *gin.Context){
	var input barang.GetCategory

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// input, _ := strconv.Atoi(c.Query("id"))

	data, err := h.barangService.GetBarangByCategory(input.Category)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, barang.FormatterGetCategory(data))
	c.JSON(http.StatusOK, response)
}

func (h *barangHandler) DeleteBarang (c *gin.Context){
	var input barang.GetBarang

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.barangService.DeleteBarang(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, "Your Barang has been succesfuly deleted")
	c.JSON(http.StatusOK, response)
}
