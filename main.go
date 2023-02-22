package main

import (
	"gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.GET("/ping", controller.Ping)
	e.Run()
}
