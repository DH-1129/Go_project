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
			v1.POST("/update/info", middleware.Token_Auth(), controller.Update_Info)
			v1.POST("/update/img", middleware.Token_Auth(), controller.Update_IMG)
			v1.POST("/upload/video", middleware.Token_Auth(), controller.Upload_Video)
			v1.POST("/del/video", middleware.Token_Auth(), controller.Del_Video)
		}
	}
	return r
}
