package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com") //填网址
	// resp 是访问网址返回的数据（是个结构体类型）
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body) //读取返回数据中的Body数据。
	// ioutil.ReadAll是读取所有的数据。适用于内容不大的数据？要看下ioutil包。
	fmt.Println(string(data))
}
