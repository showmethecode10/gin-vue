package main

import (
	"gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.GET("/ping", controller.Ping)

	e.GET("/api/user/register", controller.Register)

	err := e.Run()
	if err != nil {
		return
	}
}
