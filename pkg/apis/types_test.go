package apis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
	// todo yaml 包
)

func TestMarshalJson(t *testing.T) {
	personalInfomation := PersonalInfomation{
		Name:   "xiao",
		Sex:    "man",
		Tall:   1.7,
		Weight: 71,
		Age:    35,
	}
	fmt.Printf("%+v\n", personalInfomation)
	data, err := json.Marshal(personalInfomation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal 的结果是：", data)
	fmt.Println("marshal 的结果是：", string(data))
}

func TestUnmarshalJson(t *testing.T) {
	data := `{"name":"xiaoqiang","sex":"man","tall":1.7,"weight":71,"age":35}` // ` ` 不转义字符串
	fmt.Println(data)
	personalInfomation := PersonalInfomation{}
	json.Unmarshal([]byte(data), &personalInfomation) //[]byte(data) 把data转为[]byte类型
	fmt.Println(personalInfomation)
}

/* //todo yaml包还未安装
func TestMarshalYaml(t *testing.T) {
	personalInfomation := PersonalInfomation{
		Name:   "xiao",
		Sex:    "man",
		Tall:   1.7,
		Weight: 71,
		Age:    35,
	}
	fmt.Printf("%+v\n", personalInfomation)
	data, err := yaml.Marshal(personalInfomation)
	//yaml.Marshal(personalInfomation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal 的结果是：", data)
	fmt.Println("marshal 的结果是：", string(data))

}

*/

/*
func TestUnmarshalYaml(t *testing.T) {
	data := `{"name":"xiaoqiang","sex":"man","tall":1.7,"weight":71,"age":35}` // ` ` 不转义字符串
	fmt.Println(data)
	personalInfomation := PersonalInfomation{}
	yaml.Unmarshal([]byte(data), &personalInfomation) //[]byte(data) 把data转为[]byte类型
	fmt.Println(personalInfomation)
}


*/
func TestMarshalProtobuf(t *testing.T) {
	personalInfomation := &PersonalInfomation{
		Name:   "xiao",
		Sex:    "man",
		Tall:   1.7,
		Weight: 71,
		Age:    35,
	}
	data, err := proto.Marshal(personalInfomation) //参数对应的是个接口，所以要用指针类型。
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
	fmt.Println(string(data)) //xiaoman���?%  �B(#  人是看不懂的（二进制文件）
	// 通常在非程序交互过程中，要保留原生protobuf，可以直接写入文件。如果想要单行保存，必须转码。
	// 选择的通用转码是：base64 （base64可以把任何类型的转成可以看的ASCII码范围）
	output64Data := base64.StdEncoding.EncodeToString(data)
	fmt.Println(">>>>", output64Data) //>>>> CgR4aWFvEgNtYW4dmpnZPyUAAI5CKCM= 我们依然看不懂，但是程序可以看懂。
	//把这个保存在文件后，就可供读取使用。
}
