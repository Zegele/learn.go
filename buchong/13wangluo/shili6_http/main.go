//go语言HTTP客户端实现简述
//www.kancloud.cn/imdszxs/golang/1509677
package main

import "net/http"

// HTTP（HyperText Transfer Protocol,超文本传输协议）是互联网上应用最广泛的一种网络协议，
//定义了客户端和服务端之间请求与相应的传输标准
//go标准库提供了net/http包，涵盖了HTTP客户端和服务端的具体实现
//使用net/http包，可以方便编写HTTP客户端或服务端的程序

//基本方法：
//net/http包提供了间接的http客户端实现，可以直接使用最常见的GET和POST
//可以通过net/http包里面的Client类提供的如下方法发起HTTP请求：
//func(c *Client)Get(url string)(r *Response, err error)
//func(c *Client)Post(url string, bodyType string, body io.Reader)(r *Response, err error)
//func(c *Client)PostForm(url string, data url.Values)(r *Response, err error)
//fumc(c *Client)Head(url string)(r *Response, err error)
//func(c *Client)Do(req *Resquest)(resq *Response, err error)
//
//1. http.Get()
//要请求一个资源，只需要调用http.Get()方法（等价于http.DefaultClient.Get()）即可，示例：
/*
func main() {
	resp, err := http.Get("http://c.biancheng.net")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

*/

// 底层调用
//其实通过http.Get发起请求时，默认调用的上述http.Client缺省对象上的Get方法：
//func Get(url string)(resp *Response, err error){
//		return DefaultClient.Get(url)
//}
//而DefaultClient默认指向的正式http.Client的示例对象：
//var DefaultClient = &Client{}
//它是net/http包公开属性，当我们在http上调用Get，Post，PostFrom，Head方法时，最终调用的都是该对象上的对应方法

//返回值
//http.Get()方法的返回值有两个，分别是一个相应对象和一个error对象，
//如果请求过程中出现错误，则error对象不为空
//否则可以通过相应对象获取状态码，响应头，相应实体等信息，
//响应对象所属的类时http.Response
//可以通过查看API文档或者源码了解该类型的具体信息，
//一般我们可以
//通过resp.Body获取响应实体，
//通过resp.Header获取响应头，
//通过resp.StatusCode获取响应状态码
//获取响应成功后得调用resp.Body上的Close方法结束网络请求释放资源

//2.http.Post()
//要以POST的方式发送数据，也很简单，只需调用http.Post()方法并依次传递下面的3个参数即可：
//1. 请求的目标URL
//2. 将要post数据的资源类型（MIMEType）
//3. 数据的比特流（[]byte形式）
//示例：如何上传一张图片
/*
func main() {

	resp, err := http.Post("http://c.biancheng.net/upload", "image/jpeg", &buf)
	//其中&buf为图片的资源
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

*/

//3. http.PostFrom()
//http.PostFrom()方法实现了标准编码格式为”application/x-www-from-urlencoded“的表单提交
//下面模拟了html表单向后台提交信息的过程：
/*
func main() {
	resp, err := http.PostForm("http://www.baidu.com", url.Values{"wd": {"golang"}})
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

*/

// 注意： POST请求参数需要通过url.Values方法进行编码和封装

// 4. http.Head()
//HTTP的Head请求表示只请求目标URL的响应头信息 ，不返回响应实体
//可以通过net/http包的http.Head()方法发起Head请求，
//该方法和http.Get()方法一样
//只需要传入目标URL参数即可

/*
func main() {
	resp, err := http.Head("http://c.biancheng.net")
	if err != nil {
		fmt.Println("Request Failed:", err.Error())
		return
	}
	defer resp.Body.Close()

	//打印头信息
	for key, value := range resp.Header {
		fmt.Println(key, ":", value)
	}
}

*/

//5. (*http.Client).Do()
//在多数情况下，http.Get(), http.Post()和http.PostForm()就可以满足需求
//但是如果我们发起的HTTP请求需要更多的自定义请求信息，比如：
//1.设置自定义User-Agent, 而不是默认的Go http package；
//2.传递Cookie信息
//3.发起其他方式的HTTP请求，比如PUT，PATCH，DELETE等
//此时可以通过http.Client类提供的Do() 方法来实现，使用该方法时，就不再是通过缺省的DefaultClient对象调用http.Client 类中的方法了
//而是需要我们手动实例化Client对象并传入添加了自定义请求头信息的请求对象来发起HTTP请求：
/*
func main() {
	// 初始化客户端请求对象
	req, err := http.NewRequest("Get", "http://c.biancheng.net", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 添加自定义请求头
	req.Header.Add("Custom-Header", "Custom-Value")
	//其他请求头配置
	client := &http.Client{
		//设置客户端属性
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

*/

//用于初始化请求对象的http.NewRequest方法需要传入3个参数
//第一个是请求方法
//第二个是目标URL
//第三个是请求实体
//只有POST，PUT，DELETE之类的请求才需要设置请求实体
//对于HEAD，GET而言，传入nil即可
//http.NewRequest 方法返回的第一个值就是请求对象示例req，该实例所属的类是http.Request
//可以调用该类上的公开方法和属性对请求对象进行自定义配置
//比如请求方法，URL，请求头等
//设置完成后，就可以将请求对象传入client.Do()方法发起HTTP请求
//之后的操作和前面四个基本一样
//http.Post, http.PostForm, http.Head, http.NewRequest
//方法的底层实现及返回值和http.Get方法一样

//高级封装
//除了基本的HTTP操作，go标准库也暴露了比较底层的HTTP相关库，
//让开发者可以基于这些库灵活定制http服务器和使用http服务

//1.自定义http.Client
//前面我们使用的http.Get(), http.Post(), http.PostForm(), http.Head()
//其实都是在http.DefaultClient的基础上进行调用的
//比如http.Get()等价于http.Default-Client.Get()，以此类推
//既然存在默认的Client，那么就可以自定义
//在net/http包中，的确提供了Client类型
//看看http.Client 类型的结构
//type Client struct{
//	// Transport用于确定HTTP请求的创建机制。如果为空，将会使用DefaultTransport
//	Transport RoundTripper
//
//	//CheckRedirect定义重定向策略，如果CheckRedirect不为空，客户端将在跟踪HTTP重定向前调用该函数
//	// 两个参数req,和via分别为即将发起的请求和已经发起的所有请求，最早的,已发起请求在最前面
//	// 如果CheckRedirect返回错误，客户端将直接返回错误，不会再发起该请求
//	// 如果CheckRedirect为空，Client将采用一种确认策略，将在10个连续请求后终止	、
//	CheckRedirect func(req *Request, via []*Request)error
//
//	// 如果Jar为空，Cookie将不会在请求中发送，并会在响应中被忽略
//	Jar CookieJar
//}
//http.Client类型包含了3个公开数据成员：
//Transport RoundTripper
//CheckRedirect func(req *Request, via []*Request)error
//Jar CookieJar
//其中Transport类型必须实现http.RoundTripper接口
//Transport指定了执行一个HTTP请求的运行机制
//倘若不指定具体的Transport，默认会使用http.DefaultTransport，
//这意味着http.Transport也是可以自定义的
//net/http包中的http.Transport类型实现了http.RoundTripper接口
//CheckRedirect函数指定处理重定向的策略
//当使用HTTP Client的Get()或者是Head()方法发送HTTP请求时，若响应返回的状态码为30x（比如301/302/303/307）
//HTTP Client会在遵循跳转规则之前先调用这个CheckRedirect函数
//Jar可用于在HTTP Client中设定Cookie，Jar的类型必须实现了http.CookieJar接口
//该接口预定义了SetCookies()和Cokkies() 两个方法
//如果HTTP Client中没有定义Jar,Cookie 将被忽略而不会发送到客户端
//实际上，我们一般都用http.SetCookie()方法来设定Cookie
//使用自定义的http.Client及其Do()方法
//我们可以非常灵活地控制HTTP请求，比如发送自定义HTTP Header
//或是改写重定向策略等
//创建自定义的HTTP Client非常简单
//具体代码如下：
 client := &http.Client{
	 CheckRedirect: redirectPolicyFunc,
}
 resp, err := client.Get("http://example.com") //...
 req, err := http.NewRequest("GET","http://example.com", nil)//...
 req.Header.Add("User-Agent","Our Custom User-Agent")
 req.Header.Add("If-None-Match",`W/"TheFileEtag"`)
 resp, err := client.Do(req) //...


 // 2.自定义http.Transport
 //http.Client类型的第一个数据成员，就是http.Transport对象
 //该对象指定执行一个HTTP请求时的运行规则
 //type Transport struct{
 //	//Proxy指定用于针对特定请求返回代理的函数，如果该函数返回一个非空的错误，请求将终止并返回该错误
 //	// 如果Proxy为空，或者返回一个空的URL指针，将不使用代理
 //	Proxy func(*Request)(*url.URL, error)
 //
 //	Dial指定用于创建TCP连接的dial()函数
 //	//如果Dial为空(不指定)，将默认使用net.Dial()函数
 //	Dial func(net, addr string)(c net.Conn, err error)
 //
 //	//TLSClientConfig指定用于tls.Client的TLS配置信息
 //	// 如果为空则使用默认配置
 //	TLSClientConfig *tls.Config // SSL连接专用
 //
 //	DisableKeepAlives bool // 是否取消长连接，默认值为false，即启用长连接
 //
 //	DisableCompression bool //是否取消压缩（GZip），默认值为false,即启用压缩
 //
 //	//如果MaxIdleConnsPerHost为非零值，
 //	//它用于控制每个host所需要保持的最大空闲连接(keep-alive)数量，
 //	如果该值为空（不指定），则使用DefaultMaxIdleConnsPerHost的常量值
 //	MaxIdleConnsPerHost int //...
 //	}


 //除了http.Transport类型中定义的公开数据成员以外，它同时还提供了几个公开的成员方法
 //	func (t *Transport)CloseIdleConnections() 该方法用于关闭所有非活跃的连接
 //	func (t *Transport)RegisterProtocol(scheme string, rt RoundTripper) 该方法可用于注册
 //	并启用一个新的传输协议，比如WebSocket的传输协议标准（ws），或者FTP，File协议等
 //	func (t *Transport)RoundTrip(req *Request)(resp *Response, err error) //用于实现http.RoundTripper接口



 //自定义 http.Transport也很简单，如下：
 /*
 tr := &http.Transport{
	 TLSClientConfig : &tls.Config{RootCAs : pool},
	 DisableCompression : true,
}
 client := &http.Client{
	 Transport : tr
}
 resp, err := client.Get("https://exampele.com")
  */

 //Client和Transport在执行多个goroutine的并发过程中都是安全的
 //但出于性能考虑，应当创建一次后反复使用

 //3. 灵活的http.RoundTripper接口
 //HTTP Client可以自定义，Client的第一个公开成员就是一个http.Transport类型的实例，
 //且该成员对应的类型必须实现http.RoundTripper接口
 //现在看看http.RoundTripper接口的具体定义：
 //type RoundTripper interface{
 //	//RoundTrip执行一个单一的HTTP事务，返回响应信息
 //	//RoundTrip函数的实现不应试图去理解响应的内容，
 //	//如果RoundTrip得到一个响应，无论该响应的HTTP状态码如何，都应将返回的err设置为nil。
 //	//非空的err只意味着没有成功获取到响应
 //	//类似地，RoundTrip也不应试图处理更高级别的协议，比如重定向，认证和Cookie等
 //	// RoundTrip不应修改请求内容，除非是为了理解Body内容
 //	//每一个请求的URL和Header域都应被正确初始化
 //	RoundTrip(*Request)(*Response,error)
 //	}、

 //	上述看到，http.RoundTripper接口很简单，只定义了一个名为RoundTrip的方法
 //	任何实现了RoundTrip()方法的类型即可实现http.RoundTripper接口
 //	前面我们看到的http.Transport类型正是实现了RoundTrip()方法继而实现了该接口
 //	RoundTrip()方法用于执行一个独立的HTTP事务，接受传入的\*Request请求
 //	作为参数并返回*Response响应值，以及一个error
 //	在实现具体的RoundTrip()方法时，不应该试图在该函数里解析HTTP响应信息
 //	若响应成功，error的值必须为nil，而与返回的HTTP状态码无关
 //	若不能成功得到服务端的响应，error必须为非零值
 //	类似地，也不应该试图在RoundTrip()中处理协议层面的相关细节
 //	比如，重定向，认证或是cookie等
 //	非必要情况下，不应该在RoundTrip()中改写传入的请求体（*Resquest）
 //	请求体的内容（比如URL和Header等）必须在传入RoundTrip()之前就已组织好并完成初始化
 //	通常，我们可以在默认的http.Transport之上包一层Transport并实现RoundTrip()方法，如下：

 type OurCustomTransport struct {
	 Transport http.RoundTripper
 }
 func (t *OurCustomTransport) transport()http.RoundTripper{
	 if t.Transport != nil{
		 return t.Transport
	 }
	 return http.DefaultTransport
 }

 func (t *OurCustomTransport) RoundTrip(req *http.Request)(*http.Response, error){
	 //处理一些事情 。。
	 //发起HTTP请求。。
	 //添加一些域到req.Header中
	 return t.transport().RoundTrip(req)
 }

 func(t *OurCustomTransport)Client()*http.Client{
	 return &http.Client{
		 Transport: t,
	 }
 }

func main() {
	t := &OurCustomTransport{
		//...
	}
	c := t.Client()
	resp, err := c.Get("http://example.com")

	// ...
}
 // 因为实现了http.RoundTripper接口的代码通常需要在多个goroutine中并发执行
 //因此我们必须确保实现代码的线程安全性


 //4. 设计优雅的HTTP Client
 //综上，go语言标准库提供的HTTP Client是相当优雅的
 //一方面提供了极其简单的使用方式，另一方面又具备极大的灵活性
 //Go语言标准库提供的HTTP Client被设计成上下两层结构
 //一层是上述提到的http.Client类及其封装的基础方法，不妨称为”业务层“
 //因为调用方通常只需要关心请求的业务逻辑本身，而无需关心非业务相关的技术细节，这些细节包括：
 //1. HTTP底层传输细节
 //2. HTTP代理
 //3. gzip压缩
 //4. 连接池及其管理
 //5. 认证（SSL或其他认证方式）
 //之所以HTTP Client可以做到这么好的封装性，是因为HTTP Client在底层抽象了
 //http.RoundTripper接口，而http.Transport实现了该接口，
 //从而能够处理更多细节，不妨称为”传输层“
 //HTTP Client在业务层初始化HTTP Mehtod，目标URL，请求参数，请求内容等主要信息后，
 //经过”传输层“，“传输层”在“业务层”处理的基础上补充其他细节，
 //然后再发起HTTP请求，接收服务端返回的HTTP响应
 //，