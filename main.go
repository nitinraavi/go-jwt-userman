package main

import (
	"go-jwt/intializers"
	"go-jwt/models"
	"go-jwt/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDB()
	intializers.SyncDatabase()
}

var user models.User

func main() {
	r := gin.New()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Healthy",
		})
	})

	r.Use(gin.Logger())

	routes.AuthRoutes(r)
	r.Run(":" + os.Getenv("port")) // listen and serve on 0.0.0.0:8080
}
