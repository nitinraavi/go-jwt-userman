package routes

import (
	"go-jwt/controllers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", controllers.SignUp)
	incomingRoutes.POST("/signin", controllers.SignIn)
	incomingRoutes.GET("/validatetoken", middleware.RequireAuth, controllers.Validate)
}
