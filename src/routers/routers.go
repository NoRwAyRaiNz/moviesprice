package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
)
// 定义接收数据的结构体
type Movie struct {
	// 必须字段 binding:"required"
	Name string `form:"moviename" json:"moviename" uri:"moviename" xml:"moviename" binding:"required"`
	Price float64 `form:"movieprice" json:"movieprice" uri:"movieprice" xml:"movieprice" binding:"required"`
}

// 具体处理函数
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": " Hello, welcome to my website!",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", helloHandler)
	// TODO 拆分路由 https://www.topgoer.com/gin框架/gin路由/路由拆分与注册.html
	//2. 绑定路由规则和处理函数
	r.GET("/movie", func(c *gin.Context) {
		// 指定默认值
		// 走/user才会出默认值
		name := c.DefaultQuery("name", "sam")
		//author := c.Param("author")
		//截取/
		//author = strings.Trim(author, "/")
		c.JSON(200, gin.H{
			"name":    name,
			"message": "StatusOK",
		})
	})
	r.POST("/postandget", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		types := c.DefaultPostForm("type", "post")
		movie := c.PostForm("movie")
		author := c.PostForm("author")
		c.JSON(200, gin.H{
			"movie":   movie,
			"author":  author,
			"types":   types,
			"id": 	id,
			"page": 	page,
		})
	})

	// 限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// gin 封装了http库中的Request
		// 单文件
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
	r.POST("/multupload", func(c *gin.Context) {
		//多文件上传
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// 上传到文件指定路径
			c.SaveUploadedFile(file, file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))
	})
	r.POST("/filesize", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file size: %v", err)
		}
		// 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("file size is too large")
			return
		}
		// headers.Header.Get("Content-Type") 获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "./pic/"+headers.Filename)
		c.String(http.StatusOK, headers.Filename+"yes")
	})
	//JSON 数据解析和绑定
	r.POST("/QueryJSON", func(c *gin.Context) {
		// 声明接收变量
		var json Movie
		// 将request的body中的数据 自动按json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			//gin.H 封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断字段是否正确
		if json.Name != "1917" || json.Price != 100 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	// 表单数据解析和绑定
	r.POST("/formdata", func(c *gin.Context) {
		var form Movie
		//Bind() 默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断字段是否正确
		if form.Name != "1917" || form.Price != 100 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	//URI数据解析和绑定
	r.GET("/:moviename/:movieprice", func(c *gin.Context) {
		var movie Movie

		if err := c.ShouldBindUri(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断字段是否正确
		if movie.Name != "1917" || movie.Price != 100 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	// 结构体响应
	r.GET("/someStruct", func(c *gin.Context){
		var msg struct {
			Name string
			Message string
			Number int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})
	// XML response
	r.GET("/someXML", func(c *gin.Context){
		c.XML(200, gin.H{"message" : "XML"})
	})
	// YAML response
	r.GET("/someYAML", func(c *gin.Context){
		c.YAML(200, gin.H{"name" : "YAML"})
	})
	// protobuf
	r.GET("/someProtoBuf", func(c *gin.Context){
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		c.ProtoBuf(200, data)
	})
	// 重定向
	r.GET("toapple", func(c *gin.Context){
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	r.PUT("/xxPut", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT succeed!")
	})
	r.DELETE("xxDelete", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE succeed!")
	})
	r.PATCH("/xxPatch",func(c *gin.Context) {
		c.String(http.StatusOK, "PATCH succeed!")
	})
	r.HEAD("/xxHead",func(c *gin.Context) {
		c.String(http.StatusOK, "HEAD succeed!")
	})
	r.OPTIONS("/xxOptions",func(c *gin.Context) {
		c.String(http.StatusOK, "OPTIONS succeed!")
	})
	return r
}


