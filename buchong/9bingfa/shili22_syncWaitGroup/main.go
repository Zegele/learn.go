package main

import (
	"fmt"
	"net/http"
	"sync"
)

//www.likecs.com/show-306390439.html
// 等待组：多个任务的同步，等待组可以保证在并发环境中指定数量的任务
//在sync.WaitGroup（等待组）类型中，每个sync.WaitGroup值在内部维护着一个计数，此计数的初始默认值为零。
// 有以下方法
//(wg *WaitGroup)Add(delta int) //等待组的计数+delta
//(wg *WaitGroup)Done() //等待组的计数-1
//(wg *WaitGroup)Wait() //当等待组计数器不等于0时，阻塞，直到变0

//N个并发任务进行工作时，就将等待组的计数值增加N。每个任务完成时，这个值减1。
//同时，在另外一个goroutine中等待（Wait()）这个等待组的计数器值为0时，表示所有任务已经完成。

func main() {
	// 声明一个等待组
	var wg sync.WaitGroup

	// 准备一系列的网站地址
	var urls = []string{
		"https://www.github.com",
		"https://www.qiniu.com",
		"https://www.baidu.com",
		//"https://www.golangtc.com",
	}
	//wg.Add(3)// 为什么不放出来？可能因为不知道有多少个元素要遍历
	// 如果知道有多少个任务，可以这样。
	// 遍历这些地址
	for _, url := range urls {
		// 每一个任务开始时，将等待组增加1
		wg.Add(1)

		// 开启一个并发
		go func(url string) {
			// 使用defer，表示函数完成时将等待组值减1
			defer wg.Done()
			// 使用http访问提供的地址
			_, err := http.Get(url)
			//http包提供的Get() 函数对url进行访问，Get() 函数会一直阻塞直到网站相应或者超时

			// 访问完成后，打印地址和可能发生的错误
			fmt.Println(url, err)

			// 通过参数传递url地址
		}(url)
		// 这里将url通过goroutine的参数进行传递，是为了避免url变量通过闭包放入匿名函数后又被修改的问题。
	}

	// 等待所有的任务完成
	wg.Wait()

	fmt.Println("over")
}
