package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonData := `{"name":"小强","sex":"男","tall":1.8,"weight":70,"age":30}`
	protoDataBase64 := `EgblsI/lvLoaA+eUtyVmZuY/LQAAjEIwHg==`

	pi1 := &PersonalInformation{}
	json.Unmarshal([]byte(jsonData), pi1)
	fmt.Println("解析JSON：")
	fmt.Printf("%+v\n", *pi1) // 如果改了 标准变量后的数字，会让一些数据无法识别。

	pi2 := &PersonalInformation{}
	protoData, _ := base64.StdEncoding.DecodeString(protoDataBase64)
	proto.Unmarshal(protoData, pi2)
	fmt.Println("解析protobuf：")
	fmt.Printf("%+v\n", *pi2)

}
