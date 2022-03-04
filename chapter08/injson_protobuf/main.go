package main

//没有使用编码的情况下，要自己处理。很可能会出错。

import (
	"encoding/json"
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"io/ioutil"
	"learn.go/pkg/apis"
	"log"
	"os"
)

func main() {
	//caseStudy()
	input := &inputFromStd{} //初始信息结构体 //指针，并实例化
	calc := &Calc{}          //对初始信息进行计算，计算体脂率等
	rk := &FatRateRank{}     // 排行榜结构体
	//records := &record{}
	records := NewRecord("e:/Geek/src/learn.go/chapter08/injson_protobuf/小强.self.infomation") //初始化记录结构体

	for {
		pi := input.GetInput()                                     // 输入初始信息 返回值是json标准的
		if err := records.savePersonalInfomation(pi); err != nil { //savePersonalInfomation 的返回值时error
			log.Fatal("保存失败：", err)
		}
		fr, err := calc.FatRate(pi)
		if err != nil {
			log.Fatal("计算体脂率失败：", err)
		}

		rk.inputRecord(pi.Name, fr)

		rankResult, _ := rk.getRank(pi.Name)
		fmt.Println("排名结果：", rankResult) //直接打印结果，并没有把排名写进文件里。
	}
}

func caseStudy() {
	filePath := "e:/Geek/src/learn.go/chapter08/injson_protobuf/小强.self.infomation"
	personalInfomation := apis.PersonalInfomation{
		Name:   `""小"强"`,
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}

	fmt.Printf("%+v\n", personalInfomation) //%+v: {Name:小强 Sex:男 Tall:1.7 Weight:71 Age:35},  %v: {小强 男 1.7 71 35}

	// 把结构体格式转成json格式要用到json.marshal
	data, err := json.Marshal(personalInfomation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("marshal 的结果是(原生)：", data)
	fmt.Println("marshal 的结果是（string）：", string(data))

	writeFile(filePath, data)
	readFile(filePath)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath) // data 是读到的内容
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("读取出来的内容是：", string(data))
	//	infos := strings.Split(string(data), ",") //表示数据遇到 ","(逗号)，分割开。(小强，男 --> 小强 男) 分隔开就可以struct等使用。

	personalInfomation := apis.PersonalInfomation{}
	json.Unmarshal(data, &personalInfomation) // 把json格式的data，转成go的对象格式，并存入该对象中。

	fmt.Println("开始计算体脂信息：", personalInfomation)
	bmi, _ := gobmi.BMI(float64(personalInfomation.Weight), float64(personalInfomation.Tall)) // todo handle error
	fmt.Printf("%s 的 BMI是：%v\n", personalInfomation.Name, bmi)
	fatRate := gobmi.CalcFatRate(bmi, int(personalInfomation.Age), personalInfomation.Sex)
	fmt.Printf("%s 的 体脂率是：%v\n", personalInfomation.Name, fatRate)
}

func writeFile(filePath string, data []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(data)
	fmt.Println(err)
}
