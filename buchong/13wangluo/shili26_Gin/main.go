// go语言web框架（Gin）详解（有点了解，有点懵）
// www.kancloud.cn/imdszxs/golang/1509697
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//在go语言开发的web框架中，
//有两款著名web框架分别是Martini和Gin
//两款web框架相比较的话，Gin自己说它比Martini要强很多
//Gin是go语言写的一个web框架
//它具有运行速度快，分组的路由器
//良好的崩溃捕获和错误处理，非常好的支持中间件和json
//总之在go语言开发领域是一款值得好好研究的web框架
//开源地址：github.com/gin-gonic/gin
//首先下载安装gin包
//go get -u github.com/gin-gonic/gin
//一个简单的例子：/

/*
func main() {
	// default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// 输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.Run()//默认是0.0.0.0:8080？？
	// listen and serve on 0.0.0.0:8080
}

*/

// 编译运行程序，打开浏览器，访问localhost:8080/ping页面显示：{"message":"ping"}
//gin的功能不只是简单输出json数据
//它是一个轻量级的web框架
//支持RestFull风格API，支持GET，POST，PUT，PATCH，DELETE，OPTIONS等http方法
//支持文件上传，分组路由，Multipart/Urlencoded FORM
//以及支持JsonP，参数处理等等功能
//这些都和WEB紧密相关，通过提供这些功能，使开发人员更方便地处理WEB业务

//Gin实际应用
//接下来使用Gin作为框架来搭建一个拥有静态资源站点
//动态web站点，以及RESTFull API接口站点（可专门作为手机APP应用提供服务使用）组成的
//亦可根据情况分拆这套系统
//每种功能独立出来单独提供服务。

//下面按照一套系统但采用分站点来说明
//首先是整个系统的目录结构，website目录下面static是资源类文件
//为静态资源站点专用，photo目录是UGC上传图片目录
//tpl是动态站点的模板
//
//当然这个目录结构是一种约定
//可以根据情况来修改，
//整个项目已经开源，可以访问来详细了解：github.com/ffhelicopter/tmm
//具体每个站点的功能怎么实现呢？
//请看下面有关每个功能的讲述：

//1. 静态资源站点：
//一般网页开发中，我们会考虑把js，css，以及资源图片放在一起，作为静态站点部署在CDN，
//提升响应速度
//采用Gin实现起来非常简单
//当然也可以使用net/http包轻松实现
//但使用Gin会更方便
//不管怎么样，使用Go开发，我们可以不用花太多时间在web服务环境搭建上，
//程序启动就直接可以提供web服务了

/*
func main() {
	router := gin.Default() // 静态资源加载，本例为css，js以及资源图片
	router.StaticFS("/public", http.Dir("E:/Geek/src/learn.go/buchong/13wangluo/shili26_Gin"))
	router.StaticFile("/favicon.ico", "E:/Geek/src/learn.go/buchong/13wangluo/shili26_Gin/pic/2.png")
	// Listen and server 0.0.0.0:80
	router.Run(":80")
}

*/

//首先需要是生成一个Engine，这是gin的核心，
//默认带有Logger和Recovery两个中间件
//router := gin.Default()
//StaticFile是加载单个文件，而StaticFS是加载一个完整的目录资源：
//func (group *RouterGroup)StaticFile(relativePath, filepath string)IRoutes
//func (group *RouterGroup)StaticFS(relativePath,string, fs http.FileSystem)IRoutes
//这些目录下资源是可以随时更新，而不用重新启动程序
//现在编译运行程序，静态站点就可以正常访问了
//访问： localhost/public/images/logo.jpg
//图片加载正常
//每次请求响应都会在服务端有日志产生，包括响应时间，加载资源名称，响应状态值等

//2. 动态站点
//如果需要动态交互的功能，比如发一段文字+图片上传
//由于这些功能除了前端页面外，还需要服务端程序一起来实现
//而且迭代需要经常需要修改代码和模板
//所以把这些统一放在一个大目录下，姑且称动态站点
//tpl是动态站点所有模板的根目录，这些模板可调用静态资源站点的css，图片等
//photo是图片上传后存放的目录

func main() {
	router := gin.Default()

	// 导入所有模板，多级目录结构需要这样写
	router.LoadHTMLGlob("website/tpl/*/*") // 啥意思？
	//website分组
	v := router.Group("/")
	{
		v.GET("/index.html", handler.IndexHandler) //这些是模板？
		v.GET("/add.html", handler.AddHandler)
		v.POST("/postme.html", handler.PostmeHandler)
	}
	// router.Run(":80") 这样写就可以了，
	//下面所有代码（go1.8+）是为了优雅处理重启等动作
	srv := &http.Server{
		Addr:         ":80",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	go func() {
		// 监听请求
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 优雅Shutdown (或重启)服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	//syscall.SIGKILL
	<-quit

	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	select {
	case <-ctx.Done():
	}
	log.Println("Server exiting")
}

//在动态站点实现中，引入Web分组以及优雅重启这两个功能
//web分组功能可以通过不同的入口根路径来区别不同的模块
//这里我们可以访问localhost/index.html
//如果新增一个分组，比如：
//v := router.Group("/login")
//我们可以访问：http://localhost/login/xxxx
//xxxx是我们v.GET方法或v.POST方法中的路径
//导入所有模板，多级目录结构需要这样写
//router.LoadHTMLGlob("website/tpl/*/*")
//website分组
//v := router.Group("/")
//{
//v.GET("/index.html",handler.IndexHandler)
//v.GET("/add.html", handler.AddHandler)
//v.POST("/postme.html", hanlder.PostmeHandler)
//}
//通过router.LoadHTMLGlob("website/tpl/*/*")导入模板根目录下所有文件
//在前面讲过html/template包的使用，这里模板文件中的语法和前面的一致
//router.LoadHTMLGlob("website/tpl/*/*")
//比如v.GET("/index.html",handler.IndexHandler)
//通过访问 localhost/index.html
//这个URL，实际由handler.IndexHandler来处理
//而在tmm目录下的handler存放了package handler文件
//在包里定义了IndexHandler函数
//它使用了index.html模板
//func IndexHandler(c *gin.Context){
//	c.HTML(http.StatusOK, "index.html", gin.H{
//		"Title":"作品欣赏",
//	})
//}
//index.html模板：
//<!DOCTYPE html><html><head>{{template "header" .}}</head><body><!--导航--><div class="feeds">    <div class="top-nav">        <a href="/index.tml" class="active">欣赏</a>        <a href="/add.html" class="add-btn">            <svg class="icon" aria-hidden="true">                <use  xlink:href="#icon-add"></use>            </svg>            发布        </a>    </div>    <input type="hidden" id="showmore" value="{$showmore}">    <input type="hidden" id="page" value="{$page}">    <!--</div>--></div><script type="text/javascript">    var done = true;    $(window).scroll(function(){        var scrollTop = $(window).scrollTop();        var scrollHeight = $(document).height();        var windowHeight = $(window).height();        var showmore = $("#showmore").val();        if(scrollTop + windowHeight + 300 >= scrollHeight && showmore == 1 && done){            var page = $("#page").val();            done = false;            $.get("{:U('Product/listsAjax')}", { page : page }, function(json) {                if (json.rs != "") {                    $(".feeds").append(json.rs);                    $("#showmore").val(json.showmore);                    $("#page").val(json.page);                    done = true;                }            },'json');        }    });</script>    <script src="//at.alicdn.com/t/font_ttszo9rnm0wwmi.js"></script></body></html>
//在index.html模板中，通过{{template "header"}}语句
//嵌套了header.html 模板
//header.html模板：
//{{ define "header" }}    <meta charset="UTF-8">       <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no, minimal-ui">    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">    <meta name="format-detection" content="telephone=no,email=no">    <title>{{ .Title }}</title>    <link rel="stylesheet" href="/public/css/common.css">    <script src="/public/lib/jquery-3.1.1.min.js"></script>    <script src="/public/lib/jquery.cookie.js"></script>    <link href="/public/css/font-awesome.css?v=4.4.0" rel="stylesheet">{{ end }}
//{{define "header"}}让我们在模板嵌套时直接使用header名字
//而在inde.html中的{{template "header" .}}
//注意“.”可以使参数嵌套传递，否则不能传递
//比如这里的Title
//现在访问localhost/index.html
//可以看到浏览器显示Title是作品欣赏
//这个title是通过IndexHandler来指定的
//接下来点击“发布”按钮，我们可以进入发布也秒你，上传图片，点击“完成”提交
//会提示我们成功上传图片
//可以在photo目录中看到刚才上传的图片

//优雅重启在迭代中由较好的实际意义，每次版本发布
//如果直接停服再部署重启，对业务还是有蛮大的影响，而通过优雅重启
//这方面的体验可以做得更好些
//这里ctrl+c 过5秒服务停止
//中间件的使用，在API中可能使用限流，身份验证等
//
//go语言中net/http设计的一大特点就是特别容易构建中间件
//gin也提供了类似的中间件
//要注意的是在gin里面中间件只对注册过的路由函数起作用
//而对于分组路由，嵌套使用中间件，可以限定中间件的作用范围
//大致分为全局中间件，单个中间件和分组中间件
//即使是全局中间件，其使用前的代码不受影响
//也可在handler中局部使用，具体见api.GetUesr
//在高并发场景中，有时候需要用到限流降速的功能
//这里引入一个限流中间件
//有关限流方法常见有两种，具体可自行研究，这里只讲使用：
//导入import “github.com/didip/toollbooth/limiter”包
//在上面代码基础上增加如下语句：
//rate-limit限流中间件
//Imt := tollbooth.NewLimiter(1,nil)
//Imt.SetMessage("服务繁忙，请稍后再试。。。")
//并修改
//v.GET("/index.html",LimitHandler(Imt), handler.IndexHandler)
//当f5刷新localhost/index.html页面时，浏览器会显示：服务繁忙，请稍后再试。。。
//限流策略也可以为IP：
//tollbooth.LimitByKeys(Imt, []string{"127.0.0.1","/"})
//更多限流策略的配置，可以进一步github.com/didip/toolbooth/limiter了解

//RestFullAPI接口
//前面说了在gin里面可以采用分组来组织访问URL，这里RestFullAPI需要给出不同的访问URL来和动态站点区分，
//所以新建了一个分组v1
//在浏览器中访问localhost/v1/user/1100000/
//这里对v1.GET("/user/:id/*action", LimitHandler(Imt),api.GetUser)进行了限流控制
//所以如果频繁访问上面地址也将会有限流
//这在API接口中非常有作用
//通过api这个包，来实现所有有关API的代码
//在GetUser函数中，通过读取mysql数据库
//查找到对应userid的用户信息，并通过Json格式反馈给client
//在api.GetUser中，设置了一个局部中间件：
//CORS局部CORS，可在路由中设置全局的CORS
//c.Writer.Header().Add("Access-Control-Allow-Origin","*")
//gin关于参数的处理，api包中api.go文件中有简单说明，这里不展开
//这个项目的详细情况，github.com/ffhelicopter/tmm （作者的）
//有关gin的更多信息，访问github.com/gin-gonic/gin，该开源项目比较活跃，可以关注。
