package middleware

import (
	"gin-vue/dao"
	"gin-vue/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不够！",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := tools.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不够！",
			})
			c.Abort()
			return
		}
		//验证通过之后获取claim中的userId
		userId := claims.UserId
		user := dao.Mgr.GetUserByID(int(userId))
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不够！",
			})
			c.Abort()
			return
		}
		//用户存在,将user的信息写入context
		c.Set("user", user)
		c.Next()
	}
}
