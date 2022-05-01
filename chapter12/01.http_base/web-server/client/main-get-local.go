package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	// 这里的 5秒是，控制整个程序在5秒内运行完成。
	defer cancel()
	go httpDirectGet()
	go httpGetWithContext(ctx)
	<-ctx.Done() //等着结束
	//作用是什么？
}

func httpDirectGet() {
	resp, err := http.Get("http://localhost:8088")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("directGet", string(data))
}
func httpGetWithContext(ctx context.Context) { //一般设置参数ctx
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) //ctx传入这里
	// context 把控时间1s，如果1秒没有回应，就结束。
	// context是贯穿整个函数周期的。
	//func httpGetWithContext() {//老师刚开始是这样写的
	//	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)//老师刚开始是这样写的
	defer cancel()
	req, err := http.NewRequest("get", "http://localhost:8088", nil) // 生成一个请求
	if err != nil {
		log.Fatal("无法生成请求：", err)
	}
	req = req.WithContext(ctx) // 带上ctx  context（带一个上下文）
	//注意！ 务必重新赋值req
	// 很容易出错
	resp, err := http.DefaultClient.Do(req) // Do(rep)是把需求发送给服务端？然后返回的数据给resp
	if err != nil {
		log.Println("无法发送请求：", err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body) //读取返回数据中的Body的内容
	if err != nil {
		log.Fatal("无法读取返回内容：", err)
	}
	fmt.Println("ContextGet", string(data))
}
