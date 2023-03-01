package controller

import (
	"log"
	"net/http"
	"time"

	"dhui.com/funcs"
	"dhui.com/model"
	"github.com/gin-gonic/gin"
)

type querycomment struct {
	Uid      string `json:"uid"`
	Vid      int    `json:"vid"`
	Content  string `json:"content"`
	Reply_Id int    `json:"reply_id"`
}

// 发送评论
func Send_Comment(c *gin.Context) {
	var q querycomment
	err := c.ShouldBindJSON(&q)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 4002,
			"message":     err.Error(),
		})
		return
	}
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
			// 视频评论数量+1
			var v_info model.Video_Info
			err := model.DB.Where("id = ?", q.Vid).First(&v_info).Error
			if err != nil {
				log.Fatal(err)
				return
			}
			v_info.Comment_Count += 1
			err = model.DB.Model(&v_info).Select("comment_count").UpdateColumns(model.Video_Info{Comment_Count: v_info.Comment_Count}).Error
			if err != nil {
				log.Fatal(err)
				return
			}
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

type d_c struct {
	Uid  string
	C_id int
}

// 删除评论
func Del_Comment(c *gin.Context) {
	var q d_c
	err := c.ShouldBindJSON(&q)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 4002,
			"message":     err.Error(),
		})
		return
	}
	uid := q.Uid
	c_id := q.C_id
	tkuid, ok := c.Get("uid")
	if ok && tkuid == uid {
		var comment model.Comment
		err := model.DB.Where("id = ?", c_id).Delete(&comment).Error
		if err != nil {
			funcs.Warning(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "删除Success!")
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无权限!",
		})
		return
	}
}

type querylike struct {
	Uid string `json:"uid"`
	Vid int    `json:"vid"`
}

// 视频点赞
// func Like_Action(c *gin.Context) {
// 	var q querylike
// 	err := c.ShouldBindJSON(&q)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	tkuid, ok := c.Get("uid")
// 	if ok && tkuid == q.Uid {
// 		var v_info model.Video_Info
// 		err := model.DB.Where("id = ?", q.Vid).First(&v_info).Error
// 		if err != nil {
// 			// log.Fatal(err)
// 			c.String(http.StatusBadRequest, err.Error())
// 			funcs.Warning(err)
// 			return
// 		}
// 		v_info.Like_Count += 1
// 		err = model.DB.Model(&v_info).Select("like_count").UpdateColumns(model.Video_Info{Like_Count: v_info.Like_Count}).Error
// 		if err != nil {
// 			funcs.Warning(err)
// 			c.String(http.StatusBadRequest, err.Error())
// 			return
// 		} else {
// 			// 为用户添加点赞的vid
// 			var user_info model.User
// 			err = model.DB.Where("uid=?", q.Uid).First(&user_info).Error
// 			if err != nil {
// 				funcs.Warning(err)
// 				c.String(http.StatusBadRequest, err.Error())
// 				return
// 			}
// 			string_vid := strconv.Itoa(q.Vid)
// 			likes_string := ""
// 			if user_info.Likes != "" {
// 				likes_arr := strings.Split(user_info.Likes, ",")
// 				likes_arr = append(likes_arr, string_vid)
// 				for _, value := range likes_arr {
// 					likes_string = likes_string + "," + value
// 				}
// 			} else {
// 				likes_string = string_vid
// 			}
// 			err = model.DB.Select("likes").UpdateColumns(model.User{Likes: likes_string}).Error
// 			if err != nil {
// 				funcs.Warning(err)
// 				c.String(http.StatusBadRequest, err.Error())
// 				return
// 			}
// 			c.String(http.StatusOK, "Success!")
// 			return
// 		}
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "无权限!",
// 		})
// 		return
// 	}

// }

func User_Like(c *gin.Context) {
	var q querylike
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数传递有误!",
		})
		return
	}
	tkuid, ok := c.Get("uid")
	if ok && tkuid == q.Uid {
		if err := funcs.NewRedisLikes().Like(q.Uid, q.Vid); err != nil {
			funcs.Warning(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "点赞失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Success!",
		})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无权限!",
		})
		return
	}
}
func User_Unlike(c *gin.Context) {
	var q querylike
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数传递有误!",
		})
		return
	}
	tkuid, ok := c.Get("uid")
	if ok && tkuid == q.Uid {
		if err := funcs.NewRedisLikes().Unlike(q.Uid, q.Vid); err != nil {
			funcs.Warning(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "取消失败!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "取消成功!",
		})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无权限!",
		})
		return
	}
}

// 获取点赞列表
func Get_Likes(c *gin.Context) {
	uid := c.Query("uid")
	tkuid, ok := c.Get("uid")
	if ok && tkuid == uid {
		arr, err := funcs.NewRedisLikes().GetLikedObjects(uid)
		if err != nil {
			funcs.Warning(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "列表获取失败",
				"data":    err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "列表获取成功",
			"data":    arr,
		})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无权限!",
		})
		return
	}

}

// 用户收藏
func User_Collection(c *gin.Context) {
	var q querylike
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数传递有误!",
		})
		return
	}
	tkuid, ok := c.Get("uid")
	if ok && tkuid == q.Uid {
		if err := funcs.NewRedisLikes().Collection(q.Uid, q.Vid); err != nil {
			funcs.Warning(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "收藏失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "收藏成功!",
		})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无权限!",
		})
		return
	}
}
func User_Uncollection(c *gin.Context) {
	var q querylike
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数传递有误!",
		})
		return
	}
	tkuid, ok := c.Get("uid")
	if ok && tkuid == q.Uid {
		if err := funcs.NewRedisLikes().Uncollection(q.Uid, q.Vid); err != nil {
			funcs.Warning(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "取消失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "收藏已移除!",
		})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无权限!",
		})
		return
	}
}

// 获取收藏列表
func Get_Collections(c *gin.Context) {
	uid := c.Query("uid")
	tkuid, ok := c.Get("uid")
	if ok && tkuid == uid {
		arr, err := funcs.NewRedisLikes().GetCollectionObjects(uid)
		if err != nil {
			funcs.Warning(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "列表获取失败",
				"data":    err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "列表获取成功",
			"data":    arr,
		})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无权限!",
		})
		return
	}

}
