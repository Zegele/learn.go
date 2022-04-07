package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn.go/chapter12/apiss"
	"learn.go/chapter12/frinterface"
	"learn.go/chapter12/rank"
	"log"
	"net/http"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/testdb"))

	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func main() {
	conn := connectDb()
	var rankServer frinterface.ServeInterface = NewDbRank(conn, rank.NewFatRateRank()) //这个初始化一定要看懂。
	// NewDbRank 的返回值已经是DbRank结构体，并实现了该接口。
	// NewDbRank中的参数conn是连接数据库，ranks.NewFatRate是返回一个实例化的结构体： &FatRateRank{items: make([]RankItem, 0, 100)}

	// var rankServer frinterface.ServeInterface = ranks.NewFatRateRank()

	if initRank, ok := rankServer.(frinterface.RankInitInterface); ok {
		if err := initRank.Init(); err != nil {
			log.Fatal("初始化失败", err)
		}
	}

	r := gin.Default()
	pprof.Register(r)
	r.POST("/register", func(c *gin.Context) {
		var pi *apiss.PersonalInformation
		if err := c.BindJSON(&pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法解析注册信息" + err.Error(),
			})
			return
		}
		if err := rankServer.RegisterPersonalInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "注册失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.PUT("/personalinfo", func(c *gin.Context) {
		var pi *apiss.PersonalInformation
		if err := c.BindJSON(&pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法解析注册信息",
			})
			return
		}
		if resp, err := rankServer.UpdatePersonalInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "更新失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, resp)
		}
	})

	r.GET("/ranks/:name", func(c *gin.Context) {
		name := c.Param("name") //http://localhost:8081/rank/111
		if fr, err := rankServer.GetFatRate(name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取个人信息失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, fr)
		}
	})

	r.GET("/ranktop", func(c *gin.Context) {
		if frTop, err := rankServer.GetTop(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取排行失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, frTop)
		}
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}

}
