package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"learn.go/pkg/apis"
	"log"
	"time"
)

func main() {
	// 10 万条记录
	counter := 10
	persons := make([]*apis.PersonalInfomation, 0, counter) //[]*apis.PersonalInfomation 是这种类型的切片啊！
	for i := 0; i < counter; i++ {
		persons = append(persons, &apis.PersonalInfomation{
			Name:   "123",
			Sex:    "男",
			Tall:   1.8,
			Weight: 65,
			Age:    33,
		})
	}
	// JSON, YAML, Protobuf分别序列化，记录序列化耗时
	// 保存文件，记录耗时
	//JSON Marshal
	{
		fmt.Println("序列化JSON")
		startTime := time.Now()
		data, err := json.Marshal(persons)
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		ioutil.WriteFile("E:/Geek/src/learn.go/chapter08/0.4.think/data.json", data, 0777)
		finishWriteFileTime := time.Now()
		fmt.Println("序列化耗时：", finishMarshalTime.Sub(startTime))    //序列化耗时： 86.2638ms
		fmt.Println("写文件按耗时：", finishWriteFileTime.Sub(startTime)) //写文件按耗时： 96.2567ms

	}

	//JSON Unmarshal
	{
		startTime := time.Now()
		data, _ := ioutil.ReadFile("E:/Geek/src/learn.go/chapter08/0.4.think/data.json")
		json.Unmarshal(data, &persons)
		finishTime := time.Now()
		fmt.Println("JSON Unmarshal:", finishTime.Sub(startTime))
	}

	/*
		{
			fmt.Println("序列化YAML")
			startTime := time.Now()
			data, err := yaml.Marshal(persons)
			if err != nil {
				log.Fatal(err)
			}
			finishMarshalTime := time.Now()
			ioutil.WriteFile("E:/Geek/src/learn.go/chapter08/0.4.think/data.yaml", data, 0777)
			finishWriteFileTime := time.Now()
			fmt.Println("序列化耗时：", finishMarshalTime.Sub(startTime))
			fmt.Println("写文件按耗时：", finishWriteFileTime.Sub(startTime))
		}

	*/

	//PROTOBUF Marshal
	{
		fmt.Println("序列化PROTOBUF")
		startTime := time.Now()
		pLister := &apis.PersonalInfomationList{
			Items: persons, // 注意这里！
		}
		data, err := proto.Marshal(pLister) //proto.Marshal参数是要message类型
		// 但我们要很多数据，所以就把message嵌套成立一个很多数据一起的切片，也满足了message类型。
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		ioutil.WriteFile("E:/Geek/src/learn.go/chapter08/0.4.think/data.protobuf", data, 0777)
		finishWriteFileTime := time.Now()
		fmt.Println("序列化耗时：", finishMarshalTime.Sub(startTime)) //序列化耗时： 30.0057ms

		fmt.Println("写文件按耗时：", finishWriteFileTime.Sub(startTime)) //写文件按耗时： 33.0036ms

	}

	//Protobuf Unmarshal
	{
		startTime := time.Now()
		pLister := &apis.PersonalInfomationList{}
		data, _ := ioutil.ReadFile("E:/Geek/src/learn.go/chapter08/0.4.think/data.protobuf")
		proto.Unmarshal(data, pLister)
		//fmt.Println(pLister.Items)
		finishTime := time.Now()
		fmt.Println("Protobuf Unmarshal:", finishTime.Sub(startTime))
	}

}
