package client

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"learn.go/zuoye/zuoye5_tizhi/pkg/apii"
	"log"
	"os"
)

type InputData struct {
}

func (InputData) Input(p *apii.Person) *apii.Person {
	p = InputName(p)
	p = InputAge(p)
	p = InputSex(p)
	p = InputWeight(p)
	p = InputTall(p)
	return p
}

// 输入的方式
func InputName(p *apii.Person) *apii.Person { //输入姓名
	for {
		fmt.Print("请输入姓名：")
		fmt.Scanln(&p.Name)
		if p.Name != "" {
			return p
		}
		fmt.Println("姓名不能为空，重新输入。")
	}
}

func InputAge(p *apii.Person) *apii.Person { //输入年龄
	for {
		fmt.Print("请输入年龄：")
		fmt.Scanln(&p.Age)
		if p.Age >= 10 && p.Age <= 150 {
			return p
		}
		fmt.Println("抱歉，不在符合的年龄范围，请输入10-150的整数！")
	}
}

func InputSex(p *apii.Person) *apii.Person { //输入性别
	for {
		fmt.Print("请输入性别（男/女）：")
		fmt.Scanln(&p.Sex)
		if p.Sex == "男" || p.Sex == "man" {
			return p
		} else if p.Sex == "女" || p.Sex == "woman" {
			return p
		} else {
			fmt.Println("性别输入有误。[男（man）或女（woman）]")
		}
	}
}

func InputWeight(p *apii.Person) *apii.Person { //输入体重
	for {
		fmt.Print("请输入体重（kg）：")
		fmt.Scanln(&p.Weight)
		for p.Weight >= 20 && p.Weight < 1000 {
			return p
		}
		fmt.Println("抱歉，数据有误，请输入20-1000之间的数。")
	}
}

func InputTall(p *apii.Person) *apii.Person { //输入身高
	for {
		fmt.Print("请输入身高（米）：")
		fmt.Scanln(&p.Tall)
		for p.Tall >= 0.5 && p.Tall <= 3 {
			return p
		}
		fmt.Println("抱歉，输入有误，请输入0.5-3之间的数。")
	}
}

// 批量生成 如 生成1000个用户信息
func (*InputData) Register(number int, c chan *apii.Person) (p *apii.Person) { //第一步：把数据注入channel
	p = &apii.Person{}
	p.Name = "a" + fmt.Sprint(number)
	p.Age = 10 + int64(number)
	p.Sex = "男"
	p.Tall = 1.8
	p.Weight = 70
	c <- p
	return
}

// client信息写入

//JSON 不换行
func (*InputData) WriteClientByJSON(filePath string, ps *apii.Persons) { // 直接使用protobuf生产的apii包结构体
	data, err := json.Marshal(ps) //需要字节切片类型
	//data, err := json.Marshal(ps.Items) //这种也能正常生产json文件，但unmarshal的时候有问题
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(filePath, data, 0777)
}

// JSON 换行
func (*InputData) WriteClientByJSONHuanHang(filePath string, ps *apii.Persons) { // 直接使用protobuf生产的apii包结构体
	file, err := os.Create(filePath) //创建一个文件
	if err != nil {
		fmt.Println("无法打开文件：", filePath, "错误信息是：", err)
		os.Exit(1) //如果程序正常退出，Exit内的值是0，如果不是正常退出，给一个非0的数。
	}
	defer file.Close() // close

	for _, v := range ps.Items {
		datatemp, err := json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}
		datatemp = append(datatemp, '\n') // 加换行符
		file.Write(datatemp)
	}
}

//PROTOBUF 不换行
func (*InputData) WriteClientByProtobuf(filePath string, ps *apii.Persons) {
	data, err := proto.Marshal(ps) //proto.Marshal参数是要message类型
	if err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile(filePath, data, 0777); err != nil {
		log.Fatal(err)
	}
}

//PROTOBUF 不换行 // todo
//func (*InputData) WriteClientByProtobufHuanHang(filePath string, ps *apii.Persons) *apii.Persons {
//	file, err := os.Create(filePath) //创建一个文件
//	if err != nil {
//		fmt.Println("无法打开文件：", filePath, "错误信息是：", err)
//		os.Exit(1) //如果程序正常退出，Exit内的值是0，如果不是正常退出，给一个非0的数。
//	}
//	defer file.Close() // close
//	pshou := &apii.Persons{}
//	pItems := pshou.Items
//	for _, v := range ps.Items {
//		pItems = append(pItems, v)
//
//
//		pItems = append(pItems, '\n') // 加换行符
//
//	}
//}
