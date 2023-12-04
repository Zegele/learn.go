// go语言Cookie的设置与读取
// www.kancloud.cn/imdszxs/golang/1509689
package main

//Web开发中一个很重要的议题就是如何做好用户整个浏览过程的控制
//因为HTTP协议是无状态的，所以用户的每一次请求都是无状态的
//不知道在整个Web操作过程中那些连接与该用户有关
//应该如何来解决这个问题？
//Web里面经典的解决方案是Cookie和Session
//Cookie机制是一种客户端机制
//把用户数据保存在客户端
//而Session机制是一种服务器端的机制
//服务器使用一种类似于散列表的结构该保存信息
//每一个网站访客都会被分配给一个唯一的标识符
//即sessionID
//sessionID的存放形式无非两种：要么经过URL传递，要么保存在客户端的Cookie里
//当然，也可以将Session保存到数据库里，
//这样会更安全 ，但效率方面会有所下降

//本节主要介绍Go语言使用Cookie的方法

//设置Cookie
//Go语言中通过net/http包中的SetCookie来设置Cookie:
//http.SetCookie(w ResponseWriter, cookie *Cookie)
//w表示需要写入的response，cookie是一个struct
//让我们来看看对象是怎样的：
//type Cookie struct{
//	Name string
//	Value string
//	Path string
//	Domain string
//	Expires time.Time
//	RawExpires string
//	//MaxAge=0意味着没有指定Max-Age的值
//	// MaxAge<0意味着现在就删除Cookie，等价于Max-Age=0
//	// MaxAge>0意味着Max-Age属性存在并以秒为单位存在
//	MaxAge int
//	Secure bool
//	HttpOnly bool
//	Raw string
//	Unparsed []string  //未解析的attribute-value属性位对
//}

//下面看一个如何设置Cookie的例子
//expiration := time.Now()
//expiration = expiration.AddDate(1,0,0)
//cookie := http.Cookie{
//	Name:"username",
//	Value:"zuolan",
//	Expires:expiration,
//}
//http.SetCookie(w, &cookie)

//读取Cookie
// 上面的例子演示了如何设置Cookie数据
//这里演示如何读取Cookie
//cookie, _ :=r.Cookie("username")fmt.Fprint(w, cookie)
//还有另外一种读取方式：
//for _, cookie := range r.Cookies() {
//	fmt.Fprint(w, cookie.Name)
//}
//可以看到通过request获取Cookie非常方便
///

func main() {

}
