package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	//获取参数
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	//数据校验
	//创建用户
	//返回结果

	c.JSON(http.StatusOK, gin.H{
		"msg": "register success!",
	})
}
