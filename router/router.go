package router

import (
	"gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()

	e.POST("/api/user/register", controller.Register)

	err := e.Run()
	if err != nil {
		return
	}
}
