package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User_info struct {
	Username string `form:"username" json:"username" url:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" url:"password" xml:"password" binding:"required"`
}

func Login(c *gin.Context) {
	fmt.Println("Login")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello 好好",
	})
	// return
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is register 哈啊啊啊",
	})
	// return
}
