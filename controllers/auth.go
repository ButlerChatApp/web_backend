package controllers

import (
	"net/http"

	"butler_backend/services"
	structs "butler_backend/structs"

	"github.com/gin-gonic/gin"
)

func HandleSignUp(c *gin.Context) {
	var req structs.SignUpReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, structs.SignUpRes{
			Message: "Failed to sign up: " + err.Error(),
		})
		return
	}

	result := services.SignUp(req.UserName, req.Email, req.Password)
	c.JSON(http.StatusCreated, result)
}

func HandleSignIn(c *gin.Context) {
	var req structs.SignInReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, structs.SignInRes{
			Message: "Failed to sign in: " + err.Error(),
		})
		return
	}

	result, err := services.SignIn(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}
	c.JSON(http.StatusOK, result)
}
