package controllers

import (
	"net/http"

	services "butler_backend/services"
	structs "butler_backend/structs"

	"github.com/gin-gonic/gin"
)

func HandleChatCreation(c *gin.Context) {
	var req structs.ChatCreationReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid value for chat creation.",
		})
		return
	}

	result, err := services.ChatCreation(req.Type, req.ChatName, req.Participants)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func HandleGetAllChats(c *gin.Context) {
	uid := c.Query("uid")
	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'uid is required.",
		})
		return
	}
	result, err := services.GetAllChats(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
