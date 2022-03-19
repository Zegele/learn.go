package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() //几行代码实现了web-server（服务器端）  轻量
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte(`你好，gin！`))
	})
	r.Run(":8080")
}
