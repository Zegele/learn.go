package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn.go/zuoye/zuoye6_tizhi_dbweb/api"
	"learn.go/zuoye/zuoye6_tizhi_dbweb/dbshow"
	"learn.go/zuoye/zuoye6_tizhi_dbweb/showinterface"
	"log"
	"net/http"
	"strconv"
)

func connectDb() *gorm.DB { //
	conn, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/testdb")) //与数据库建立连接

	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func main() {
	conn := connectDb()

	var ShowServer showinterface.ServeInterface = dbshow.NewDbShow(conn)
	// 和数据库对接 完成
	//创建20个数据
	for i := 0; i < 20; i++ {
		var ps = &api.PersonalShow{
			PersonID: 1,
			Name:     "a" + strconv.Itoa(i+1),
			//ShowTime: time.Now().Unix(),
			//ShowTime:      ,//时间格式和数据库怎么匹配？？？
			ShowDescription: "初来乍到，出来炸道",
			Weight:          float32(70 + i),
			Tall:            1.8 + float32(i),
			Age:             int64(20 + i),
			Visiable:        true,
		}
		if err := ShowServer.SaveShowInformation(ps); err != nil { //通过
			fmt.Println(err)
		}
	}

	//修改一个数据
	psUpdate := &api.PersonalShow{
		Id:              166,
		PersonID:        1,
		Name:            "a11",
		ShowDescription: "我怎么这么优秀？？？",
		Weight:          70.0,
		Tall:            1.8,
		Age:             30,
		Visiable:        true,
	}
	ShowServer.UpdatePersonalInformation(psUpdate) //通过

	// 浏览所有数据
	if _, err := ShowServer.GetShow(); err != nil { //通过
		fmt.Println(err)
	}

	//真删除
	for i := 4; i < 161; i++ {
		ShowServer.DeleteTrue(int64(i), 1) //通过
	}

	//假删除
	ShowServer.DeleteFalse(170, 1) //通过

	// 和web对接  get可以获取数据。但是post怎么增加和删除数据？？？不会
	r := gin.Default() //生成了一个实例
	pprof.Register(r)

	//对接注册
	r.POST("/register", func(c *gin.Context) {
		var ps *api.PersonalShow
		if err := c.BindJSON(&ps); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法解析注册信息" + err.Error(),
			})
			return
		}
		if err := ShowServer.SaveShowInformation(ps); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "注册失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	//对接更新
	r.PUT("/update", func(c *gin.Context) {
		var ps *api.PersonalShow
		if err := c.BindJSON(&ps); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "更新失败",
			})
			return
		}
		if err := ShowServer.UpdatePersonalInformation(ps); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "更新失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, ps)
		}
	})

	//对接查看
	r.GET("/showall", func(c *gin.Context) {
		if psall, err := ShowServer.GetShow(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "SHOW失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, psall)
		}
	})

	//查看某个人的SHOW
	r.GET("/show/:name", func(c *gin.Context) {
		name := c.Param("name") //http://localhost:8082/show/a1

		if ps, err := ShowServer.GetOneShow(name); err != nil { //通过
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取信息失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, ps)
		}
	})

	//删除
	//真删除
	r.POST("/deletetrue/:id", func(c *gin.Context) {
		id := c.Param("id")
		idd, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("转换失败", err)
		}
		if err := ShowServer.DeleteTrue(int64(idd), 1); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "真！！！删除失败！",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": "真！！！删除成功",
			})
		}
	})

	//假删除
	r.POST("/deletefalse/:id", func(c *gin.Context) {
		id := c.Param("id")
		idd, err := strconv.Atoi(id)
		fmt.Println(idd)
		if err != nil {
			fmt.Println("转换失败", err)
		}
		if err := ShowServer.DeleteFalse(int64(idd), 1); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "假！！！删除失败！",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": "假！！！删除成功",
			})
		}
	})

	if err := r.Run(":8082"); err != nil {
		log.Fatal(err)
	}

}
