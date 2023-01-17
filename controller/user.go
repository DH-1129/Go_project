package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"dhui.com/funcs"
	"dhui.com/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println("Login")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello 好好",
	})

}

func Get_Username() (name string) {
	var user model.User
	for {
		username := funcs.RandStringBytes(12)
		err := model.DB.Where("username=?", username).First(&user).Error
		if err != nil {
			return name
		}
	}
}

func Get_Uid() (uid string) {
	letter := "0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func Register(c *gin.Context) {
	// 邮箱注册或者电话号码注册
	// types := c.DefaultPostForm("type", "post")
	email := c.PostForm("email")
	v_code := c.PostForm("v_code")
	password := c.PostForm("password")
	// phone := c.PostForm("phone")
	_, err := model.GetEmail(email)
	if email != "" && v_code != "" && password != "" && err != nil && funcs.VerificationFormat(email) {
		user := model.User{
			Email:       email,
			Username:    Get_Username(),
			Uid:         Get_Username(),
			Password:    password,
			Img_url:     "wwww.image",
			Create_time: time.Now(),
			Last_time:   time.Now(),
		}
		err := model.DB.Create(&user).Error
		if err != nil {
			funcs.Danger("新用户插入失败! ", err)
			c.JSON(http.StatusOK, gin.H{
				"status_code": 5004,
				"message":     "数据库插入失败",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 4000,
				"message":     "用户创建成功",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 4003,
			"message":     "邮箱或者验证码或者密码没有填写或者邮箱格式不对！",
		})
	}

}

func Send_Verificaton(c *gin.Context) {
	email := c.PostForm("email")
	if email != "" {
		// 邮箱注册
		user, err := model.GetEmail(email)
		if err != nil {
			// 邮箱不存在，需要发送验证码
			code, err := funcs.Send_Email(email)
			if err != nil {
				funcs.Danger("邮箱注册验证码发送失败! ", err)
				c.JSON(http.StatusOK, gin.H{
					"status_code": 4003,
					"message":     "邮箱注册验证码发送失败! 请检查邮箱号或者格式是否正确！",
				})
			} else {
				// 保存code
				err = funcs.Set_Code(email, code, 120*time.Second)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"status_code": 4002,
						"message":     "邮箱注册验证码获取失败! 请重新获取！",
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"status_code": 4000,
						"message":     "验证码已发送,两分钟内有效！",
					})
				}
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 4001,
				"message":     "该邮箱已存在，无法重复注册！",
				"create_time": user.Create_time,
			})
		}

	}
}
