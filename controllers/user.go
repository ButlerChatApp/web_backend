package controllers

import (
	"net/http"

	services "butler_backend/services"
	structs "butler_backend/structs"

	"github.com/gin-gonic/gin"
)

func HandleGetUserName(c *gin.Context) {
	var req structs.GetUserNameReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid value",
		})
		return
	}

	result, err := services.GetUserName(req.Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, result)
}