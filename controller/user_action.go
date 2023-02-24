package controller

import (
	"fmt"
	"net/http"
	"time"

	"dhui.com/model"
	"github.com/gin-gonic/gin"
)

type querycomment struct {
	Uid      string `json:"uid"`
	Vid      int    `json:"vid"`
	Content  string `json:"content"`
	Reply_Id int    `json:"reply_id"`
}

func Send_Comment(c *gin.Context) {
	// uid := c.PostForm("uid")
	// vid := c.PostForm("vid")
	// content := c.PostForm("content")
	// reply_id := c.PostForm("reply_id")
	var q querycomment
	err := c.ShouldBindJSON(&q)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 4002,
			"message":     err.Error(),
		})
		return
	}
	fmt.Println(q)
	tkuid, ok := c.Get("uid")
	if ok && tkuid == q.Uid {
		comment := model.Comment{
			Uid:         q.Uid,
			V_id:        q.Vid,
			Content:     q.Content,
			Reply_Id:    q.Reply_Id,
			Create_Time: time.Now(),
		}
		err = model.DB.Create(&comment).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": 4001,
				"message":     err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "successful!",
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无权限!",
		})
		return
	}

}
