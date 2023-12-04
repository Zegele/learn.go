// go语言反射——性能和灵活性的双刃剑
// blog.csdn.net/Mr_XiMu/article/details/122386171
package main

import (
	"reflect"
	"testing"
)

/*
// 现在一些流行设计思想需要建立在反射基础上，如控制反转（Inversion Of Control,IOC）和依赖注入（Dependency Injection，DI）。
//Go语言有名的web框架martini（github.com/go-martini/martini）就是通过依赖注入计数进行中间件的实现，例如使用martini框架搭建的http的服务器：
import (
	_ "github.com/go-martini/martini"
) //这个库还没拉

func main() {
	m := martini.Classic()
	m.Get("/", func() string { // 相应路径/的代码使用一个闭包实现。
		return "Hello World"
	})
	//如果希望获得Go语言中提供的请求和响应接口，可以直接修改为;
	//m.Get("/",func(res http.ResponseWriter,req *http.Request)string{
	// 响应处理代码。。。
	//})
	m.Run()
}

//martini的底层会自动通过识别Get获得的闭包参数情况，通过动态反射调用这个函数并传入需要的参数。
//martini的设计广受好评，但同时也有人指出，其运行效率较低，其中最主要的因素是大量使用了反射

*/

//虽然一般，I/O的延迟远远大于反射代码所造成的延迟。但是，更低的响应速度和更低的cpu占用依然是web服务器追求的目标。
//因此，反射再带来灵活性的同时，也带上了性能低下的桎梏
//了解反射的性能，一些基准测试从多方面对比了原生调用和反射调用的区别。.

// 1. 结构体成员赋值对比
//反射经常被使用在结构体上，因此结构体的成员访问性能就成为了关注的重点。
//例子，使用一个被实例化的结构体，访问它的成员，然后使用Go语言的基准化测试可以迅速测试出结果
//基准测试以Benchmark为前缀，需要一个testing.B类型的参数b。

// 声明一个结构体，拥有一个字段
type data struct { // 声明一个普通结构体，拥有一个成员变量
	Hp int
}

// 原生结构体的赋值过程：
func BenchmarkNativeAssign(b *testing.B) { // 使用基准化测试的入口
	// 实例化结构体
	v := data{Hp: 2} // 实例化data结构体，并给Hp成员赋值
	// 停止基准测试的计时器
	b.StopTimer()
	// 重置基准测试计时器数据
	b.ResetTimer() // 将基准测试的计时器复位
	//重启启动基准测试计时器
	b.StartTimer() //将基准测试的计时器重启，这样做更精确

	// 根据基准测试数据进行循环测试
	for i := 0; i < b.N; i++ {
		// 结构体成员赋值测试
		v.Hp = 3
	}
}

// 使用反射访问结构体成员并赋值的过程
func BenchmarkReflectAssign(b *testing.B) {
	v := data{Hp: 2}

	// 取出结构体指针的反射值对象并取其元素
	vv := reflect.ValueOf(&v).Elem()
	// 取v的地址并转为反射值对象。此时值对象里的类型为*data，使用值的Elem()方法取元素，获得data的反射值对象。

	// 根据名字取结构体成员
	f := vv.FieldByName("Hp")

	b.StartTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		// 反射测试设置成员值性能
		f.SetInt(3)
		// 使用反射值对象的SetInt()方法，给data结构的Hp字段设置数值3
	}
}

// SetInt() 源码
//func(v Value) SetInt(x int64){
//	v.mustBeAssignable()
//	switch k:=v.kind(); k{
//	default:
//		panic(&ValueError{"reflect.Value.SetInt", v.kind()})
//	case Int:
//		*(*int)(v.ptr) = int(x)
//	case Int8:
//		*(*int8)(v.ptr) = int8(x)
//	case Int16:
//		*(*int16)(v.ptr) = int16(x)
//	case Int32:
//		*(*int32)(v.ptr) = int32(x)
//	case Int64:
//		*(*int64)(v.ptr) =x
//	}
//}
// 可以发现，整个设置过程都是指针转换即赋值，没有遍历及内存操作等相对耗时的算法。

// 2. 结构体成员搜索并赋值对比
func BenchmarkReflectFindFieldAndAssign(b *testing.B) {
	v := data{Hp: 2}

	vv := reflect.ValueOf(&v).Elem()

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		// 测试结构体成员的查找和设置成员的性能
		vv.FieldByName("Hp").SetInt(3)
	}
}

// 上面这段代码，将反射值对象的FieldByName()方法与SetInt()方法放在循环里进行检测，
//主要对比测试FieldByName()方法对性能的影响。
// FieldByName()方法源码如下：
//func (v Value) FieldByName(name string) Value {
//	v.mustBe(Struct)
//	if f, ok := v.typ.FieldByName(name); ok { //通过名字查询类型对象，这里有一次遍历过程
//		return v.FieldByIndex(f.Index) // 找到类型对象后，使用FieldByIndex()继续在值中查找，这里又是一次遍历。
//	}
//	return Value{}
//}
// 经过底层代码分析得出，随着结构体字段数量和相对位置的变化，FieldByName()方法比较严重的低效率问题。

// 3. 调用函数对比
// 反射的函数调用，也是使用反射中容易忽视的性能点，下面展示对普通函数的调用过程。

// 一个普通函数
func foo(v int) { // 一个普通的只有一个参数的函数
}

func BenchmarkNativeCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 原生函数调用
		foo(2) // 对原生函数调用的性能测试
	}
}

func BenchmarkReflectCall(b *testing.B) {
	// 取函数的反射值对象
	v := reflect.ValueOf(foo) // 根据函数名取出反射值对象

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		// 反射调用函数
		v.Call([]reflect.Value{reflect.ValueOf(2)})
		// 使用reflect.ValueOf(2)将2构造为反射值对象，因为反射函数调用的参数必须全是反射值对象
		// 再使用[]reflect.Value构造多个参数列表传给反射值对象的Call()方法进行调用。
	}
}

//反射函数调用的参数构造过程非常复杂，构建很多对象会造成很大的内存回收负担。
//Call()方法内部就更为复杂，需要将参数列表的每个值从reflect.Value类型转化为内存
//调用完毕后，还要将函数返回值重新转换为reflect.Value类型返回。
//因此，反射调用函数的性能堪忧.

// go test -v - bench=.  // 在git中会有详细信息。 goland中没有详细信息
//$ go test -v -bench=.
//goos: windows
//goarch: amd64
//pkg: learn.go/buchong/10fanshe_reflect/shili3_shuangrenjian
//cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
//BenchmarkNativeAssign
//BenchmarkNativeAssign-4                 1000000000               0.3482 ns/op // 原生结构体成员赋值：每一部操作耗时0.32纳秒，这是参考基准
//BenchmarkReflectAssign
//BenchmarkReflectAssign-4                344812029                3.393 ns/op // 反射赋值，每次操作耗时3.393纳秒，是原生的10倍
//BenchmarkReflectFindFieldAndAssign
//BenchmarkReflectFindFieldAndAssign-4    13665636                85.66 ns/op // 反射查找结构体成员并反射赋值，每次操作耗时85.66纳秒， FieldByName导致性能下降太多
//BenchmarkNativeCall
//BenchmarkNativeCall-4                   1000000000               0.4009 ns/op
//BenchmarkReflectCall
//BenchmarkReflectCall-4                   5909800               195.9 ns/op // 反射函数调用，性能差到爆，单次耗时195.9纳秒,是原生的496倍。
//PASS
//ok      learn.go/buchong/10fanshe_reflect/shili3_shuangrenjian    5.263s

// 经过基准测试结果的数值分析及对比，最终得出一下结论：
//1.能使用原生代码时，尽量避免反射操作
//2. 提前缓冲反射值对象，对性能有很大的帮助
//3. 避免反射函数调用，实在需要调用时，先提前缓冲函数参数列表，并且尽量少地使用返回值.
