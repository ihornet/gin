package http

import (
	"gin/middleware"
	"gin/modules/user"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Init(port string) {

	// 禁用控制台颜色
	//gin.DisableConsoleColor()

	// 创建记录日志的文件
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("user/:id", middleware.MiddleWare(), user.GetInfo)
	router.POST("/user", user.AddUser)

	router.Run(port)

}
