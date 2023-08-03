package handler

import (
	"blog/berita"
	"blog/helper"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

type beritaHandler struct {
	beritaService berita.Service
}

func NewBeritaHandler(beritaService berita.Service) *beritaHandler {
	return &beritaHandler{beritaService}
}

func (h *beritaHandler) DeleteBerita(c *gin.Context){
	var input berita.GetBerita

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.beritaService.DeleteBerita(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
		
	}
	response := helper.APIresponse(http.StatusOK, "your news has been succesfuly deleted")
	c.JSON(http.StatusOK, response)

}

func (h *beritaHandler) GetOneBerita(c *gin.Context){
	var input berita.GetBerita

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.beritaService.GetOneBerita(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
		
	}
	response := helper.APIresponse(http.StatusOK, berita.FormatterBerita(newDel))
	c.JSON(http.StatusOK, response)

}

func (h *beritaHandler) GetByTags (c *gin.Context){
	var input berita.GetTags

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	barang, err := h.beritaService.FindByTags(input.Tags)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	response := helper.APIresponse(http.StatusOK, berita.FormatterGetBerita(barang))
	c.JSON(http.StatusOK, response)

}

func (h *beritaHandler) GetByKarya(c *gin.Context){
	var input berita.GetKarya

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	barang, err := h.beritaService.FindByKarya(input.Karya)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	response := helper.APIresponse(http.StatusOK, berita.FormatterGetBerita(barang))
	c.JSON(http.StatusOK, response)
}

func (h *beritaHandler) CreateBerita(c *gin.Context){
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

	img,err:=beritaBase64toEncode(buf.Bytes())
	if err!=nil{
		fmt.Println("error reading image %v",err)
	}

	fmt.Println("image base 64 format : %v",img)

	imageKitURL, err := beritaImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var input berita.CreateBerita

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

	_, err = h.beritaService.CreateBerita(input, imageKitURL)
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

func beritaBase64toEncode(bytes []byte) (string,error){
	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += beritatoBase64(bytes)

	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)
	return base64Encoding,nil
}

func beritatoBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func beritaImageKit(ctx context.Context, base64Image string) (string, error) {
	fmt.Println("start uploading image ...")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	ik := imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  "private_iitVYNY2fbOQSJgHtSccK9agJz0=",
		PublicKey:   "public_OtKeno1x/kY3zk5m7I9eNwnAtrY=",
		UrlEndpoint: "https://ik.imagekit.io/raihan2",
	})

	resp, err := ik.Uploader.Upload(ctx, base64Image, uploader.UploadParam{
		FileName: "test_image.jpg",
		Tags:     "barang",
		Folder:   "unj",
	})

	if err != nil {
		fmt.Printf("an error occurred when uploading image %v", err)
		return "", err
	}

	if resp.StatusCode != 200 {
		fmt.Printf("an error occurred when uploading image %v", resp)
		return "", errors.New("failed to upload image")
	}

	// Return the ImageKit URL
	return resp.Data.Url, nil
}


func (h *beritaHandler) GetAllBerita(c *gin.Context){
	input, _ := strconv.Atoi(c.Query("id"))

	newBerita, err := h.beritaService.GetAllBerita(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, "Eror")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, berita.FormatterGetBerita(newBerita))
	c.JSON(http.StatusOK, response)
}
