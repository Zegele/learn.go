package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"learn.go/pkg/apis"
	"log"
	"os"
)

func NewRecord(filePath string) *record { //注意：初始化函数  关于record结构体的初始化函数 可以保证整个使用过程正确。不用在主函数中，专门对该结构体进行初始化
	return &record{
		filePath:         filePath,           // 默认是json文件 .json?
		yamlFilePath:     filePath + ".yaml", //没有取掉 .json
		protobufFilePath: filePath + ".proto.base64",
	}
}

type record struct {
	filePath         string
	yamlFilePath     string
	protobufFilePath string
}

func (r *record) savePersonalInfomation(pi *apis.PersonalInfomation) error {
	{ //JSON
		data, err := json.Marshal(pi) //把pi数据转成json格式
		if err != nil {
			fmt.Println("marshal 出错：", err)
			return err
		}
		if err := r.writeFileWithAppendJson(data); err != nil {
			log.Println("写入JSON时出错：", err)
			return err
		}
	}
	//{//todo yaml 包还没有安装
	//	data, err := yaml.Marshal(pi) //把pi数据转成json格式
	//	if err != nil {
	//		fmt.Println("marshal 出错：", err)
	//		return err
	//	}
	//	if err := r.writeFileWithAppendYaml(data); err != nil {
	//		log.Println("写入YAML时出错：", err)
	//		return err
	//	}
	//	return nil
	//}
	{
		data, err := proto.Marshal(pi) //把pi数据转成json格式
		if err != nil {
			fmt.Println("marshal 出错：", err)
			return err
		}
		if err := r.writeFileWithAppendProtobuf(data); err != nil {
			log.Println("写入PROTOBUF时出错：", err)
			return err
		}
		return nil
	}
}

func (r *record) writeFileWithAppendJson(data []byte) error {
	file, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //创建一个文件； 写入r.filePath地址的文件
	if err != nil {
		fmt.Println("无法打开文件：", r.filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close() // close

	//_, err = file.Write(data) //  没有换行符
	_, err = file.Write(append(data, '\n')) // 使用了换行
	//这里用了单引号 单引号里面是 单个 字符，对应的值是改为字符的ASCII值。
	// 双引号：里面可以是单个字符，也可以是字符串，且可以有转义字符 如\n \r（表示回车）等。 对应string类型
	// 反引号：表示原生意思，反引号中的内容可以是多行内容，不支持转义。
	if err != nil {
		return err
	}
	return nil
}
func (r *record) writeFileWithAppendYaml(data []byte) error {
	file, err := os.OpenFile(r.yamlFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //创建一个文件； 写入r.filePath地址的文件
	if err != nil {
		fmt.Println("无法打开文件：", r.filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close() // close

	//_, err = file.Write(data) //  没有换行符
	_, err = file.Write(append(data, '\n')) // 使用了换行
	//这里用了单引号 单引号里面是 单个 字符，对应的值是改为字符的ASCII值。
	// 双引号：里面可以是单个字符，也可以是字符串，且可以有转义字符 如\n \r（表示回车）等。 对应string类型
	// 反引号：表示原生意思，反引号中的内容可以是多行内容，不支持转义。
	if err != nil {
		return err
	}
	return nil
}

func (r *record) writeFileWithAppendProtobuf(data []byte) error {
	file, err := os.OpenFile(r.protobufFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //创建一个文件； 写入r.filePath地址的文件
	if err != nil {
		fmt.Println("无法打开文件：", r.filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close() // close

	//_, err = file.Write(data) //  没有换行符
	_, err = file.Write([]byte(base64.StdEncoding.EncodeToString(data))) // EncodeToString 返回值是string类型，用[]byte()转换成字节切片类型。
	//这里用了单引号 单引号里面是 单个 字符，对应的值是改为字符的ASCII值。
	// 双引号：里面可以是单个字符，也可以是字符串，且可以有转义字符 如\n \r（表示回车）等。 对应string类型
	// 反引号：表示原生意思，反引号中的内容可以是多行内容，不支持转义。
	if err != nil {
		return err
	}
	return nil
}
