package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"dhui.com/funcs"
	"dhui.com/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 随机生成用户名
func Get_Username(name_len int) (name string) {
	var user model.User
	for {
		username := funcs.RandStringBytes(name_len)
		err := model.DB.Where("username=?", username).First(&user).Error
		if err != nil {
			return username
		}

	}
}

// 随机生成8为位的uid
func Get_Uid(uid_len int) (uid string) {
	var user model.User
	for {
		fmt.Println("开始生成uid")
		letter := "0123456789"
		b := make([]byte, uid_len)
		for i := range b {
			b[i] = letter[rand.Intn(len(letter))]
		}
		err := model.DB.Where("uid=?", string(b)).First(&user).Error
		if err != nil {
			return string(b)
		}
	}

}

// 用户登录
func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println(email != " ", password != " ")
	if email != " " && password != " " {
		fmt.Println("raSr")
		var user model.User
		err := model.DB.Where("email=?", email).First(&user).Error
		fmt.Println("err: ", err)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 4004,
				"message":     "密码或者邮箱不正确!",
			})
		} else {
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"Status_code": 4001,
					"message":     "密码或者邮箱不正确!",
				})
			} else {
				user.Password = ""
				c.JSON(http.StatusOK, gin.H{
					"status_code": 200,
					"message":     "登录成功",
					"token":       funcs.Get_Token(user.Uid, user.Username),
					"data":        user,
				})
			}
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": "4001",
			"message":     "邮箱或者密码不完整!",
		})
	}

}

// 邮箱注册
func Register(c *gin.Context) {

	email := c.PostForm("email")
	v_code := c.PostForm("v_code")
	password := c.PostForm("password")

	if email != " " && v_code != " " && password != " " && funcs.VerificationFormat(email) {
		flag := funcs.Veri_Code(v_code, email) // 校验验证码
		if flag {
			user := model.User{
				Email:       email,
				Username:    Get_Username(12),
				Uid:         Get_Uid(8),
				Password:    funcs.Encrypthon_PW(password),
				Img_url:     "wwww.image",
				Sex:         "未知",
				Create_Time: time.Now(),
				Update_Time: time.Now(),
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
					"status_code": 2000,
					"message":     "用户创建成功",
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 4006,
				"message":     "验证码不对或已过期",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 4003,
			"message":     "邮箱格式不正确或该邮箱已存在",
		})
	}

}

// 发送邮箱注册验证码
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
				"create_time": user.Create_Time,
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 4004,
			"message":     "请输入正确的邮箱！",
		})
	}
}

// 用户密码修改
func Update_Password(c *gin.Context) {
	new_ps := c.PostForm("new_password")
	uid := c.PostForm("uid")
	var user model.User
	if new_ps != "" && uid != "" {
		err := model.DB.Where("uid=?", uid).First(&user).Error
		// fmt.Println(err, user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 4000,
				"message":     "用户密码修改失败!",
			})
		} else {
			model.DB.Model(&user).Select("password", "update_time").UpdateColumns(model.User{Password: funcs.Encrypthon_PW(new_ps), Update_Time: time.Now()})
			c.JSON(http.StatusOK, gin.H{
				"status_code": 2000,
				"message":     "用户密码修改成功!",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 4004,
			"message":     "新密码未输入!",
		})
	}
}

// 用户其他信息修改

func Update_Info(c *gin.Context) {
	// fmt.Printf("%T", c.Get("claim"))
	new_username := c.PostForm("new_username")
	new_sex := c.PostForm("new_sex")
	new_signature := c.PostForm("new_signature")
	uid := c.PostForm("uid")
	tkuid, ok := c.Get("uid") // 避免利用合法token修改别人信息
	// fmt.Println(ok, uid, tkuid, uid == tkuid)
	if ok && uid != "" && uid == tkuid {
		var user model.User
		err := model.DB.Where("uid = ?", uid).First(&user).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 4000,
				"message":     "该用户uid不存在!",
			})
		} else {
			user.Username = new_username
			user.Sex = new_sex
			user.Signature = new_signature
			model.DB.Model(&user).Select("username", "sex", "signature").UpdateColumns(model.User{Username: user.Username, Sex: user.Sex, Signature: user.Signature})
			c.JSON(http.StatusOK, gin.H{
				"status_code": 2000,
				"message":     "信息修改成功!",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 4000,
			"message":     "授权用户和当前修改用户不一致!",
		})
	}
}
