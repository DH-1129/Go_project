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
			v1.POST("/upload/video", middleware.Token_Auth(), controller.Upload_Video)
			v1.POST("/del/video", middleware.Token_Auth(), controller.Del_Video)
			v1.POST("/del/comment", middleware.Token_Auth(), controller.Del_Comment)
			v1.POST("/commit/comment", middleware.Token_Auth(), controller.Send_Comment)
			v1.POST("/commit/like", middleware.Token_Auth(), controller.User_Like)
			v1.POST("/commit/unlike", middleware.Token_Auth(), controller.User_Unlike)
			v1.POST("/commit/collection", middleware.Token_Auth(), controller.User_Collection)
			v1.POST("/commit/uncollection", middleware.Token_Auth(), controller.User_Uncollection)
		}
		v2 := api.Group("v2")
		{
			v2.POST("/update/password", controller.Update_Password)
			v2.POST("/update/info", middleware.Token_Auth(), controller.Update_Info)
			v2.POST("/update/img", middleware.Token_Auth(), controller.Update_IMG)
			v2.GET("/get_collections", middleware.Token_Auth(), controller.Get_Collections)
			v2.GET("/get_likes", middleware.Token_Auth(), controller.Get_Likes)

		}
	}
	return r
}
