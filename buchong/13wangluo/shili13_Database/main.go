// go语言数据库（database）相关操作（懵的，结合小强看看）
// www.kancloud.cn/imdszxs/golang/1509684
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//本节将对db/spl官方标准库作一些简单分析，并介绍一些应用比较广泛的开源ORM和SQL Builder
//并从企业级应用开发和公司架构的角度来分析哪种技术栈对于现代的企业级应用更为合适
//从database/sql讲起
//go语言官方提供了database/sql包来给用户进行和数据库打交道的工作
//实际上database/sql库就只是提供了一套操作数据库的接口和规范
//例如抽象好的SQL预处理（prepare），连接池管理，数据绑定，事务，错误处理等
//官方并没有提供母体某种数据库实现的协议支持
//和具体的数据库，例如MySQL打交道
//还需要再引入MySQL的驱动，如下：
//import "database/sql"
//import _ "github.com/go-sql-driver/mysql"
//db, err := sql.Open("mysql", "user:password@/dbname")

//import_"github.com/go-sql-driver/mysql" 这句实际是调用了mysql包的init函数
//做的事情也很简单
//func init(){
//	sql.Register("mysql", &MySQLDriver{})
//}
//在sql包的全局map里把mysql这个名字的driver注册上
//实际上Driver在sql包中是一个接口：
//type Driver interface{
//	Open(name string)(Conn, error)
//}
//调用sql.Open()返回的db对象实际上就是这里的Conn，
//type Conn interface{
//	Prepare(query string)(Stmt, error)
//	Close()error
//	Begin()(Tx, error)
//}
//Conn也是一个接口。
//实际上如果仔细查看database/sql/driver/driver.go 的代码会发现
//这个文件里所有的成员全都是接口
//对这些类型进行操作，实际上还是会调用具体的driver里的方法
//从用户的角度来讲，在使用database/sql包的过程中，
//能够使用的也就是这些接口里提供的函数
//来看一个使用database/sql和go-sql-driver/mysql的完整例子：

func main() {
	// db 是一个sql.DB类型的对象
	// 该对象线程安全，且内部已包含了一个连接池
	//连接池的选项可以在sql.DB的方法中设置，这里为了简单省略了
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var (
		id   int
		name string
	)
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// 必须要把rows里的内容读完，或者显示调用Close()方法
	//否则在defer的rows.Close()执行之前，连接永远不会释放
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//如果大家像了解这个database/sql库更加详细的用法，可以参考http://godatabase-sql:org/ 貌似要看看
//包括该库的功能介绍，用法，注意事项和反直觉的一些实现方式（例如同一个goroutine内对sql.DB的查询，可能在多个连接上）都有涉及
//这里不赘述

//通过上面的介绍，也许大家已经发现了一些问题
//官方的db库提供的功能这么简单，
//我们每次去数据库里读取内容岂不是都要去写这么一套差不多的代码？
//或者如果我们的对象是结构体，把sql.Rows绑定到对象的工作就会变得更加得重复而无聊
//所以社区才会有各种各样的SQL Builder和ORM百花齐放
//提高生产效率的ORM和SQL Builder
//在web开发领域常常提到的ORM是什么？我们先看看万能的维基百科：
//对象关系映射（英语：Object Relational Mapping,简称ORM，或O/RM，或O/Rmapping）
//是一种程序设计技术，
//用于实现面向对象编程语言里不同类型系统的数据之间的转换
//从效果上说，它其实是创建了一个可在编程语言里使用的”虚拟对象数据库“
//最为常见的ORM实际上做的是从db到程序的类或结构体这样的映射
//所以你手边的程序可能是从MySQL的表映射你的程序内的类。
//我们可以先来看看其它的程序语言里的ORM写起来怎么样的感觉：
//>>>from blog.models import Blog
//>>>b = Blog(name='Beatles Blog',tagline='All the latest Beatles news.')
//>>>b.save()
//完全没有数据库的痕迹，没错ORM的目的就是屏蔽掉DB层，
//实际上很多语言的ORM只要把你的类或结构体定义好，
//再用特定的语法将结构体之间的一对一或者一对多关系表达出来
//那么任务就完成了
//然后你就可以对这些映射好了数据库表的对象进行各种操作
//例如：save，create，retrieve， delete
//至于ORM在背地里做了什么，你是不一定清楚的，
//使用ORM的时候，我们往往比较容易有一种忘记了数据库的直观感受
//例子，我们有个需求：向用户展示最新的商品列表，我们再假设，商品和商家是1：1的关联关系
//我们就很容易写出像下面这样的代码：
//shopList := []
//for product in productList{
//	shopList = append(shopList, product.GetShop)
//}
//当然了，我们不能批判这样写代码的程序员是偷懒的
//因为ORM一类的工具再出发点上就是屏蔽sql，
//让我们对数据库的操作更接近于人类的思维方式，
//这样很多只接触果ORM而且又是刚入行的程序员就很容易写出上面这样的代码
//这样的代码将对数据库的读写请求放大了n倍
//也就是说，如果你的商品列表有15个SKU，
//那么每次用户打开这个页面，至少需要执行1（查询商品列表）+15（查询相关的商铺信息）次查询
//这里N是16。
//如果你的列表页很大，比如有600个条目，那么就至少执行1+600次查询
//如果说你的数据库能够承受的最大的简单查询是12万QPS，而上述这样的查询正好是最常用的查询的话
//实际上能对外提供的服务能力是多少呢？
//是200qps！
//互联网系统的忌讳之一，就是这种无端的读放大
//当然，也可以说这不是ORM的问题，
//如果手写sql还是可能会写出差不多的程序
//那么再来看两个demo:
//o := orm.NewOrm()
//num, err := o.QueryTable("cardgroup").Filter("Cards_Card_Name", cardName).All(&cardgroups)
//很多ORM都提供了这种Filter类型的查询方式，
//不过实际上在某些ORM背后甚至隐藏了非常难以察觉的细节
//比如生成的SQL语句会自动limit1000
//也许喜欢ORM的读者读到这里会反驳了，你是没有认真阅读文档就瞎写
//是的，尽管这些ORM工具在文档里说明了All查询在不显式地指定Limit的话会自动limit1000
//但对于很多没有阅读果文档或着看过ORM源码的人，
//这依然是一个非常难以察觉的”魔鬼“细节

// 喜欢强类型语言的人一般都不喜欢隐式地去做什么事情
// 例如各种语言在赋值操作时进行的隐式类型转换后又在转换中丢失了精度的勾当
// 一定让你非常的头疼
// 所以一个程序库背地里做的事情还是越少越好
// 如果一定要做，那也一定要在显眼的地方做
// 比如上面的例子，
// 去掉这种默认的自作聪明的行为，
// 或者要求用户强制传入limit参数都是更好的选择
// 除了limit的问题，我们再看一遍这个下面的查询：
// num, err := o.QueryTable("cardgroup").Filter("Cards_Card_Name",cardName).All(&cardgroups)
// 可以看得出来这个Filter是有表join的操作么？
// 当然了，有深入使用经验的用户还是会觉得这是在吹毛求疵
// 但这样的分析想证明的是，ORM想从设计上隐去太多的细节
// 而方便的代价是其背后的运行完全失控
// 这样的项目在经过几任维护人员之后
// 将变得面目全非，难以维护
//
// 当然，我们不能否认ORM的进步意义
// 它的设计初衷就是为了让数据的操作和存储的具体实现所剥离
// 但是在上了规模的公司的人们渐渐达成了一个共识，
// 由于隐藏重要的细节，ORM可能是失败的设计
// 其所隐藏的重要细节对于上了规模的系统开发来说至关重要
//
// 相比ORM来说，SQL Builder在SQL和项目可维护性之间取得了比较好的平衡
// 首先，sql builder不像ORM那样屏蔽了过多的细节
// 其次从开发的角度来讲，
// sql builder简单进行封装后可以非常高效的完成开发，如下：
//
//	where := map[string]interface{
//		"order_id>?" : 0,
//		"customer_id != ?" : 0,
//		}
//
// limit := []int{0,100}
// orderBy := []string{"id asc", "create_time desc"}
// orders := orderModel.GetList(where, limit, orderBy)
// 写SQL Builder的相关代码，或者读懂都不费劲。
// 把这些代码脑内转换为sql也不会太费劲
// 所以通过代码就可以对这个查询是否命中数据库索引
// 是否走了覆盖索引，是否能够用上联合索引进行分析了。
// 说白了SQL Builder 是sql在代码里的一种特殊方言，
// 如果你们没有DBA但研发有自己分析和优化sql的能力，
// 或者你们公司的DBA对于学习这样一些sql的方言没有异议
// 那么使用SQL Builder是一个比较好的选择，不会导致什么问题
// 另外在一些本来也不需要DBA介入的场景内，使用SQL Builder也是可以的
// 例如要做一套运维系统，且将MySQL当做了系统中的一个组件，系统的QPS不高
// 查询不复杂等等
// 一旦你做的是高并发的OLTP在线系统，且想在人员充足分工明确的前提下最大程度控制系统的风险
// 使用SQL Builder就不合适了
//
// 脆弱的数据库
// 无论是ORM还是SQL Builder都有一个致命的缺点
// 就是没有办法尽心该系统上线前sql审核
// 虽然很多ORM和SQL Builder 也提供了运行期打印sql的功能
// 但旨在查询的时候才能进行输出
// 而SQL Builder和ORM本身提供的功能太过灵活，
// 使得你不可能通过测试枚举出所有可能在线上执行sql
// 例如你可能用SQL Builder写出下面这样的代码：
//
//	where := map[string]interface{}{
//		"product_id=?" : 10,
//		"user_id=?" : 1232,
//	}
//
//	if order_id != 0{
//		where["order_id=?"] = order_id
//	}
//
// res, err := historyModel.GetList(where, limit, orderBy)
// 你的系统里有类似上述样例的大量if的话，就难以通过测试用例来覆盖到所有可能的sql组合了
// 这样的系统只要发布，就已经孕育了初期的巨大风险
//
// 对于现在7*24服务的互联网公司来说，服务不可用是非常重大的问题
// 存储层的技术栈虽经历了多年的发展，
// 在整个系统中依然是最为脆弱的一环，
// 系统宕机对于24小时对外提供服务的公司来说，意味着直接的经济损失， 个中风险不可忽视
// 从行业分工的角度来讲，现金的互联网公司都有专职的DBA。
// 大多数DBA并不一定有写代码的能力
// 去阅读SQL Builder的相关“拼SQL”代码多多少少还是会有一点障碍
// 从DBA角度出发，还是希望能够有专门的事前SQL审核机制，
// 并能让其低成本地获取到系统的所有SQL内容
// 而不是去阅读业务研发编写的SQL Builder的相关代码
// 所以现如今，大型的互联网公司核心线上业务都会在代码中把SQL
// 放在显眼的位置提供给DBA评审，举一个例子：
const (
	getAllByProductIDAndCustomerID = `select * from p_orders where product_id in (:product_id) and customer_id =:customer_id `
) // GetAllByProductIDAndCustomerID
// @param driver_id
// @param rate_data
// @return []Order, errorfunc
func getAllByProductIDAndCustomerID(ctx context.Context, productIDs []uint64, customerID uint64) ([]Order, error) {
	var orderList []Order
	params := map[string]interface{}{
		"product_id":  productIDs,
		"customet_id": customID,
	}
	//getAllByProductIDAndCustomerID 是const类型的sql字符串
	sql, args, err := sqlutil.Named(getAllByProductIDAndCustomerID, params)
	if err != nil {
		return nil, err
	}
	return orderList, err
}

// 像这样的代码，在上线之前把DAO层的变更集的const部分直接俄拿给DBA来进行审核
//就比较方便了，代码中的sqlutil.Named是类似于sqlx中的Named函数
//同时支持where表达式中比较操作符和in，
