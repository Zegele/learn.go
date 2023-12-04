//go语言开发一个简单的相册网站
//www.kancloud.cn/imdszxs/golang/1509683

/*
package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

//开发一个简单但五脏俱全的相册网站

//首先创建一个用于存放工程源代码的目录:
//photoweb
// 	---uploads
//	---photoweb.go

//基本功能
//1. 支持图片上传
//2. 在网页中可以查看已上传的图片
//3. 能看到所有上传的图片列表
//4. 可以删除指定的图片
//
//使用net/http包提供网络服务
//1. 上传图片：

const (
	UPLOAD_DIR = "./uploads"
)
*/

//func uploadHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "Get" {
//		io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
//			"enctype=\"multipart/form-data\">"+
//			"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
//			"<input type=\"submit\" value=\"Upload\" />"+
//			"</form>")
//		return
//	}
//	if r.Method == "POST" {
//		f, h, err := r.FormFile("image")
//		if err != nil {
//			http.Error(w, err.Error(),
//				http.StatusInternalServerError)
//			return
//		}
//		filename := h.Filename
//		defer f.Close()
//
//		t, err := os.Create(UPLOAD_DIR + "/" + filename)
//		if err != nil {
//			http.Error(w, err.Error(),
//				http.StatusInternalServerError)
//			return
//		}
//		defer t.Close()
//
//		if _, err := io.Copy(t, f); err != nil {
//			http.Error(w, err.Error(),
//				http.StatusInternalServerError)
//			return
//		}
//		http.Redirect(w, r, "/view?id="+filename,
//			http.StatusFound)
//	}
//}

//func main() {
//	http.HandleFunc("/", listHandler)
//	http.HandleFunc("/view", viewHandler) // 注册viewHandler方法
//	http.HandleFunc("/upload", uploadHandler)
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err.Error())
//	}
//}

//可以看到，集合main() uploadHandler()方法，
//针对HTTP GET方式请求/upload路径，
//程序将会往http.ResponseWriter类型的实例对象w中写入一段HTML文本
//即输出一个HTML上传表单
//如果我们使用浏览器访问这个地址，那么网页上将会是一个可以上传文件的表单。
//只有上传表单还不能完成图片上传
//服务端程序还必须有接受上传图片的相关处理
//针对上传表单提交过来的文件，
//我们对uploadHandler()方法再添加些业务逻辑程序：

//对uploadHandler()增加后
//如果是客户端发起的HTTP POST，那么首先从表单提交过来的字段寻找名为image的文件域并对其接值
//调用r.FormFile()方法会返回3个值
//各个值的类型分别是multipart.File, *multipart.FileHeader 和 error
//如果上传的图片接收不成功，那么在示例程序中返回一个HTTP服务端的内部错误给客户端
//如果上传的图片接收成功，则该图片的内容复制到一个临时文件里
//如果临时文件创建失败，或者图片副本保存失败，
//都将触发服务端内部错误
//如果临时文件创建成功兵器图片副本保存成功， 即表示图片上传成功
//就跳转到查看图片页面
//此外，我们还定义了两个defer语句
//无论图片上传成功还是失败，
//当uploadHandler() 方法执行结束时，
//都会先关闭临时文件句柄，继而关闭图片上传到服务器文件流的句柄
//别忘了上传成功后，我们即可在网页上查看这张图片
//顺便确认图片是否真正上传到了服务端

//2.在网页上显示图片
//要在网页上显示图片，必须有一个可以访问到该图片的网址
//在前面的示例代码中，图片上传成功后会跳转到/view?id=这样的网址
//因此我们的程序要能够将对/view路径的访问映射到某个具体的业务逻辑处理方法
//首先，在photoweb程序中新增一个名为viewHandler()的方法，如下：
/*
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

*/

//上述代码中，我们首先从客户端请求中对参数进行接值
//r.FormValue("id")即可得到客户端请求传递的图片唯一ID
//然后我们将图片ID结合之前保存图片用的目录进行组装
//即可得到文件在服务器上的存放路径
//接着，调用http.ServeFile()方法将该路径下的文件从磁盘中读取并作为服务端的返回信息输出给客户端
//同时，也将http响应头输出格式预设为image类型
//这是一个比较简单的示意写法
//实际上应该严谨些，准确解析出文件的Mime Type将其作为Content-Type进行输出
//具体可参考go语言标准库的http.DetectContentType()方法和mime包提供的相关方法
//完成viewHandler()的业务逻辑后，我们将该方法注册到程序的main()方法中，
//与/view路径访问形成映射关联
//这样当客户端（浏览器）访问/view路径并传递id参数时，即可直接以HTTP形式看到图片的内容
//在网页上，将会呈现一张可视化的图片

//3.处理不存在的图片访问
//理论上，只要是uploads/目录下有的图片，都能够访问到
//但我们还是假设有意外情况，如：网址中传入的图片ID在uploads/没有对应的文件
//这是，我们的viewHandler()方法就显得很脆弱了
//不管是给出友好的错误提示还是返回404页面，
//都应该对这种情况作相应处理
//我们不妨先以最简单有效的方式对其进行处理
//修改viewHandler(),具体如下：
/*
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err) //这是干啥的
}

 */

// 增加了isExiists辅助函数，用于检查文件是否真的存在

// 4. 列出所有已上传图片
//应该有个入口，可以看到所有已上传的图片
//对于多有列出的这些图片，我们可以选择进行查看或删除等操作
//下面假设在访问首页时列出所有上传的图片
//由于我们将客户端上传的图片全部保存在工程的./uploads目录下
//所以程序中应该有个名叫listHandler()的方法，
//用于在网页上列出该目录下存放的所有文件
//暂时我们不考虑缩略图的形式列出所有已上传图片，只需列出可供访问的文件名即可
//实现listHandler()
//func listHandler(w http.ResponseWriter, r *http.Request) {
//	fileInfoArr, err := ioutil.ReadDir("./uploads")
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	var listHtml string
//	for _, fileInfo := range fileInfoArr {
//		imgid := fileInfo.Name
//		listHtml += "<li><a href=\"/view?id=" + imgid + "\">imgid</a></li>"
//	}
//	io.WriteString(w, "<ol>"+listHtml+"</ol>")
//}

//从上面的listHandler()方法中可以看到
//程序先从./uploads目录中遍历得到所有文件并赋值到fileInfoArr变量里
//fileInfoArr是一个数组，其中的每一个元素都是一个文件对象
//然后，程序遍历fileInfoArr数组并从中得到图片的名称
//用于在后续的HTML片段中显示文件名和传入的参数内容
//listHtml变量用于在for循环中将图片名称一一串联起来生成一段HTML
//最后调用io.WriteString()方法将这段HTML输出返回给客户端
//然后在phtotweb.go程序的main()方法中，我们将对首页的访问映射到listHandler()方法

//这样在访问网站首页的时候，即可看到已上传的所有图片列表
//是否注意到：在phtotweb.go程序的uploadHandler()和listHandler()
//都使用io.WriteString()方法输出HTML
//在业务逻辑处理程序中混杂HTML可不是什么好事
//代码多起来后会导致程序不够清晰
//而且改动程序里面的HTML文本时，
//每次都要重新编译整个工程的源代码才能看到修改后的效果。
//正确的做法是，应该将业务逻辑程序和表现层分离开来，各自单独处理，
//这时候，就需要使用网页模板技术了
//Go标准库中的html/template包对网页模板有着良好的支持
//接下来，让我们来了解如何在photoweb.go程序中用上了Go的模板功能

//渲染网页模板
//使用go标准库提供的html/template包
//可以让我们将HTML从业务逻辑程序中抽离出来形成独立的模板文件
//这样业务逻辑程序只负责处理业务逻辑部分和提供模板需要的数据
//模板文件负责数据表现的具体形式
//
//然后模板解析器将这些数据以定义好的模板规则结合模板文件进行渲染
//最终将渲染后的结果一并输出，构成一个完整的网页
//下面我们把photoweb.go程序的uploadHandler()和listHandler()方法中的HTML文本抽出来，生成模板文件
//新建一个名为upload.html的文件。
//新建一个名为list.html的文件。???新建到哪里？
//上述模板中
//双大括号{{}}是区分模板代码和HTML的分隔符
//括号里面可以是要显示输出的数据，或者是控制语句
//比如if判断式或者range循环体等
//range语句在模板中式一个循环过程体，紧跟在range后面的必须是一个array，slice或map类型的变量
//在list.html模板中，images是一组string类型的切片
//在使用range语句遍历的过程中，即表示该循环体中的当前元素
//.|formatter表示对当前这个元素的值以formatter方式进行格式化输出
//比如.|urlquery}即表示对当前元素的值进行转换以适合作为URL一部分
//而{{.|html表示对当前元素的值进行适合用于HTML显示的字符转化，
//比如：">"会转译成">"   //???? 没看懂。。。
//如果range关键字后面紧跟着的是map这样的多维复合结构体，
//循环体中的当前元素可以用.key1.key2.keyN这样的形式表示
//如果要更改模板中默认的分隔符，可以使用template包提供的Delims()方法
//在了解模板语法后
//接着我们修改photoweb.go源文件，
//引入html/template包，
//并修改uploadHandler()和listHandler()方法
//具体如下：
//func uploadHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "GET" {
//		t, err := template.ParseFiles("upload.html")
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		t.Execute(w, nil)
//		return
//	}
//	if r.Method == "POST" {
//		f, h, err := r.FormFile("image")
//		if err != nil {
//			http.Error(w, err.Error(),
//				http.StatusInternalServerError)
//			return
//		}
//		filename := h.Filename
//		defer f.Close()
//
//		t, err := os.Create(UPLOAD_DIR + "/" + filename)
//		if err != nil {
//			http.Error(w, err.Error(),
//				http.StatusInternalServerError)
//			return
//		}
//		defer t.Close()
//
//		if _, err := io.Copy(t, f); err != nil {
//			http.Error(w, err.Error(),
//				http.StatusInternalServerError)
//			return
//		}
//		http.Redirect(w, r, "/view?id="+filename,
//			http.StatusFound)
//	}
//}

//func listHandler(w http.ResponseWriter, r *http.Request) {
//	fileInfoArr, err := ioutil.ReadDir("./uploads")
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	locals := make(map[string]interface{})
//	images := []string{}
//	for _, fileInfo := range fileInfoArr {
//		images = append(images, fileInfo.Name())
//	}
//	locals["images"] = images
//	t, err := template.ParseFiles("list.html")
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	t.Execute(w, locals)
//}

//上面的代码中，template.ParseFiles()函数将会
//读取指定模板的内容并且返回一个*template.Template值
//t.Execute()方法会根据模板语法来执行模板的渲染，
//并将渲染后的结果作为HTTP的返回数据输出

//在uploadHandler()方法和listHandler()方法中
//均调用了template.ParseFiles()和t.Execute()这两个方法
//根据DRY（Don't Repeat Yourself）原则
//我们可以将模板渲染代码分离出来，单独写一个处理函数
//以便其他业务逻辑函数都可以使用，
//于是，我们可以定义一个名为renderHtml()的方法用来渲染模板：
//func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
//	t, err := template.ParseFiles(tmpl + ".html")
//	if err != nil {
//		return
//	}
//	err = t.Execute(w, locals)
//	return err
//}

// 有了renderHtml()这个通用的模板渲染方法，
//uploadHandler() , listHandler()方法的代码可以再精简下：
/*
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		//交给渲染
		if err := renderHtml(w, "upload", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename,
			http.StatusFound)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images

	//交给渲染
	if err = renderHtml(w, "list", locals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

 */

// 当我们引入go标准库中的html/template包，实现了业务逻辑层与表现层分离后
//对模板渲染逻辑去重，编写并使用通用模板渲染方法renderHtml()
//这让业务逻辑处理层的代码看起来确实要清晰简洁许多
//不过，直觉敏锐的你可能已经发现
//无论是重构后的uploadHandler()还是listHandler()方法
//每次调用这两个方法时都会重新读取并渲染模板
//很明显，这很低效，也比较浪费资源，有没有一种办法可以让模板只加载一次呢？

//模板缓存
//对模板进行缓存，即指一次性预加载模板，
//我们可以在photoweb程序初始化运行的时候
//将所有模板一次性加载到程序中
//正好Go的包加载机制允许我们在init()函数中做这样的事情
//init()会在main()函数之前执行
//首先，我们在photoweb程序中声明并初始化一个全局变量templates
//用于存放所有模板内容
//templates := make(map[string]*template.Template)
//templates是一个map类型的复合解构，map的键（key）是字符串类型，
//即模板的名字，值（value）是*template.Template类型
//接着，我们在photoweb程序的init()函数中一次性加载所有模板：
//func init() {
//	templates := make(map[string]*template.Template)
//	for _, tmpl := range []string{"upload", "list"} {
//		t := template.Must(template.ParseFiles(tmpl + ".html"))
//		templates[tmpl] = t // templates是什么？？？
//	}
//}

//在上面的代码中，在template.ParseFiles()方法的外层强制使用template.Must()进行封装
//template.Must()确保了模板不能解析成功时，一定会触发处理流程
//之所以这么做，时因为倘若模板不能成功加载，程序能做的唯一有意义的事情就是退出
//
//在range语句中，包含了我们希望加载的upload.html和list.html两个模板
//如果我们想加载更多模板
//只需要往这个数组中添加更多元素即可
//当然，最好的办法应该是将所有HTML模板文件统一放到一个子文件夹中，
//然后对这个模板问及爱你加进行遍历和预加载
//如果需要加载新的模板，只需要这个文件夹中新建模板即可
//这样做的好处是不用反复修改代码即可重新编译程序
//而且实现了业务层和表现层真正意义上的分离

//试试
//首先创建一个名为./views的目录，
//然后将当前目录下所有html文件移动到该目录下
//接着适当的对init()方法中的代码进行改写
//好让程序初始化时即可预加载该目录下的所有模板文件
//如下：
/*
const (
	TEMPLATE_DIR = "./views"
)

func init() {
	templates := make(map[string]*template.Template)
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templatePath); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templatePath] = t
	}
}

// 同时别忘了对renderHtml()的代码进行相应的调整：
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
	templates := make(map[string]*template.Template)
	err = templates[tmpl].Execute(w, locals)
	return err
} //这个不完整吧


 */

//此时，renderHtml()函数的代码也变得更为简洁
//还好我们之前单独封装了renderHtml()函数
//这样全局代码中值更改这一个地方，这无疑是代码解耦的好处之一

//错误处理
//在前面的代码中，有不少地方对于出错处理都是直接返回http.Error() 50x系列的服务端内部错误
//从DRY的原则来看，不应该在程序中导出使用一样的代码
//我们可以定义一个名为check()的方法
//用于统一捕获50x系列的服务端内部错误
/*
func check(err error) {
	if err != nil {
		panic(err)
	}
}

 */

// 此时，我们可以将photoweb程序中出现的“报错”代码替换为check()处理
//check(err)
//错误处理虽然简单很多，但是也带来一个问题，
//由于发送错误触发处理流程必然会引发程序停止运行
//这种改发有点像搬起石头砸自己的脚
//其实我们可以换一种思维方式
//尽管我们从书写上能保证大多数错误都能得到相应的处理
//但根据墨菲定律，有可能出问题的地方就一定会出问题
//在计算机程序里尤其如此
//如果程序中我们正确处理99个错误
//但若有一个系统错误意外导致程序出现异常
//那么程序同样还是会终止运行
//我们不能预计一个工程里边会出现多少意外的情况
//但是不管什么意外
//只要会触发错误流程，我们就有办法对其进行处理
//如果这样思考
//那么前面这种改法又何尝不是置死地而后生呢？
//接下来，让我们了解如何处理panic导致程序崩溃的情况

//巧用闭包避免程序运行时出错崩溃
//go支持闭包
//闭包可以是一个函数里边返回的另一个匿名函数
//该匿名函数包含了定义在它外面的值
//使用闭包，可以让我们网址的业务逻辑处理程序更安全地运行
//我们可以在photoweb程序中针对所有的业务逻辑处理函数（listHandler(),viewHandler(),uploadHandler()） 再进行一次包装
//再如下的代码中，我们定义了一个名为safeHandler()的函数，
//该函数有哦一个参数并返回一个值
//传入的参数和返回值都是一个函数
//且都是http.HandlerFunc类型
//这种类型的函数有两个参数：http.ResponseWriter和*http.Request
//函数规格同photoweb的业务逻辑处理函数完全一致
//事实上，我们正是要吧业务逻辑处理函数作为参数传入到safeHandler()方法中
//这样任何一个错误流程向上回溯的时候，
//我们都能对其进行拦截处理，从而也能避免程序停止运行:
/*
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				// 或者输出自定义的50x错误页面
				//w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e)
				// logging
				log.Printf("WARN: panic in %v - %v", fn, err)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

 */

//在上述这段代码中，我们巧妙使用了defer关键字搭配recover()方法终结panic的
//safeHandler()接收一个业务逻辑处理函数作为参数，
//同时调用这个业务逻辑处理函数
//该业务逻辑处理函数里面引发了panic，
//则调用recover()对其进行检测
//若为一般性的错误，则输出HTTP 50x，出错信息并记录记录日志
//而程序将继续良好运行
//要应用safeHandler()函数，只需要在main()中对各个业务逻辑处理函数做一次包装
//如下：
//func main() {
//	http.HandleFunc("/", safeHandler(listHandler))
//	http.HandleFunc("/view", safeHandler(viewHandler))
//	http.HandleFunc("/upload", safeHandler(uploadHandler))
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		log.Fatal("ListenAndServe:", err.Error())
//	}
//}

// 动态请求和静态资源分离
//还有一个疑问，前面的业务逻辑层都是动态请求
//但若是针对静态资源（比如CSS和JavaScript等）
//是没有业务逻辑处理的
//只需提供静态输出
//在go里，当然是可行的
//还记得在viewHandler()函数里，用到http.ServeFile()这个方法吗？
//net/http包提供的这个ServeFile()函数可以将服务端的一个文件内容读写到http.Response-Writer并返回给请求来源的*http.Request客户端
//用前面介绍的闭包技巧结合这个http.ServeFile()方法
//我们就能轻而易举地实现业务逻辑的动态请求和静态资源的完全分离
//假设我们有./public目录下相应的文件：
//[GET]/assets/css/*.css
//[GET]/assets/js/*.js
//[GET]/assets/images/*.js
//然后，我们定义一个名为staticDirHandler()方法，用于实现上述需求：
/*
const (
	ListDir = 0x0001
)

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	mux.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}

 */

// 如此，即完美实现了静态资源和动态请求的分离
//当然，我们要思考是否确实需要用go来提供静态资源的访问
//如果使用外部web服务器（比如Nginx等）
//就没必要使用Go编写的静态文件服务了
//在本机做开发时有一个程序内置的静态文件服务器还是很实用的

//重构（见photowebmain.go ）


//更多资源
//Go的第三方库很丰富，无论是对于关系型数据库驱动还是非关系型的键值存储系统的接入
//都有着良好地支持，
//而且还有丰富的Go语言Web开发框架以及用于Web开发的相关工具包
//可以访问http://godashboard.appspot.com/project
//了解更多第三方库的详细信息
