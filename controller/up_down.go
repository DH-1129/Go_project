package controller

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"dhui.com/configs"
	"dhui.com/funcs"
	"dhui.com/model"
	"github.com/gin-gonic/gin"
)

// 随机生成36位的文件名
func Get_V_Name(v_name_len int) (uid string) {
	var user model.Video_Info
	for {
		fmt.Println("开始生成video_name")
		letter := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM0123456789"
		b := make([]byte, v_name_len)
		for i := range b {
			b[i] = letter[rand.Intn(len(letter))]
		}
		err := model.DB.Where("file_name=?", string(b)).First(&user).Error
		fmt.Println(err)
		if err != nil {
			return string(b)
		}
	}

}

// 视频上传
func Upload_Video(c *gin.Context) {
	uid := c.PostForm("uid")
	title := c.PostForm("title")
	content := c.PostForm("content")
	tag := c.PostForm("tag")
	tkuid, ok := c.Get("uid")
	if ok && tkuid == uid {
		file, err := c.FormFile("file")
		if err != nil {
			log.Fatal(err)
			c.String(http.StatusOK, "Fatal!")
		}
		file_name := Get_V_Name(36) + ".pm4"
		err = c.SaveUploadedFile(file, configs.VIDEO_SAVE_DIR+"/"+file_name)
		if err != nil {
			funcs.Danger(err)
			log.Fatal(err)
			c.JSON(http.StatusOK, gin.H{
				"status_code": 4001,
				"message":     "视频保存失败!",
			})
		} else {
			video_info := model.Video_Info{
				Uid:         uid,
				Title:       title,
				Content:     content,
				Tag:         tag,
				Create_Time: time.Now(),
				Prefix:      configs.VIDEO_SAVE_DIR,
				File_Name:   file_name,
			}
			err := model.DB.Create(&video_info).Error
			if err != nil {
				funcs.Warning(err)
				c.JSON(http.StatusOK, gin.H{
					"status_code": 4001,
					"message":     "视频信息入库失败!",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status_code": 2000,
					"message":     "success!",
				})
			}
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 4000,
			"message":     "权限错误!",
		})
	}

}

func Del_Video(c *gin.Context) {
	uid := c.PostForm("uid")
	v_id := c.PostForm("v_id")
	v_name := c.PostForm("v_name")
	tkuid, ok := c.Get("uid")
	if ok && uid == tkuid {
		v_info := model.Video_Info{}
		fmt.Println(configs.VIDEO_SAVE_DIR + "/" + v_name)
		err := os.Remove(configs.VIDEO_SAVE_DIR + "/" + v_name)
		if err != nil {
			log.Fatal(err)
			c.String(http.StatusOK, "删除失败!")
		}
		err = model.DB.Where("id = ?", v_id).Delete(&v_info).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 4001,
				"message":     "删除失败!",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 2000,
				"message":     "删除成功!",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 4000,
			"message":     "权限错误!",
		})
	}

}
