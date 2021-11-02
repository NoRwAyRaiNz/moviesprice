package main

import (
	"MoviesPrice/src/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)




func main() {
	//强制开启日志颜色
	gin.ForceConsoleColor()
	// 创建日志文件
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
	//1. 创建路由
	r := routers.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
