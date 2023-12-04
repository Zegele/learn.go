package main

import (
	"context"
	"fmt"
)

// func WithValue(parent Context, key, val interface{}) Context
// WithValue 函数接收context并返回派生的context， 其中值val与key关联，并通过context树与context一起传递。
// 这意味着一旦获得带有值的context，从中派生的任何context都会获得此值。不建议使用context值传递关键参数，函数应该接收签名中的那些值，使其显示化。
// 所提供的键必须是可比较的，并且不应该是string类型或任何其他类型，以避免使用上下文在包之间发生冲突。
// WithValue的用户应该为键定义自己的类型，为了避免在分配给接口{}时进行分配，上下文键通常具有具体类型struct{}。
//或者，导出的上下文关键变量的静态类型应该是指针或接口。

func main() {
	type favContextKey string // 定义一个key类型
	// f : 一个从上下文中根据key取value的函数
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	k := favContextKey("language")
	k2 := favContextKey("lan")
	// 创建一个携带key为k， value为“GO”的上下文
	ctx := context.WithValue(context.Background(), k, "GO")
	ctx = context.WithValue(ctx, k2, "GOO")

	f(ctx, k)
	f(ctx, k2)
	fmt.Println(ctx.Value(k2))
}

//使用Context的注意事项：
// 1、 不要把Context放在结构体中，要以参数的方式显示传递
// 2. 以Context作为参数的函数方法，应该把Context作为第一个参数。
// 3. 给一个函数方法传递 Context 的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
// 4. Context的Value相关方法应该传递请求域的必要数据，不应该拥有传递可选参数
// 5. Context是线程安全的，可以放心的在多个Goroutine中传递。

// 在真正使用值传递的功能时我们也应该非常谨慎，不能将请求的所有参数都使用Context进行传递，这是一种非常差的设计，
// 比较常见的使用场景是传递请求对应用户的认证令牌以及用于进行分布式追踪的请求ID
