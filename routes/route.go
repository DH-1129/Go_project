package routes

import (
	"dhui.com/controller"
	"dhui.com/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/login", controller.Login)
			v1.POST("/register", controller.Register)
			v1.POST("/verification_code", controller.Send_Verificaton)

			v1.POST("/update/password", controller.Update_Password)
			// v1.Use(gin.BasicAuth(gin.Accounts{
			// 	"admin": "123456",
			// }))
			v1.Use(middleware.Token_Auth())
			v1.POST("/update/info", controller.Update_Info)
		}
	}
	return r
}
