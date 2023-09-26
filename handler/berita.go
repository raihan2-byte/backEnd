package handler

import (
	"blog/berita"
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
	barang, err := h.beritaService.FindByKarya()
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
	var imagesKitURLs []string

    // Loop through all file input fields (e.g., "file1", "file2", etc.)
    for i := 1; ; i++ {
        fileKey := fmt.Sprintf("file%d", i)
        file, err := c.FormFile(fileKey)
        
        // If there are no more files to upload, break the loop
        if err == http.ErrMissingFile {
            break
        }

        if err != nil {
            fmt.Printf("Error when opening file %s: %v\n", fileKey, err)
            continue // Skip to the next file
        }

        src, err := file.Open()
        if err != nil {
            fmt.Printf("Error when opening file %s: %v\n", fileKey, err)
            continue
        }
        defer src.Close()

        buf := bytes.NewBuffer(nil)
        if _, err := io.Copy(buf, src); err != nil {
            fmt.Printf("Error reading file %s: %v\n", fileKey, err)
            continue
        }

        img, err := imagekits.Base64toEncode(buf.Bytes())
        if err != nil {
            fmt.Printf("Error reading image %s: %v\n", fileKey, err)
            continue
        }

        fmt.Printf("Image base64 format %s: %v\n", fileKey, img)

        imageKitURL, err := imagekits.ImageKit(context.Background(), img)
        if err != nil {
            fmt.Printf("Error uploading image %s to ImageKit: %v\n", fileKey, err)
            continue
        }

        imagesKitURLs = append(imagesKitURLs, imageKitURL)
    }
		// if err != nil{
		// 	return err
		// }
		var input berita.CreateBerita

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Create a new news item with the provided input
newNews, err := h.beritaService.CreateBerita(input)
fmt.Println(newNews)
if err != nil {
    response := helper.APIresponse(http.StatusUnprocessableEntity, err)
    c.JSON(http.StatusUnprocessableEntity, response)
    return
}

// Associate the uploaded images with the news item
for _, imageURL := range imagesKitURLs {
    // Create a new BeritaImage record for each image and associate it with the news item
    err := h.beritaService.CreateBeritaImage(newNews.ID, imageURL)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
        c.JSON(http.StatusUnprocessableEntity, response)
        return
    }
}

// Respond with a success message
data := gin.H{"is_uploaded": true}
response := helper.APIresponse(http.StatusOK, data)
c.JSON(http.StatusOK, response)

}

func (h *beritaHandler) UpdateBerita (c *gin.Context){
	// file, err := c.FormFile("file")
	// if err != nil {
	// 	fmt.Printf("error when open file: %v", err)
	// 	return
	// }
	
	// src, err := file.Open()
	// if err != nil {
	// 	fmt.Printf("error when open file: %v", err)
	// 	return
	// }
	// defer src.Close()
	
	// buf:=bytes.NewBuffer(nil)
	// if _, err := io.Copy(buf, src); err != nil {
	// 	fmt.Printf("error read file %v",err)
	// 	return 
	// }	

	// img,err:=imagekits.Base64toEncode(buf.Bytes())
	// if err!=nil{
	// 	fmt.Println("error reading image %v",err)
	// }

	// fmt.Println("image base 64 format : %v",img)

	// imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	// if err != nil {
	// 	// Tangani jika terjadi kesalahan saat upload gambar
	// 	// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
	// 	response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
	// 	c.JSON(http.StatusInternalServerError, response)
	// 	return
	// }

	var inputID berita.GetBerita
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
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

	_, err = h.beritaService.UpdateBerita(inputID, input)
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
