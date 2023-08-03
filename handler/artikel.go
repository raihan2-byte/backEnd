package handler

import (
	"blog/artikel"
	"blog/helper"
	"blog/user"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
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
	var input artikel.GetArtikel

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
	file, err := c.FormFile("file")
	src,err:=file.Open()
	defer 	src.Close()
	if err!=nil{
		fmt.Printf("error when open file %v",err)
	}
	
	buf:=bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v",err)
		return 
	}	

	img,err:=y(buf.Bytes())
	if err!=nil{
		fmt.Println("error reading image %v",err)
	}

	fmt.Println("image base 64 format : %v",img)

	x(context.Background(),img)
	

	var input artikel.CreateArtikel

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

	currentUser := c.MustGet("currentUser").(user.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	// input.User = currentUser
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.artikelService.CreateArtikel(input, path)
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

func y(bytes []byte) (string,error){
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
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)
	return base64Encoding,nil
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func x(ctx context.Context, base64Image string){
	fmt.Println("start uploading image ...")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()
	ik:=imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey: "private_iitVYNY2fbOQSJgHtSccK9agJz0=",
		PublicKey: "public_OtKeno1x/kY3zk5m7I9eNwnAtrY=",
		UrlEndpoint: "https://ik.imagekit.io/raihan2",
	})
	
	resp,err:=ik.Uploader.Upload(ctx,base64Image,uploader.UploadParam{
		FileName: "test_image.jpg",
		Tags: "barang",
		Folder: "unj",
	})

	if err!=nil{
		fmt.Printf("an error occurred when uploading image %v",err)
	}

	if resp.StatusCode!=200{
		fmt.Printf("an error occurred when uploading image %v",resp)
	}

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
