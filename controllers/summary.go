package controllers

import (
	"net/http"

	"butler_backend/services"
	structs "butler_backend/structs"

	"github.com/gin-gonic/gin"
)

func HandleRequestSummary(c *gin.Context) {
	var req structs.SummaryRequestReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid value",
		})
		return
	}

	result, err := services.RequestSummary(req.Uid, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func HandleGetSummaries(c *gin.Context) {
	uid := c.Query("uid")

	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid value",
		})
		return
	}

	result, err := services.GetSummaries(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func HandleGetSelectedSummary(c *gin.Context) {
	summaryId := c.Query("summaryId")

	if summaryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid value",
		})
		return
	}

	result, err := services.GetSelectedSummary(summaryId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func HandleGetRecentSummaries(c *gin.Context) {
	uid := c.Query("uid")

	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid value.",
		})
		return
	}

	result, err := services.GetRecentSummaries(uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}