package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learn.go/chapter12/apiss"
	"learn.go/chapter12/frinterface"
	"learn.go/chapter12/rank"
	"log"
	"net/http"
	"strings"
)

func main() {
	var rankServer frinterface.ServeInterface = rank.NewFatRateRank() //要学会这样使用

	m := http.NewServeMux()

	m.Handle("/register", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "post") { //strings.EqualFold 忽略大小写，且对这两个参数进行相等
			//if request.Method != "post" // request.Method  是对request进行什么方法的操作？
			// 如果request.Method不是post方法
			writer.WriteHeader(http.StatusBadRequest) //400 4开头客户端出问题
			return
		}

		if request.Body == nil {
			//如果Body是空的， 不能注册了，但没有内容。Body就是payload
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()

		payload, err := ioutil.ReadAll(request.Body) // 如果有body，读取body，到payload
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法读取数据： %s", err)))
			return
		}

		var pi *apiss.PersonalInformation
		if err := json.Unmarshal(payload, &pi); err != nil { //把读取到的数据unmarshal解析
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法解析数据： %s", err)))
			return
		}

		if err := rankServer.RegisterPersonalInformation(pi); err != nil {
			writer.WriteHeader(http.StatusInternalServerError) //如果还是nil，就说明是服务端出错。
			writer.Write([]byte(fmt.Sprintf("注册失败： %s", err)))
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`success`))

	}))

	m.Handle("/personalinfo", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "post") { //strings.EqualFold 忽略大小写
			// 如果不是post方法
			writer.WriteHeader(http.StatusBadRequest) //400
			return
		}

		if request.Body == nil {
			//如果Body是空的
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()

		payload, err := ioutil.ReadAll(request.Body) // 读取body
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法读取数据： %s", err)))
			return
		}

		var pi *apiss.PersonalInformation
		if err := json.Unmarshal(payload, &pi); err != nil { //把读取到的数据unmarshal解析
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法解析数据： %s", err)))
			return
		}

		if fr, err := rankServer.UpdatePersonalInformation(pi); err != nil {
			writer.WriteHeader(http.StatusInternalServerError) //如果还是nil，就说明是客户端出错。
			writer.Write([]byte(fmt.Sprintf("更新失败： %s", err)))
			return
		} else {

			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(fr) //unmarshal的数据是解析的有不可掌控，但marshal的数据是我们可以掌控的。
			writer.Write(data)
			return
		}

	}))

	m.Handle("/ranks", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "get") { //strings.EqualFold 忽略大小写
			// 如果不是post方法
			writer.WriteHeader(http.StatusBadRequest) //400
			return
		}

		name := request.URL.Query().Get("name") // localhost:8080/ranks?name=xiaoqiang
		if name == "" {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("name参数未设置"))
			return
		}

		if fr, err := rankServer.GetFatRate(name); err != nil {
			writer.WriteHeader(http.StatusInternalServerError) //如果还是nil，就说明是客户端出错。
			writer.Write([]byte(fmt.Sprintf("获取排行失败： %s", err)))
			return
		} else {

			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(fr) //unmarshal的数据是解析的有不可掌控，但marshal的数据是我们可以掌控的。
			writer.Write(data)
			return
		}

	}))

	m.Handle("/ranktop", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "get") { //strings.EqualFold 忽略大小写
			// 如果不是post方法
			writer.WriteHeader(http.StatusBadRequest) //400
			return
		}

		if frTop, err := rankServer.GetTop(); err != nil {
			writer.WriteHeader(http.StatusInternalServerError) //如果还是nil，就说明是客户端出错。
			writer.Write([]byte(fmt.Sprintf("获取排行失败： %s", err)))
			return
		} else {

			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(frTop) //unmarshal的数据是解析的有不可掌控，但marshal的数据是我们可以掌控的。
			writer.Write(data)
			return
		}
	}))

	if err := http.ListenAndServe(":8080", m); err != nil { //注意！！！ ":8080"冒号不能省略
		log.Fatal(err)
	}

}

// 运行后
//http://localhost:8080/ranktop  --> []
//http://localhost:8080/rank?name=nine
//{"Name":"nine","Sex":"","RankNumber":0,"FatRate":0}
