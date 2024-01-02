package handler

import (
	endpointcount "blog/endpointCount"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	statisticsService endpointcount.StatisticsService
}

func NewStatisticsHandler(statisticsService endpointcount.StatisticsService) *StatisticsHandler {
	return &StatisticsHandler{
		statisticsService: statisticsService,
	}
}

func (h *StatisticsHandler) IncrementCountHandler(c *gin.Context) {
	endpoint := c.Request.URL.Path
	// userAgent := c.GetHeader("User-Agent")


	err := h.statisticsService.IncrementCount(endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to increment count"})
		return	
	}

	c.JSON(http.StatusOK, gin.H{"message": "Count incremented"})
}

func (h *StatisticsHandler) GetStatisticsHandler(c *gin.Context) {
	statistics, err := h.statisticsService.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve statistics"})
		return
	}

	c.JSON(http.StatusOK, endpointcount.FormatterGet(statistics))
}

func (h *StatisticsHandler) GetTotalCountForEndpointHandler(c *gin.Context) {
    endpoint := c.Param("endpoint") // Misalkan endpoint diambil dari parameter

    totalCount, err := h.statisticsService.GetTotalCountForEndpoint(endpoint)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve total count for endpoint"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"totalCount": totalCount})
}



