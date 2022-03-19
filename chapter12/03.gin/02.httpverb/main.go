package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"learn.go/chapter12/apiss"
	"learn.go/chapter12/frinterface"
	"learn.go/chapter12/rank"
	"net/http"
)

func main() {
	var rankServer frinterface.ServeInterface = rank.NewFatRateRank()

	r := gin.Default()
	pprof.Register(r)

	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte(`你好，gin！`))
	})

	r.GET("/history", func(c *gin.Context) {
		c.JSON(200, []*apiss.PersonalInformation{ // 如果用JSON，就可以返回json格式，并自带content-type: application/json
			{
				Name:   "小强",
				Sex:    "男",
				Tall:   1,
				Weight: 2,
				Age:    3,
			},
			{
				Name:   "小强",
				Sex:    "男",
				Tall:   1.8,
				Weight: 65,
				Age:    33,
			},
		})
	})

	r.POST("/register", func(c *gin.Context) {
		pi := &apiss.PersonalInformation{}
		if err := c.BindJSON(pi); err != nil {
			c.JSON(400, gin.H{
				"message": "无法读取个人信息",
			})
			return
		}
		//todo 注册到排行榜 不知道是否作对
		if err := rankServer.RegisterPersonalInformation(pi); err != nil {
			c.Writer.Write([]byte(fmt.Sprintf("注册失败： %s", err)))
			//c.Writer([]byte(fmt.Sprintf("注册失败： %s", err)))
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(`success`))

		c.JSON(200, nil)
	})
	//r.PUT（一般用于更新） r.Delete（是删除，不是隐藏）等在ppt里。

	r.GET("/getwithquery", func(c *gin.Context) {
		//http://localhost:8081/getwithquery?name=xiaoqiang1  //这个url长一些
		//得到：{"payload":"eGlhb3FpYW5nMQ=="}
		name := c.Query("name")
		c.JSON(200, gin.H{
			"payload": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})

	r.GET("/getwithparam/:name", func(c *gin.Context) { ///:name" 注意这里有冒号，使用中没有冒号
		//http://localhost:8081/getwithparam/xiaoqiang1 // 结果和上面一样，但是url短一些。
		//得到：{"payload":"eGlhb3FpYW5nMQ=="}
		name := c.Param("name")
		c.JSON(200, gin.H{
			"payload": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})

	r.GET("/gettwithparam/xiaoqiang", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"special": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})

	r.Run(":8081")
}
