package middleware

import (
	"net/http"

	"dhui.com/funcs"
	"github.com/gin-gonic/gin"
)

// token认证
func Token_Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, err := funcs.ParseToken(c.Request.Header.Get("Authorization"))
		// fmt.Println(claim, err)
		if err == nil {
			// fmt.Println("claim", claim.Uid)
			c.Set("uid", claim.Uid)
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"status_code": 4000,
				"message":     "无权限或者token过期!",
			})

		}

	}
}
