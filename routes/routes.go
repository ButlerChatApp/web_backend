package routes

import (
	"net/http"
	"time"

	controllers "butler_backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // フロントエンドのURLを指定
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/api/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Butler !!!",
		})
	})

	r.POST("/api/signup", func(c *gin.Context) {
		controllers.HandleSignUp(c)
	})

	r.POST("/api/signin", func(c *gin.Context) {
		controllers.HandleSignIn(c)
	})

	return r
}
