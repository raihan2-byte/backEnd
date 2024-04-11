package handler

import (
	endpointcount "blog/endpointCount"
	"blog/helper"
	"blog/home"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type homeHandler struct {
	homeService     home.Service
	endpointService endpointcount.StatisticsService
}

func NewHomeHandler(homeService home.Service, endpointService endpointcount.StatisticsService) *homeHandler {
	return &homeHandler{homeService, endpointService}
}

func (h *homeHandler) CreateTagHome(c *gin.Context) {
	var input home.CreateTagHome

	err := c.ShouldBind(&input)

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

	data, err := h.homeService.CreateTagHome(input)
	if err != nil {
		// data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

func (h *homeHandler) UpdateTagHome(c *gin.Context) {
	var inputID home.GetIdTagHome
	var input home.CreateTagHome

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

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

	data, err := h.homeService.UpdateTagHome(inputID.ID, input)
	if err != nil {
		// data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

func (h *homeHandler) GetAllTagHome(c *gin.Context) {

	input, _ := strconv.Atoi(c.Query("id"))

	newBerita, err := h.homeService.GetAllTagLineHome(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, "Eror")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// err = h.endpointService.IncrementCount("View-Site")
	// if err != nil {
	// 	response := helper.APIresponse(http.StatusUnprocessableEntity, err)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	response := helper.APIresponse(http.StatusOK, newBerita)
	c.JSON(http.StatusOK, response)

}

func (h *homeHandler) DeleteHome(c *gin.Context) {
	var input home.GetIdTagHome

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.homeService.DeleteTagLineHome(input.ID)
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
