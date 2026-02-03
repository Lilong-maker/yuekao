package middleware

import (
	"net/http"
	"yuekao/bff/pkg"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "未登录",
			})
			return
		}
		getToken, err := pkg.GetToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "token解析失败",
			})
			return
		}

		c.Set("userId", getToken["userId"].(string))
		c.Next()
	}
}

//func Sx(c *gin.Context)  {
//	c.GetHeader()
//}
