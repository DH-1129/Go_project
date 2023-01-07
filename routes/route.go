package routes

import (
	"dhui.com/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/login", controller.Login)
		}
	}
	return r
}
