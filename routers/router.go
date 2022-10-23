package routers

import (
	"final_project/controllers"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	router := gin.Default()

	router.GET("/users/", controllers.GetUsers)
	router.POST("/users/register/", controllers.RegisterUser)
	router.POST("/users/login/", controllers.LoginUser)
	router.PUT("/users/:userId", controllers.UpdateUser)

	return router
}
