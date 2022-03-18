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
	defer cancel()
	go httpDirectGet()
	go httpGetWithContext(ctx)
	<-ctx.Done()
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
	//func httpGetWithContext() {//老师刚开始是这样写的
	//	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)//老师刚开始是这样写的
	defer cancel()
	req, err := http.NewRequest("get", "http://localhost:8088", nil)
	if err != nil {
		log.Fatal("无法生成请求：", err)
	}
	req = req.WithContext(ctx) //注意！ 务必重新赋值req
	// 很容易出错
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("无法发送请求：", err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("无法读取返回内容：", err)
	}
	fmt.Println("ContextGet", string(data))
}
