// go语言session的创建和管理（懵的）
// www.kancloud.cn/imdszxs/golang/1509695
package main

import (
	"container/list"
	_ "github.com/astaxie/session"
	"sync"
	"time"
)

//前面《Cookie设置与读取》一节我们介绍了Cookie的应用，
//本节我们将讲解session的应用，
//我们知道session是在服务器端实现的一种用户和服务器之间认知的解决方案
//目前Go语言标准包没有为session提供任何支持
//接下来我们将会自己动手来实现go版本的session管理和创建

//session创建过程
//session的基本原理是由服务器为每个会话维护一份信息数据
//客户端和服务端依靠一个全局唯一的表示来访问这个数据
//以达到交互的目的
//当用户访问web 时
//服务端程序会随需要创建session这个过程可以概括为3个步骤：
//1. 生成全局唯一标识符（sessionid）
//2. 开辟数据存储空间
//一般会在内存中创建相应的数据结构
//但这种情况下，系统一旦掉电，所有的会话数据就会丢失
//如果时电子商务类网站
//这将造成严重的后果，所以为了解决这类问题
//可以将会话数据写到文件里或存储在数据库中，
//当然这样会增加I/O开销
//但是它可以实现某种程度的session持久化，也更有利于session的共享
//3. 将session的全局唯一标识符发送给客户端

//以上3个步骤中，最关键的是如何发送这个session的唯一标识这一步上
//考虑到HTTP协议的定义，数据无非可以放到请求行，头域或Body里，
//所以一般来说会有两种常用的方式：
//cookie和URL重写
//Cookie
//服务端通过设置Set-cookie头就可以将session的标识符传送到客户端
//而客户端此后的每一次请求都会带上这个标识符
//另外，一般包含session信息的cookie会将失效时间设置为0（会话cookie）
//即浏览器进程有效时间
//至于浏览器怎么处理这个0，每个浏览器都有自己的方案
//但差别都不会太大（一般体现在新建浏览器窗口的时候）

//URL重写
//就是在返回给用户的页面里的所有URL后面追加session标识符，
//这样用户在收到相应之后，无论点击相应页面里的哪个链接或提交表单
//都会自动带上session标识符
//从而就实现了会话的保持
//虽然这种做法比较麻烦，但是，如果客户端禁用了cookie的话，此种方案将会是首选

//Go实现session管理
//通过上面session创建过程的讲解
//应该对session有一个大体的认识，但是具体到动态页面技术里面
//又是怎么实现session的呢？下面我们将结合session的生命周期（lifecycle）
//来实现Go语言版本的session管理
//session管理设计
//我们知道session管理涉及到如下几个因素
//1. 全局session管理器
//2. 保证sessionid的全局唯一性
//3. 为每个客户关联一个session
//4. session的存储（可以存储到内存，文件，数据库等）
//5. session过期处理
//接下来讲解以下关于session管理的整个设计思路以及相应的go代码示例

//session管理器
//定义一个全局的session管理器
//type Manager struct{
//	cookieName string
//	// private cookiename
//	lock sync.Mutex
//	//protexts session
//	provider Provider
//	maxLifeTime int64
//}
//
//func NewManager(provideName, cookieName string, maxLifeTime int64)(*Manage, error){
//	provider, ok := provides[provideName]
//	if !ok {
//		return nil, fmt.Errorf("session: unknow provide %q (forgotten import?)", provideName)
//	}
//	return &Manager{
//		provider : provider,
//		cookieName : cookieName,
//		maxLifeTime : maxLifeTime
//	}, nil
//}

//go实现整个的流程应该也是这样的
//在main包中创建一个全局的session管理器
//var globalSession *session.Manager
//然后在init函数中初始化
//func init(){
//	globalSessions, _ = NewManager("memory","gosessionid",3600)
//}
//我们知道session是保存在服务器端的数据
//它可以以任何的方式存储
//比如存储在内存，数据库或者文件中
//因此我们抽象出一个Provider接口
//用以表征session管理器底层存储结构
//type Provider interface{
//	Sessioninit(sid string)(Session, error)
//	SessionRead(sid string)(Session, error)
//	SessionDestroy(sid string)error
//	SessionGC(maxLifeTime int64)
//}
//1. SessionInit 函数实现session的初始化，操作成功则返回此新的session变量
//2. SessionRead 函数返回sid所代表的session变量，如果不存在，那么将以sid为参数调用SessionInit函数创建并返回一个新的session变量
//3. SessionDestroy 函数用来销毁sid对于的session变量
//4. SessionGC根据maxLifeTime来删除过期的数据

//那么session接口需要实现什么样的功能呢？
//有过Web开发经验的读者都知道
//对Session的处理基本就：设置值，读取值，删除值，以及获取当前sessionID这四个操作
//所以我们的session接口也就实现这四个操作
//type Session interface{
//	Set(key, value interface{})error // set session value
//	Get(key, interface{}) interface //get session value
//	Delete(key interface{})error //delete session value
//	SessionID()string // back current sessionID
//}
//以上设计思路来源于database/sql/driver
//先定义好接口，然后具体的存储session的结构实现相应的接口并注册后
//相应功能这样就可以使用了
//以下是用来随需注册存储session的结构的Register函数的实现
//var provides = make(map[string]Provider) // register 通过提供的名称提供会话
//	// 如果用想用的名称调用两次register，或者如果driver为nil，它会恐慌。
//func Register(name string, provider Provider){
//	if provider == nil{
//		panic("session: Register provider is nil")
//	}
//	if _, dup := provides[name]; dup{
//		panic("session: Register called twice for provider " + name)
//	}
//	provides[name] = provider
//}

//全局唯一的session ID
//session ID是用来识别访问Web应用的每一个用户
//因此必须保证它是全局唯一的（GUID）
//下面代码展示了如何满足这一需求：
//func (manager *Manager)sessionID()string{
//	b := make([]byte, 32)
//	if _, err := rand.Read(b); err != nil{
//		return ""
//	}
//	return base64.URLEncoding.EncodeToString(b)
//}
//
//session创建
//我们需要为每个来访用户分配或获取与他相关连的session
//以便后面根据session信息来验证操作
//SessionStart 这函数就是用来检测是否已经有某个session与当前来访用户发生了关联，
//如果没有则创建之
//func (manager *Manager)SessionStart(w http.ResponseWriter, r *http.Request)(session Session){
//	manager.lock.Lock()
//	defer manager.lock.Unlock()
//	cookie, err := r.Cookie(manager.cookieName)
//	if err != nil || cookie.Value == ""{
//		sid := manager.sessionId()
//		session, _ = manager.provider.SessionInit(sid)
//		cookie := http.Cookie{
//			Name: manager.cookieName,
//			Value: url.QueryEscape(sid),
//			Path:"/",
//			HttpOnly:true,
//			MaxAge: int(manager.maxLifeTime)
//		}
//	http.SetCookie(w, &cookie)
//	}else{
//		sid, _ := url.QueryUnescape(cookie.Value)
//		session, _ = manager.provider.SessionRead(sid)
//	}
//	return
//}

//我们用前面login操作来演示session的运用：
//func login(w http.ResponseWriter, r *http.Request){
//	sess := globalSessions.SessionStart(w,r)
//	r.ParseForm()
//	if r.Method == "GET"{
//		t, _ := template.ParseFile("login.gtpl")
//		w.Header().Set("Content-Type", "text/html")
//		t.Execute(w, sess.Get("username"))
//	}else{
//		sess.Set("username", r.Form["username"])
//		http.Redirect(w, r, "/", 302)
//	}
//}

//操作值：设置，读取，和删除
//SeesionStart函数返回的是一个满足session接口的变量
//那么我们该如何用他来对session数据进行操作呢？
//上面的例子中的代码session.Get("uid")已经展示了基本的读取数据的操作
//现在我们再来看一下详细的操作：
//func count(w http.ResponseWriter, r *http.Request){
//	sess := globalSession.SessionStart(w, r)
//	createtime := sess.Get("createtime")
//	if createtime == nil{
//		sess.Set("createtime", time.Now().Unix())
//	}else if (createtime.(int64)+360)<(time.Now().Unix()){
//		globalSessions.SessionDestroy(w, r)
//		sess = globalSessions.SessionStart(w, r)
//	}
//	ct := sess.Get("countnum")
//	if ct == nil{
//		sess.Set("countnum", 1)
//	}else{
//		sess.Set("countnum", (ct.(int)+1))
//	}
//	t, _ := template.ParseFiles("count.gtpl")
//	w.Header().Set("Content-Type","text/html")
//	t.Execute(w, sess.Get("countnum"))
//}

//通过上面的例子可以看到，session的操作和操作key/value数据库类似：
//Set，Get，Delete等操作
//因为session有过期的概念
//所以我们定义了GC操作，当访问过期时间满足GC的触发条件后将会引起GC
//但是当我们进行了任意一个session操作
//都会对session实体进行更新
//都会触发对最后访问时间的修改
//这样当GC的时候就不会误删除还在使用的session实体

//session重置
//我们知道，Web应用中有用户退出这个操作，
//那么当用户退出应用的时候，我们需要对该用户的session数据进行销毁操作
//上面的代码已经演示了如何使用session重置操作
//下面这个函数就是实现了这个功能
//Destroy sessionid
//func (manager *Manager)SessionDestroy(w http.ResponseWriter, r *http.Request){
//	cookie, err := r.Cookie(manager.cookieName)
//	if err != nil || cookie.Value == ""{
//		return
//	}else{
//		manager.lock.Lock()
//		defer manager.lock.Unlock()
//		manager.provider.SessionDestroy(cookie.Value)
//		expiration := time.Now()
//		cookie := http.Cookie{
//			Name:manager.cookieName,
//			Path:"/",
//			HttpOnly:true,
//			Expires:expiration,
//			MaxAge: -1
//		}
//		http.SetCookie(w, &cookie)
//	}
//}

//session销毁
//来看以下session管理器如何来管理销毁，
//只要我们在main启动的时候启动
//func init(){
//	go globalSessions.GC()
//}
//func(manager *Manager)GC(){
//	manager.lock.Lock()
//	defer manager.lock.Unlock()
//	manager.provider.SessionGC(manager.maxLifeTime)
//	time.AfterFunc(time.Duration(manager.maxLifeTime), func(){
//		manager.GC()
//	})
//}

//我们可以看到GC充分利用了time包中的定时器功能，
//当超时maxLifeTime之后调用GC函数
//这样就可以保证maxLifeTime时间内的session都是可用的
//类似的方案也可以用于统计在线用户数之类的
//至此，我们实现了一个用来在web应用中全局管理session的SessionManager
//定义了用来提供session存储实现Provider的接口
//接下来，我们将通过接口定义来实现一些Provider，供参考学习：

//session存储
//上面我们介绍了session管理器的实现原理，
//定义了存储session的接口
//接下来我们将示例一个基于内存的session存储接口的实现
//其他的存储方式，
//大家可以自定参考示例来实现
//内存的实现请看下面的例子：

var pder = &Provider{
	list: list.New(),
}

type SessionStore struct {
	sid          string                      // session id 唯一标识
	timeAccessed time.Time                   // 最后访问时间
	value        map[interface{}]interface{} //session里面存储的值
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

type Provider struct {
	lock     sync.Mutex               // 用来锁
	sessions map[string]*list.Element // 用来存储在内存
	list     *list.List               //用来做gc
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{
		sid:          sid,
		timeAccessed: time.Now(),
		value:        v,
	}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
	}
	return nil
}

func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
}

// 上面这个代码实现了一个内存存储的session机制
// 通过init函数注册到session管理器中
//这样就可以方便的调用了
//我们如何来调用该引擎呢？请看下面的代码：
//import(
//"github.com/astaxie/session"
//_ "github.com/astaxie/session/providers/memory"
//)
//当import的时候已经执行了memory函数里面的init函数
//这样就已经注册到session管理器中
//我们就可以使用了
//通过如下方式就可以初始化一个session管理器
//var globalSessions *session.Manager // 然后在init函数中初始化
//func inti(){
//	globalSessions, _ = session.NewManager("memory","gosessionid",3600)
//	go globalSessions.GC()
//}

//预防session劫持
//我们写了如下的代码来展示一个count计数器
//func count(w http.ResponseWriter, r *http.Request){
//	sess := globalSessions.SessionStart(w, r)
//	ct := sess.Get("countnum")
//	if ct == nil{
//		sess.Set("countnum", 1)
//	}else{
//		sess.Set("countnum", (ct.(int)+1))
//	}
//	t, _ := template.ParseFiles("count.gtpl")
//	w.Header().Set("Content-Type", "text/html")
//	t.Execute(w, sess.Get("countnum"))
//}

//count.gtpl的代码如下所示：
//Hi.Now count:{{.}}
//然后我们在浏览器里面刷新可以看到如下内容：
//图：浏览器端显示count数
//随着刷新，数字将不断增长，当数字显示为6的时候
//打开浏览器（以chrome为例）的cookie管理器，可以看到类似如下的信息
//图：获取浏览器端保存的cookie

//下面这个步骤最为关键：打开另一个浏览器（这里我们打开了firefox浏览器）
//赋值chrome地址栏里的地址到新打开的浏览器的地址栏中
//然后打开firefox的cookie模拟插件
//新建一个cookie，把按上图中的cookie内容原样在firefox中重建一份
//图：模拟cookie
//回车后，大家将看到如下内容：
//图：劫持session成功
//可以看到虽然换了浏览器，但是我们却获得了sessionID，
//然后模拟了cookie存储的过程
//这个例子是在同一台计算机上做的
//不过计时换用两台来做，其结果任然一样
//此时如果交替点击两个浏览器里的链接您会发现他们其实操纵的是同一个计数器
//不必惊讶，此处firefox盗用了chrome和goserver之间的维持会话的钥匙
//即gosessionid，这是一种类型的“会话劫持”
//在goserver看来，它从http请求中得到了一个gosessionid
//由于HTTP协议的无状态性，它无法得知这个gosessionid是从chrome那里“劫持”来的
//它依然会查找对应的session，并执行相关计算
//与此同时chrome也无法得知自己保持的会话已经被“劫持”

//session劫持防范
//cookieonly和token
//通过上面session劫持的简单演示可以了解到session一旦被其他人劫持
//就非常危险，劫持者可以假装成劫持者进行很多非法操作
//那么如何有效的防止session劫持呢？
//其中一个解决方案就是sessionID的值只允许cookie设置
//而不是通过URL重置方式设置
//同时设置cookie的httponly为true
//这个属性是设置是否可通过客户端脚本访问这个设置的cookie
//第一这个可以防止这个cookie被XSS读取从而引起session劫持
//第二cookie设置不会像URL重置方式那么容易获取sessionID
//第二步就是在每个请求里面加上token
//实现类似前面章节里面讲的防止form重复提交类似的功能
//我们在每个请求里面加上一个隐藏token
//然后每次验证这个token
//从而保证用户的请求都是唯一性
//h := md5.New()
//salt:="astaxie%^7&8888"
//io.WriteString(h, salt+time.Now().String())
//token := fmt.Sprintf("%x", h.Sum(nil))
//if r.Form["token"]!=token{
//	//提示登录
//}
//sess.Set("token",token)

//间隔生成新的SID
//还有一个解决方案就是，我们给session额外设置一个创建时间的值，
//一旦过了一定的时间，
//我们销毁这个sessionID，
//重新生成新的session
//这样可以一定程度上防止session劫持的问题
//createtime := sess.Get("createtime")
//if createtime == nil{
//	sess.Set("createtime", time.Now().Unix())
//}else if (createtime.(int64)+60)<(time.Now().Unix()){
//	globalSessions.SessionDestory(w, r)
//	sess = globalSessions.SessionStart(w, r)
//}

//session启动后，我们设置了一个值，用于记录生成sessionID的时间
//通过判断每次请求是否过期（这里设置了60秒） 定期生成新的ID
//这样使得攻击者获取有效sessionID的机会大大降低
//上面两个手段的组合可以在实践中消除session劫持的风险
//一方面，由于sessinoID频繁改变
//使攻击者难有机会获取有效的sessionID
//另一方面，因为sessionID只能在cookie中传递
//然后设置了httponly
//所以基于URL攻击的可能性为零
//同时被XSS获取sessionID也不可能
//最后，由于我们还设置了MaxAge=0
//这就相当于session cookie不会留在浏览器的历史记录里面

func main() {

}
