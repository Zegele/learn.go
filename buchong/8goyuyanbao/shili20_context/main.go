package main

// cancel
// 自己改的
//func genn(ctx context.Context) <-chan int {
//	dst := make(chan int)
//	n := 1
//	go func() {
//		for {
//			select {
//			case <-ctx.Done():
//				return // return 结束该goroutine， 防止泄露
//			case dst <- n:
//				n++
//			}
//		}
//	}()
//	return dst
//}

//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel() // 当我们取完需要的整数后调用cancel
//	var gg <-chan int
//
//	gg = genn(ctx)
//
//	for n := range gg {
//		fmt.Println(n)
//		if n == 5 {
//			break
//		}
//	}
//}

/*

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // return 结束该goroutine， 防止泄露
				case dst <- n:
					n++
				}
			}
		}()
		fmt.Println("gaahfaga")
		defer fmt.Println(1243)
		fmt.Println(222)
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //当我们取完需要的整数后调用canceel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

*/

/* WithDeadline

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的时间。
	// 尽管不这样做，可能会使上下文及其父类存活的时间超过必要的时间
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Printf("%v\n", ctx)
		fmt.Println(ctx)
		fmt.Println(ctx.Err())
	}
	// 上面的代码中，定义了一个50ms之后 过期的deadline， 然后我们调用 context.WithDeadline(context.Backround(), d)
	// 得到一个上下文（ctx）和一个取消函数（cancel），然后使用一个select让主函数陷入等待，等待1m后打印overslept退出，
	//或者等待ctx过期后退出。因为ctx50ms后就过期，所以ctx.Done()会先接收到值，然后打印ctx.Err()取消原因。
}

*/

/*  WithTimeout

func main() {
	// 传递带有超时的上下文
	// 告诉阻塞函数在超时结束后应该放弃工作
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 终端输出“context deadline exceeded”
	}

}

*/
