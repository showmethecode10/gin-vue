package router

import (
	"gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()

	//e.Static("/static", "static")
	//e.LoadHTMLGlob("template/*")

	e.GET("/")

	//router group
	api := e.Group("api")
	{
		//用户路由
		user := api.Group("user")
		{
			user.POST("/register", controller.Register)
			user.GET("/login", controller.Login)
		}
	}

	err := e.Run()
	if err != nil {
		return
	}
}
