package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/controllers"

)

func UserRoutes(r *gin.Engine) {

	userGroup := r.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}
}
