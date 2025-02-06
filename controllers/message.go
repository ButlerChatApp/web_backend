package controllers

import (
	"net/http"

	services "butler_backend/services"
	structs "butler_backend/structs"

	"github.com/gin-gonic/gin"
)

func HandlePostMessage(c *gin.Context) {
	var req structs.PostMessageReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request.",
		})
		return
	}

	result, err := services.PostMessage(req.ChatId, req.SenderId, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func HandleGetMessages(c *gin.Context) {
	chatId := c.Query("chatId")

	if chatId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ChatId is required",
		})
		return
	}

	result, err := services.GetMessages(chatId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func HandleEditMessage(c *gin.Context) {
	var param structs.EditMessageParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data. ",
		})
		return
	}

	result, err := services.EditMessage(param.MessageId, param.NewContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update message. " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func HandleDeleteMessage(c *gin.Context) {
	messageId := c.Query("messageId")

	if messageId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "MessageId is required.",
		})
		return
	}

	result, err := services.DeleteMessage(messageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete message." + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}