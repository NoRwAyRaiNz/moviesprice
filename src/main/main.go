package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
func main() {
	//1. 创建路由
	r := gin.Default()
	//2. 绑定路由规则和处理函数
	r.GET("/movie", func(c *gin.Context) {
		// 指定默认值
		// 走/user才会出默认值
		name := c.DefaultQuery("name","sam")
		//author := c.Param("author")
		//截取/
		//author = strings.Trim(author, "/")
		c.JSON(200, gin.H{
			"name": name,
			"message": "StatusOK",
		})
	})
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type","post")
		movie := c.PostForm("movie")
		author := c.PostForm("author")
		c.JSON(200, gin.H{
			"movie": movie,
			"author": author,
			"types": types,
			"message": "StatusOK",
		})
	})

	// 限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context){
		// gin 封装了http库中的Request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(500, gin.H{
				"message": "upload pic failed",
			})
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})

	r.POST("/filesize", func(c *gin.Context){
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file size: %v", err)
		}
		// 获取文件大小
		if headers.Size > 1024 * 1024 * 2 {
			fmt.Println("file size is too large")
			return
		}
		// headers.Header.Get("Content-Type") 获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "./pic/" + headers.Filename)
		c.String(http.StatusOK, headers.Filename + "yes")
	})

	r.PUT("/xxPut", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT succeed!")
	})
	r.DELETE("xxDelete", func(c *gin.Context) {
		c.String(http.StatusOK,"delete succeed!")
	})
	//3.监听端口
	r.Run(":8080")


}

